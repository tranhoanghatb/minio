// Copyright (c) 2015-2021 MinIO, Inc.
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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio/internal/logger"
	"github.com/minio/pkg/v2/wildcard"
)

// BucketQuotaSys - map of bucket and quota configuration.
type BucketQuotaSys struct{}

// Get - Get quota configuration.
func (sys *BucketQuotaSys) Get(ctx context.Context, bucketName string) (*madmin.BucketQuota, error) {
	cfg, _, err := globalBucketMetadataSys.GetQuotaConfig(ctx, bucketName)
	return cfg, err
}

// NewBucketQuotaSys returns initialized BucketQuotaSys
func NewBucketQuotaSys() *BucketQuotaSys {
	return &BucketQuotaSys{}
}

var bucketStorageCache timedValue

// Init initialize bucket quota.
func (sys *BucketQuotaSys) Init(objAPI ObjectLayer) {
	bucketStorageCache.Once.Do(func() {
		// Set this to 10 secs since its enough, as scanner
		// does not update the bucket usage values frequently.
		bucketStorageCache.TTL = 10 * time.Second
		// Rely on older value if usage loading fails from disk.
		bucketStorageCache.Relax = true
		bucketStorageCache.Update = func() (interface{}, error) {
			ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
			defer done()

			return loadDataUsageFromBackend(ctx, objAPI)
		}
	})
}

// GetBucketUsageInfo return bucket usage info for a given bucket
func (sys *BucketQuotaSys) GetBucketUsageInfo(bucket string) (BucketUsageInfo, error) {
	v, err := bucketStorageCache.Get()
	timedout := OperationTimedOut{}
	if err != nil && !errors.Is(err, context.DeadlineExceeded) && !errors.As(err, &timedout) {
		if v != nil {
			logger.LogOnceIf(GlobalContext, fmt.Errorf("unable to retrieve usage information for bucket: %s, relying on older value cached in-memory: err(%v)", bucket, err), "bucket-usage-cache-"+bucket)
		} else {
			logger.LogOnceIf(GlobalContext, errors.New("unable to retrieve usage information for bucket: %s, no reliable usage value available - quota will not be enforced"), "bucket-usage-empty-"+bucket)
		}
	}

	var bui BucketUsageInfo
	dui, ok := v.(DataUsageInfo)
	if ok {
		bui = dui.BucketsUsage[bucket]
	}

	return bui, nil
}

// parseBucketQuota parses BucketQuota from json
func parseBucketQuota(bucket string, data []byte) (quotaCfg *madmin.BucketQuota, err error) {
	quotaCfg = &madmin.BucketQuota{}
	if err = json.Unmarshal(data, quotaCfg); err != nil {
		return quotaCfg, err
	}
	if !quotaCfg.IsValid() {
		if quotaCfg.Type == "fifo" {
			logger.LogIf(GlobalContext, errors.New("Detected older 'fifo' quota config, 'fifo' feature is removed and not supported anymore. Please clear your quota configs using 'mc admin bucket quota alias/bucket --clear' and use 'mc ilm add' for expiration of objects"))
			return quotaCfg, fmt.Errorf("invalid quota type 'fifo'")
		}
		return quotaCfg, fmt.Errorf("Invalid quota config %#v", quotaCfg)
	}
	return
}

func (sys *BucketQuotaSys) enforceQuotaHard(ctx context.Context, bucket, api string, size int64) error {
	if size < 0 {
		return nil
	}

	q, err := sys.Get(ctx, bucket)
	if err != nil {
		return err
	}

	var quotaSize uint64
	if q != nil && q.Type == madmin.HardQuota {
		if q.Size > 0 {
			quotaSize = q.Size
		} else if q.Quota > 0 {
			quotaSize = q.Quota
		}
	}
	if quotaSize > 0 {
		if uint64(size) >= quotaSize { // check if file size already exceeds the quota
			return BucketQuotaExceeded{Bucket: bucket}
		}

		bui, err := sys.GetBucketUsageInfo(bucket)
		if err != nil {
			return err
		}

		if bui.Size > 0 && ((bui.Size + uint64(size)) >= quotaSize) {
			return BucketQuotaExceeded{Bucket: bucket}
		}
	}

	return enforceBucketThrottle(bucket, api, q)
}

func enforceBucketQuotaHard(ctx context.Context, bucket, api string, size int64) error {
	if globalBucketQuotaSys == nil {
		return nil
	}
	return globalBucketQuotaSys.enforceQuotaHard(ctx, bucket, api, size)
}

func enforceBucketThrottle(bucket, api string, qCfg *madmin.BucketQuota) error {
	// if throttle rules not set, return
	if !qCfg.IsBucketThrottled() {
		return nil
	}

	// Create map of unique APIs and their allowed no of calls from throttle rules
	s3AllowedReqCountMap := make(map[string]uint64)
	for _, rule := range qCfg.ThrottleRules {
		for _, api := range rule.APIs {
			if count, ok := s3AllowedReqCountMap[api]; ok {
				// apply the smaller value for the concurrent request count for the API
				if count > rule.ConcurrentRequestsCount {
					s3AllowedReqCountMap[api] = rule.ConcurrentRequestsCount
				}
			} else {
				s3AllowedReqCountMap[api] = rule.ConcurrentRequestsCount
			}
		}
	}

	// Get the current count for the API from stats
	currReqCount, ok := globalHTTPStats.toServerHTTPStats().CurrentS3Requests.APIStats[strings.ToLower(api)]
	if !ok {
		// return nil as stats for called API not found.
		// ideally it wont happen ever
		return nil
	}

	// Check the breach of throttle rules
	validAllowedCount := uint64(0) // set an invalid value
	for allowedAPI, allowedCount := range s3AllowedReqCountMap {
		// rules can have exact API names or patterns like `Get*`
		// first check exact match for API name, if not found then check wildcard
		if strings.EqualFold(allowedAPI, api) {
			validAllowedCount = allowedCount
			break // exact match found, break
		}
		if wildcard.MatchSimple(strings.ToLower(allowedAPI), strings.ToLower(api)) {
			validAllowedCount = allowedCount
			break // pattern match found, break
		}
	}
	if validAllowedCount > 0 && uint64(currReqCount) > validAllowedCount {
		return BucketThrottleQuotaExceeded{
			Bucket: bucket,
			Err:    fmt.Errorf("no of requests exceeded the allowed upper limit %d", validAllowedCount),
		}
	}

	return nil
}
