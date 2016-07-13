/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"path"
	"sort"
	"sync"
	"time"
)

const (
	// Erasure related constants.
	erasureAlgorithmKlauspost = "klauspost/reedsolomon/vandermonde"
	erasureAlgorithmISAL      = "isa-l/reedsolomon/cauchy"
)

// objectPartInfo Info of each part kept in the multipart metadata
// file after CompleteMultipartUpload() is called.
type objectPartInfo struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	ETag   string `json:"etag"`
	Size   int64  `json:"size"`
}

// byObjectPartNumber is a collection satisfying sort.Interface.
type byObjectPartNumber []objectPartInfo

func (t byObjectPartNumber) Len() int           { return len(t) }
func (t byObjectPartNumber) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byObjectPartNumber) Less(i, j int) bool { return t[i].Number < t[j].Number }

// checkSumInfo - carries checksums of individual scattered parts per disk.
type checkSumInfo struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
}

// erasureInfo - carries erasure coding related information, block
// distribution and checksums.
type erasureInfo struct {
	Algorithm    string         `json:"algorithm"`
	DataBlocks   int            `json:"data"`
	ParityBlocks int            `json:"parity"`
	BlockSize    int64          `json:"blockSize"`
	Index        int            `json:"index"`
	Distribution []int          `json:"distribution"`
	Checksum     []checkSumInfo `json:"checksum,omitempty"`
}

// IsValid - tells if the erasure info is sane by validating the data
// blocks, parity blocks and distribution.
func (e erasureInfo) IsValid() bool {
	return e.DataBlocks != 0 && e.ParityBlocks != 0 && len(e.Distribution) != 0
}

// pickValidErasureInfo - picks one valid erasure info content and returns, from a
// slice of erasure info content. If no value is found this function panics
// and dies.
func pickValidErasureInfo(eInfos []erasureInfo) erasureInfo {
	for _, eInfo := range eInfos {
		if eInfo.IsValid() {
			return eInfo
		}
	}
	panic("Unable to look for valid erasure info content")
}

// statInfo - carries stat information of the object.
type statInfo struct {
	Size    int64     `json:"size"`    // Size of the object `xl.json`.
	ModTime time.Time `json:"modTime"` // ModTime of the object `xl.json`.
}

// A xlMetaV1 represents `xl.json` metadata header.
type xlMetaV1 struct {
	Version string   `json:"version"` // Version of the current `xl.json`.
	Format  string   `json:"format"`  // Format of the current `xl.json`.
	Stat    statInfo `json:"stat"`    // Stat of the current object `xl.json`.
	// Erasure coded info for the current object `xl.json`.
	Erasure erasureInfo `json:"erasure"`
	// Minio release tag for current object `xl.json`.
	Minio struct {
		Release string `json:"release"`
	} `json:"minio"`
	// Metadata map for current object `xl.json`.
	Meta map[string]string `json:"meta"`
	// Captures all the individual object `xl.json`.
	Parts []objectPartInfo `json:"parts,omitempty"`
}

// newXLMetaV1 - initializes new xlMetaV1, adds version, allocates a fresh erasure info.
func newXLMetaV1(object string, dataBlocks, parityBlocks int) (xlMeta xlMetaV1) {
	xlMeta = xlMetaV1{}
	xlMeta.Version = "1.0.0"
	xlMeta.Format = "xl"
	xlMeta.Minio.Release = minioReleaseTag
	xlMeta.Erasure = erasureInfo{
		Algorithm:    erasureAlgorithmKlauspost,
		DataBlocks:   dataBlocks,
		ParityBlocks: parityBlocks,
		BlockSize:    blockSizeV1,
		Distribution: hashOrder(object, dataBlocks+parityBlocks),
	}
	return xlMeta
}

// IsValid - tells if the format is sane by validating the version
// string and format style.
func (m xlMetaV1) IsValid() bool {
	return m.Version == "1.0.0" && m.Format == "xl"
}

// ObjectPartIndex - returns the index of matching object part number.
func (m xlMetaV1) ObjectPartIndex(partNumber int) (index int) {
	for i, part := range m.Parts {
		if partNumber == part.Number {
			index = i
			return index
		}
	}
	return -1
}

