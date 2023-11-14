// Copyright (c) 2015-2023 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7/pkg/tags"
	"github.com/minio/minio/internal/bucket/versioning"
	xhttp "github.com/minio/minio/internal/http"
	"github.com/minio/minio/internal/logger"
	"github.com/minio/pkg/env"
	"github.com/minio/pkg/wildcard"
	"github.com/minio/pkg/workers"
)

// expire: # Expire objects that match a condition
//   apiVersion: v1
//   bucket: mybucket # Bucket where this batch job will expire matching objects from
//   prefix: myprefix # (Optional) Prefix under which this job will expire objects matching the rules below.
//   rules:
//     - type: object  # regular objects with zero ore more older versions
//       name: NAME # match object names that satisfy the wildcard expression.
//       olderThan: 70h # match objects older than this value
//       createdBefore: "2006-01-02T15:04:05.00Z" # match objects created before "date"
//       tags:
//         - key: name
//           value: pick* # match objects with tag 'name', all values starting with 'pick'
//       metadata:
//         - key: content-type
//           value: image/* # match objects with 'content-type', all values starting with 'image/'
//       size:
//         lesserThan: "10MiB" # match objects with size lesser than this value (e.g. 10MiB)
//         greaterThan: 1MiB # match objects with size greater than this value (e.g. 1MiB)
//       purge:
//           # retainVersions: 0 # (default) delete all versions of the object. This option is the fastest.
//           # retainVersions: 5 # keep the latest 5 versions of the object.
//
//     - type: deleted # objects with delete marker as their latest version
//       name: NAME # match object names that satisfy the wildcard expression.
//       olderThan: 10h # match objects older than this value (e.g. 7d10h31s)
//       createdBefore: "2006-01-02T15:04:05.00Z" # match objects created before "date"
//       purge:
//           # retainVersions: 0 # (default) delete all versions of the object. This option is the fastest.
//           # retainVersions: 5 # keep the latest 5 versions of the object including delete markers.
//
//   notify:
//     endpoint: https://notify.endpoint # notification endpoint to receive job completion status
//     token: Bearer xxxxx # optional authentication token for the notification endpoint
//
//   retry:
//     attempts: 10 # number of retries for the job before giving up
//     delay: 500ms # least amount of delay between each retry

//go:generate msgp -file $GOFILE

// BatchJobExpirePurge type accepts non-negative versions to be retained
type BatchJobExpirePurge struct {
	RetainVersions int `yaml:"retainVersions" json:"retainVersions"`
}

// Validate returns nil if value is valid, ie > 0.
func (p BatchJobExpirePurge) Validate() error {
	if p.RetainVersions < 0 {
		return errors.New("retainVersions must be >= 0")
	}
	return nil
}

// BatchJobExpireFilter holds all the filters currently supported for batch replication
type BatchJobExpireFilter struct {
	OlderThan     time.Duration `yaml:"olderThan,omitempty" json:"olderThan"`
	CreatedBefore time.Time     `yaml:"createdBefore,omitempty" json:"createdBefore"`
	Tags          []BatchJobKV  `yaml:"tags,omitempty" json:"tags"`
	Metadata      []BatchJobKV  `yaml:"metadata,omitempty" json:"metadata"`
	Size          struct {
		LesserThan  BatchJobSize `yaml:"lesserThan,omitempty" json:"lesserThan"`
		GreaterThan BatchJobSize `yaml:"greaterThan,omitempty" json:"greaterThan"`
	} `yaml:"size" json:"size"`
	Type  string              `yaml:"type" json:"type"`
	Name  string              `yaml:"name" json:"name"`
	Purge BatchJobExpirePurge `yaml:"purge" json:"purge"`
}