// AddObjectPart - add a new object part in order.
func (m *xlMetaV1) AddObjectPart(partNumber int, partName string, partETag string, partSize int64) {
	partInfo := objectPartInfo{
		Number: partNumber,
		Name:   partName,
		ETag:   partETag,
		Size:   partSize,
	}

	// Update part info if it already exists.
	for i, part := range m.Parts {
		if partNumber == part.Number {
			m.Parts[i] = partInfo
			return
		}
	}

	// Proceed to include new part info.
	m.Parts = append(m.Parts, partInfo)

	// Parts in xlMeta should be in sorted order by part number.
	sort.Sort(byObjectPartNumber(m.Parts))
}

// ObjectToPartOffset - translate offset of an object to offset of its individual part.
func (m xlMetaV1) ObjectToPartOffset(offset int64) (partIndex int, partOffset int64, err error) {
	if offset == 0 {
		// Special case - if offset is 0, then partIndex and partOffset are always 0.
		return 0, 0, nil
	}
	partOffset = offset
	// Seek until object offset maps to a particular part offset.
	for i, part := range m.Parts {
		partIndex = i
		// Offset is smaller than size we have reached the proper part offset.
		if partOffset < part.Size {
			return partIndex, partOffset, nil
		}
		// Continue to towards the next part.
		partOffset -= part.Size
	}
	// Offset beyond the size of the object return InvalidRange.
	return 0, 0, InvalidRange{}
}

// pickValidXLMeta - picks one valid xlMeta content and returns from a
// slice of xlmeta content. If no value is found this function panics
// and dies.
func pickValidXLMeta(xlMetas []xlMetaV1) xlMetaV1 {
	for _, xlMeta := range xlMetas {
		if xlMeta.IsValid() {
			return xlMeta
		}
	}
	panic("Unable to look for valid XL metadata content")
}

// list of all errors that can be ignored in a metadata operation.
var objMetadataOpIgnoredErrs = []error{
	errDiskNotFound,
	errDiskAccessDenied,
	errFaultyDisk,
	errVolumeNotFound,
	errFileAccessDenied,
}

// readXLMetadata - returns the object metadata `xl.json` content from
// one of the disks picked at random.
func (xl xlObjects) readXLMetadata(bucket, object string) (xlMeta xlMetaV1, err error) {
	for _, disk := range xl.getLoadBalancedQuorumDisks() {
		if disk == nil {
			continue
		}
		xlMeta, err = readXLMeta(disk, bucket, object)
		if err == nil {
			return xlMeta, nil
		}
		// For any reason disk or bucket is not available continue
		// and read from other disks.
		if isErrIgnored(err, objMetadataOpIgnoredErrs) {
			continue
		}
		break
	}
	// Return error here.
	return xlMetaV1{}, err
}

// Undo rename xl metadata, renames successfully renamed `xl.json` back to source location.
func (xl xlObjects) undoRenameXLMetadata(srcBucket, srcPrefix, dstBucket, dstPrefix string, errs []error) {
	var wg = &sync.WaitGroup{}
	srcJSONFile := path.Join(srcPrefix, xlMetaJSONFile)
	dstJSONFile := path.Join(dstPrefix, xlMetaJSONFile)

	// Undo rename `xl.json` on disks where RenameFile succeeded.
	for index, disk := range xl.storageDisks {
		if disk == nil {
			continue
		}
		// Undo rename object in parallel.
		wg.Add(1)
		go func(index int, disk StorageAPI) {
			defer wg.Done()
			if errs[index] != nil {
				return
			}
			_ = disk.RenameFile(dstBucket, dstJSONFile, srcBucket, srcJSONFile)
		}(index, disk)
	}
	wg.Wait()
}

// deleteXLMetadata - deletes `xl.json` on a single disk.
func deleteXLMetdata(disk StorageAPI, bucket, prefix string) error {
	jsonFile := path.Join(prefix, xlMetaJSONFile)
	return disk.DeleteFile(bucket, jsonFile)
}

// writeXLMetadata - writes `xl.json` to a single disk.
func writeXLMetadata(disk StorageAPI, bucket, prefix string, xlMeta xlMetaV1) error {
	jsonFile := path.Join(prefix, xlMetaJSONFile)

	// Marshal json.
	metadataBytes, err := json.Marshal(&xlMeta)
	if err != nil {
		return err
	}
	// Persist marshalled data.
	return disk.AppendFile(bucket, jsonFile, metadataBytes)
}

// deleteAllXLMetadata - deletes all partially written `xl.json` depending on errs.
func deleteAllXLMetadata(disks []StorageAPI, bucket, prefix string, errs []error) {
	var wg = &sync.WaitGroup{}
	// Delete all the `xl.json` left over.
	for index, disk := range disks {
		if disk == nil {
			continue
		}
		// Undo rename object in parallel.
		wg.Add(1)
		go func(index int, disk StorageAPI) {
			defer wg.Done()
			if errs[index] != nil {
				return
			}
			_ = deleteXLMetdata(disk, bucket, prefix)
		}(index, disk)
	}
	wg.Wait()
}

// writeUniqueXLMetadata - writes unique `xl.json` content for each disk in order.
func writeUniqueXLMetadata(disks []StorageAPI, bucket, prefix string, xlMetas []xlMetaV1, writeQuorum, readQuorum int) error {
	var wg = &sync.WaitGroup{}
	var mErrs = make([]error, len(disks))

	// Start writing `xl.json` to all disks in parallel.
	for index, disk := range disks {
		if disk == nil {
			mErrs[index] = errDiskNotFound
			continue
		}
		wg.Add(1)
		// Write `xl.json` in a routine.
		go func(index int, disk StorageAPI) {
			defer wg.Done()

			// Pick one xlMeta for a disk at index.
			xlMetas[index].Erasure.Index = index + 1

			// Write unique `xl.json` for a disk at index.
			err := writeXLMetadata(disk, bucket, prefix, xlMetas[index])
			if err != nil {
				mErrs[index] = err
			}
		}(index, disk)
	}

	// Wait for all the routines.
	wg.Wait()

	// Do we have write quorum?.
	if !isDiskQuorum(mErrs, writeQuorum) {
		// Validate if we have read quorum.
		if isDiskQuorum(mErrs, readQuorum) {
			// Return success.
			return nil
		}
		// Delete all `xl.json` successfully renamed.
		deleteAllXLMetadata(disks, bucket, prefix, mErrs)
		return errXLWriteQuorum
	}

	// Reduce errors and verify quourm and return.
	if errCount, reducedErr := reduceErrs(mErrs); reducedErr != nil {
		if errCount < writeQuorum {
			// Delete all `xl.json` successfully renamed.
			deleteAllXLMetadata(disks, bucket, prefix, mErrs)
			return errXLWriteQuorum
		}
		if isErrIgnored(reducedErr, []error{
			errDiskNotFound,
			errDiskAccessDenied,
			errFaultyDisk,
			errVolumeNotFound,
		}) {
			// Success.
			return nil
		}
		return reducedErr
	}

	// Success.
	return nil
}

// writeSameXLMetadata - write `xl.json` on all disks in order.
func writeSameXLMetadata(disks []StorageAPI, bucket, prefix string, xlMeta xlMetaV1, writeQuorum, readQuorum int) error {
	var wg = &sync.WaitGroup{}
	var mErrs = make([]error, len(disks))

	// Start writing `xl.json` to all disks in parallel.
	for index, disk := range disks {
		if disk == nil {
			mErrs[index] = errDiskNotFound
			continue
		}
		wg.Add(1)
		// Write `xl.json` in a routine.
		go func(index int, disk StorageAPI, metadata xlMetaV1) {
			defer wg.Done()

			// Save the disk order index.
			metadata.Erasure.Index = index + 1

			// Write xl metadata.
			err := writeXLMetadata(disk, bucket, prefix, metadata)
			if err != nil {
				mErrs[index] = err
			}
		}(index, disk, xlMeta)
	}

	// Wait for all the routines.
	wg.Wait()

	// Do we have write Quorum?.
	if !isDiskQuorum(mErrs, writeQuorum) {
		// Do we have readQuorum?.
		if isDiskQuorum(mErrs, readQuorum) {
			// Return success.
			return nil
		}
		// Delete all `xl.json` successfully renamed.
		deleteAllXLMetadata(disks, bucket, prefix, mErrs)
		return errXLWriteQuorum
	}

	// Reduce errors and verify quourm and return.
	if errCount, reducedErr := reduceErrs(mErrs); reducedErr != nil {
		if errCount < writeQuorum {
			// Delete all `xl.json` successfully renamed.
			deleteAllXLMetadata(disks, bucket, prefix, mErrs)
			return errXLWriteQuorum
		}
		// Ignore specific errors if we are under write quorum.
		if isErrIgnored(reducedErr, []error{
			errDiskNotFound,
			errDiskAccessDenied,
			errFaultyDisk,
			errVolumeNotFound,
		}) {
			// Success.
			return nil
		}
		return reducedErr
	}

	// Success.
	return nil
}