// Matches returns true if obj matches the filter conditions specified in ef.
func (ef BatchJobExpireFilter) Matches(obj ObjectInfo, now time.Time) bool {
	switch ef.Type {
	case BatchJobExpireObject:
		if obj.DeleteMarker {
			return false
		}
	case BatchJobExpireDeleted:
		if !obj.DeleteMarker {
			return false
		}
	default:
		// we should never come here, Validate should have caught this.
		return false
	}

	if len(ef.Name) > 0 && !wildcard.Match(ef.Name, obj.Name) {
		return false
	}
	if ef.OlderThan > 0 && now.Sub(obj.ModTime) <= ef.OlderThan {
		return false
	}

	if !ef.CreatedBefore.IsZero() && (obj.ModTime.Equal(ef.CreatedBefore) || obj.ModTime.After(ef.CreatedBefore)) {
		return false
	}

	if len(ef.Tags) > 0 && !obj.DeleteMarker {
		// Only parse object tags if tags filter is specified.
		var tagMap map[string]string
		if len(obj.UserTags) != 0 {
			t, err := tags.ParseObjectTags(obj.UserTags)
			if err != nil {
				return false
			}
			tagMap = t.ToMap()
		}

		for _, kv := range ef.Tags {
			// Object (version) must match all tags specified in
			// the filter
			var match bool
			for t, v := range tagMap {
				if kv.Match(BatchJobKV{Key: t, Value: v}) {
					match = true
				}
			}
			if !match {
				return false
			}
		}

	}
	if len(ef.Metadata) > 0 && !obj.DeleteMarker {
		for _, kv := range ef.Metadata {
			// Object (version) must match all x-amz-meta and
			// standard metadata headers
			// specified in the filter
			var match bool
			for k, v := range obj.UserDefined {
				if !stringsHasPrefixFold(k, "x-amz-meta-") && !isStandardHeader(k) {
					continue
				}
				// We only need to match x-amz-meta or standardHeaders
				if kv.Match(BatchJobKV{Key: k, Value: v}) {
					match = true
				}
			}
			if !match {
				return false
			}
		}
	}

	if ef.Size.LesserThan > 0 {
		if !ef.Size.LesserThan.LesserThan(obj.Size) {
			return false
		}
	}

	if ef.Size.GreaterThan > 0 {
		if !ef.Size.GreaterThan.GreaterThan(obj.Size) {
			return false
		}
	}

	return true
}

const (
	// BatchJobExpireObject - object type
	BatchJobExpireObject string = "object"
	// BatchJobExpireDeleted - delete marker type
	BatchJobExpireDeleted string = "deleted"
)

// Validate returns nil if ef has valid fields, validation error otherwise.
func (ef BatchJobExpireFilter) Validate() error {
	switch ef.Type {
	case BatchJobExpireObject:
	case BatchJobExpireDeleted:
		if len(ef.Tags) > 0 || len(ef.Metadata) > 0 {
			return errors.New("invalid batch-expire rule filter")
		}
	default:
		return errors.New("invalid batch-expire type")
	}

	for _, tag := range ef.Tags {
		if err := tag.Validate(); err != nil {
			return err
		}
	}

	for _, meta := range ef.Metadata {
		if err := meta.Validate(); err != nil {
			return err
		}
	}
	if err := ef.Purge.Validate(); err != nil {
		return err
	}
	return nil
}

// BatchJobExpireV1 v1 of batch expiration job
type BatchJobExpireV1 struct {
	APIVersion      string                 `yaml:"apiVersion" json:"apiVersion"`
	Bucket          string                 `yaml:"bucket" json:"bucket"`
	Prefix          string                 `yaml:"prefix" json:"prefix"`
	NotificationCfg BatchJobNotification   `yaml:"notify" json:"notify"`
	Retry           BatchJobRetry          `yaml:"retry" json:"retry"`
	Rules           []BatchJobExpireFilter `yaml:"rules" json:"rules"`
}

// Notify notifies notification endpoint if configured regarding job failure or success.
func (r BatchJobExpireV1) Notify(ctx context.Context, body io.Reader) error {
	if r.NotificationCfg.Endpoint == "" {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.NotificationCfg.Endpoint, body)
	if err != nil {
		return err
	}

	if r.NotificationCfg.Token != "" {
		req.Header.Set("Authorization", r.NotificationCfg.Token)
	}

	clnt := http.Client{Transport: getRemoteInstanceTransport}
	resp, err := clnt.Do(req)
	if err != nil {
		return err
	}

	xhttp.DrainBody(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// Expire expires object versions which have already matched supplied filter conditions
func (r *BatchJobExpireV1) Expire(ctx context.Context, api ObjectLayer, vc *versioning.Versioning, objsToDel []ObjectToDelete) []error {
	opts := ObjectOptions{
		PrefixEnabledFn:  vc.PrefixEnabled,
		VersionSuspended: vc.Suspended(),
	}
	_, errs := api.DeleteObjects(ctx, r.Bucket, objsToDel, opts)
	return errs
}

const (
	batchExpireName                 = "batch-expire.bin"
	batchExpireFormat               = 1
	batchExpireVersionV1            = 1
	batchExpireVersion              = batchExpireVersionV1
	batchExpireAPIVersion           = "v1"
	batchExpireJobDefaultRetries    = 3
	batchExpireJobDefaultRetryDelay = 250 * time.Millisecond
)

type objInfoCache map[string]*ObjectInfo

func newObjInfoCache() objInfoCache {
	return objInfoCache(make(map[string]*ObjectInfo))
}

func (oiCache objInfoCache) Add(toDel ObjectToDelete, oi *ObjectInfo) {
	oiCache[fmt.Sprintf("%s-%s", toDel.ObjectName, toDel.VersionID)] = oi
}

func (oiCache objInfoCache) Get(toDel ObjectToDelete) (*ObjectInfo, bool) {
	oi, ok := oiCache[fmt.Sprintf("%s-%s", toDel.ObjectName, toDel.VersionID)]
	return oi, ok
}

func batchObjsForDelete(ctx context.Context, r *BatchJobExpireV1, ri *batchJobInfo, job BatchJobRequest, api ObjectLayer, wk *workers.Workers, expireCh <-chan []expireObjInfo) {
	defer wk.Give()

	vc, _ := globalBucketVersioningSys.Get(r.Bucket)
	retryAttempts := r.Retry.Attempts
	delay := job.Expire.Retry.Delay
	if delay == 0 {
		delay = batchExpireJobDefaultRetryDelay
	}

	var i int
	for toExpire := range expireCh {
		if i > 0 {
			if wait := globalBatchConfig.Clone().ExpirationWorkersWait; wait > 0 {
				time.Sleep(wait)
			}
		}
		i++
		wk.Take()
		go func(toExpire []expireObjInfo) {
			defer wk.Give()

			toExpireAll := make([]ObjectInfo, 0, len(toExpire))
			toDel := make([]ObjectToDelete, 0, len(toExpire))
			oiCache := newObjInfoCache()
			for _, exp := range toExpire {
				if exp.ExpireAll {
					toExpireAll = append(toExpireAll, exp.ObjectInfo)
					continue
				}
				// Cache ObjectInfo value via pointers for
				// subsequent use to track objects which
				// couldn't be deleted.
				od := ObjectToDelete{
					ObjectV: ObjectV{
						ObjectName: exp.Name,
						VersionID:  exp.VersionID,
					},
				}
				toDel = append(toDel, od)
				oiCache.Add(od, &exp.ObjectInfo)
			}

			// DeleteObject(deletePrefix: true) to expire all versions of an object
			for _, exp := range toExpireAll {
				var success bool
				for attempts := 1; attempts <= retryAttempts; attempts++ {
					stopFn := globalBatchJobsMetrics.trace(batchJobMetricExpire, ri.JobID, attempts)
					_, err := api.DeleteObject(ctx, exp.Bucket, exp.Name, ObjectOptions{
						DeletePrefix: true,
						NoLock:       true,
					})
					if err != nil {
						stopFn(exp, err)
						logger.LogIf(ctx, fmt.Errorf("Failed to expire %s/%s versionID=%s due to %v (attempts=%d)", toExpire[i].Bucket, toExpire[i].Name, toExpire[i].VersionID, err, attempts))
					} else {
						stopFn(exp, err)
						success = true
						break
					}
				}
				ri.trackMultipleObjectVersions(r.Bucket, exp, success)
			}

			// DeleteMultiple objects
			toDelCopy := make([]ObjectToDelete, len(toDel))
			for attempts := 1; attempts <= retryAttempts; attempts++ {
				stopFn := globalBatchJobsMetrics.trace(batchJobMetricExpire, ri.JobID, attempts)
				// Copying toDel to select from objects whose
				// deletion failed
				copy(toDelCopy, toDel)
				var failed int
				errs := r.Expire(ctx, api, vc, toDel)
				// reslice toDel in preparation for next retry
				// attempt
				toDel = toDel[:0]
				for i, err := range errs {
					if err != nil {
						stopFn(toDelCopy[i], err)
						logger.LogIf(ctx, fmt.Errorf("Failed to expire %s/%s versionID=%s due to %v (attempts=%d)", ri.Bucket, toDelCopy[i].ObjectName, toDelCopy[i].VersionID, err, attempts))
						failed++
						if attempts == retryAttempts { // all retry attempts failed, record failure
							if oi, ok := oiCache.Get(toDelCopy[i]); ok {
								ri.trackCurrentBucketObject(r.Bucket, *oi, false)
							}
						} else {
							toDel = append(toDel, toDelCopy[i])
						}
					} else {
						stopFn(toDelCopy[i], nil)
						if oi, ok := oiCache.Get(toDelCopy[i]); ok {
							ri.trackCurrentBucketObject(r.Bucket, *oi, true)
						}
					}
				}

				globalBatchJobsMetrics.save(ri.JobID, ri)

				if failed == 0 {
					break
				}

				// Add a delay between retry attempts
				if attempts < retryAttempts {
					time.Sleep(delay)
				}
			}
		}(toExpire)
	}
}

type expireObjInfo struct {
	ObjectInfo
	ExpireAll bool
}

// Start the batch expiration job, resumes if there was a pending job via "job.ID"
func (r *BatchJobExpireV1) Start(ctx context.Context, api ObjectLayer, job BatchJobRequest) error {
	ri := &batchJobInfo{
		JobID:     job.ID,
		JobType:   string(job.Type()),
		StartTime: job.Started,
	}
	if err := ri.load(ctx, api, job); err != nil {
		return err
	}

	globalBatchJobsMetrics.save(job.ID, ri)
	lastObject := ri.Object

	now := time.Now().UTC()

	workerSize, err := strconv.Atoi(env.Get("_MINIO_BATCH_EXPIRATION_WORKERS", strconv.Itoa(runtime.GOMAXPROCS(0)/2)))
	if err != nil {
		return err
	}

	wk, err := workers.New(workerSize)
	if err != nil {
		// invalid worker size.
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make(chan ObjectInfo, workerSize)
	if err := api.Walk(ctx, r.Bucket, r.Prefix, results, ObjectOptions{
		WalkMarker:       lastObject,
		WalkLatestOnly:   false, // we need to visit all versions of the object to implement purge: retainVersions
		WalkVersionsSort: WalkVersionsSortDesc,
	}); err != nil {
		// Do not need to retry if we can't list objects on source.
		return err
	}

	// Goroutine to periodically save batch-expire job's in-memory state
	saverQuitCh := make(chan struct{})
	go func() {
		saveTicker := time.NewTicker(10 * time.Second)
		defer saveTicker.Stop()
		for {
			select {
			case <-saveTicker.C:
				// persist in-memory state to disk after every 10secs.
				logger.LogIf(ctx, ri.updateAfter(ctx, api, 10*time.Second, job))

			case <-ctx.Done():
				// persist in-memory state immediately before exiting due to context cancelation.
				logger.LogIf(ctx, ri.updateAfter(ctx, api, 0, job))
				return

			case <-saverQuitCh:
				// persist in-memory state immediately to disk.
				logger.LogIf(ctx, ri.updateAfter(ctx, api, 0, job))
				return
			}
		}
	}()

	expireCh := make(chan []expireObjInfo, workerSize)
	wk.Take() // For batchObjsForDelete goroutine

	go batchObjsForDelete(ctx, r, ri, job, api, wk, expireCh)

	var (
		prevObj       ObjectInfo
		matchedFilter BatchJobExpireFilter
		versionsCount int
		toDel         []expireObjInfo
	)
	for result := range results {
		result := result
		// Apply filter to find the matching rule to apply expiry
		// actions accordingly.
		// nolint:gocritic
		if result.IsLatest {
			// send down filtered entries to be deleted using
			// DeleteObjects method
			if len(toDel) > 10 { // batch up to 10 objects/versions to be expired simultaneously.
				xfer := make([]expireObjInfo, len(toDel))
				copy(xfer, toDel)
				expireCh <- xfer
				toDel = toDel[:0] // resetting toDel
			}
			var match BatchJobExpireFilter
			var found bool
			for _, rule := range r.Rules {
				if rule.Matches(result, now) {
					match = rule
					found = true
					break
				}
			}
			if !found {
				continue
			}

			prevObj = result
			matchedFilter = match
			versionsCount = 1
			// Include the latest version
			if matchedFilter.Purge.RetainVersions == 0 {
				toDel = append(toDel, expireObjInfo{
					ObjectInfo: result,
					ExpireAll:  true,
				})
				continue
			}
		} else if prevObj.Name == result.Name {
			if matchedFilter.Purge.RetainVersions == 0 {
				continue // including latest version in toDel suffices, skipping other versions
			}
			versionsCount++
		} else {
			continue
		}

		if versionsCount <= matchedFilter.Purge.RetainVersions {
			continue // retain versions
		}
		toDel = append(toDel, expireObjInfo{
			ObjectInfo: result,
		})
	}
	// Send any remaining objects downstream
	if len(toDel) > 0 {
		select {
		case <-ctx.Done():
		case expireCh <- toDel:
		}
	}
	close(expireCh)

	wk.Wait() // waits for all expire goroutines to complete

	ri.Complete = ri.ObjectsFailed == 0
	ri.Failed = ri.ObjectsFailed > 0
	globalBatchJobsMetrics.save(job.ID, ri)

	// Close the saverQuitCh - this also triggers saving in-memory state
	// immediately one last time before we exit this method.
	close(saverQuitCh)

	// Notify expire jobs final status to the configured endpoint
	buf, _ := json.Marshal(ri)
	if err := r.Notify(ctx, bytes.NewReader(buf)); err != nil {
		logger.LogIf(ctx, fmt.Errorf("unable to notify %v", err))
	}

	return nil
}

//msgp:ignore batchExpireJobError
type batchExpireJobError struct {
	Code           string
	Description    string
	HTTPStatusCode int
}

func (e batchExpireJobError) Error() string {
	return e.Description
}

// Validate validates the job definition input
func (r *BatchJobExpireV1) Validate(ctx context.Context, job BatchJobRequest, o ObjectLayer) error {
	if r == nil {
		return nil
	}

	if r.APIVersion != batchExpireAPIVersion {
		return errInvalidArgument
	}

	if r.Bucket == "" {
		return errInvalidArgument
	}

	if _, err := o.GetBucketInfo(ctx, r.Bucket, BucketOptions{}); err != nil {
		if isErrBucketNotFound(err) {
			return batchExpireJobError{
				Code:           "NoSuchSourceBucket",
				Description:    "The specified source bucket does not exist",
				HTTPStatusCode: http.StatusNotFound,
			}
		}
		return err
	}

	if len(r.Rules) > 10 {
		return errors.New("Too many rules. Batch expire job can't have more than 10 rules")
	}

	for _, rule := range r.Rules {
		if err := rule.Validate(); err != nil {
			return err
		}
	}

	if err := r.Retry.Validate(); err != nil {
		return err
	}
	return nil
}
