/*
 * Minio Cloud Storage, (C) 2016, 2017 Minio, Inc.
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

package cmd

import (
	"bytes"
	"path/filepath"
	"testing"
	"time"
)

// validates functionality provided to find most common
// time occurrence from a list of time.
func TestCommonTime(t *testing.T) {
	// List of test cases for common modTime.
	testCases := []struct {
		times []time.Time
		time  time.Time
	}{
		{
			// 1. Tests common times when slice has varying time elements.
			[]time.Time{
				time.Unix(0, 1).UTC(),
				time.Unix(0, 2).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 2).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 1).UTC(),
			}, time.Unix(0, 3).UTC(),
		},
		{
			// 2. Tests common time obtained when all elements are equal.
			[]time.Time{
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
			}, time.Unix(0, 3).UTC(),
		},
		{
			// 3. Tests common time obtained when elements have a mixture
			// of sentinel values.
			[]time.Time{
				time.Unix(0, 3).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 2).UTC(),
				time.Unix(0, 1).UTC(),
				time.Unix(0, 3).UTC(),
				time.Unix(0, 4).UTC(),
				time.Unix(0, 3).UTC(),
				timeSentinel,
				timeSentinel,
				timeSentinel,
			}, time.Unix(0, 3).UTC(),
		},
	}

	// Tests all the testcases, and validates them against expected
	// common modtime. Tests fail if modtime does not match.
	for i, testCase := range testCases {
		// Obtain a common mod time from modTimes slice.
		ctime, _ := commonTime(testCase.times)
		if testCase.time != ctime {
			t.Fatalf("Test case %d, expect to pass but failed. Wanted modTime: %s, got modTime: %s\n", i+1, testCase.time, ctime)
		}
	}
}

// partsMetaFromModTimes - returns slice of modTimes given metadata of
// an object part.
func partsMetaFromModTimes(modTimes []time.Time, checkSums []string) []xlMetaV1 {
	var partsMetadata []xlMetaV1
	for i, modTime := range modTimes {
		partsMetadata = append(partsMetadata, xlMetaV1{
			Erasure: erasureInfo{
				Checksum: []checkSumInfo{{Hash: checkSums[i]}},
			},
			Stat: statInfo{
				ModTime: modTime,
			},
			Parts: []objectPartInfo{
				{
					Name: "part.1",
				},
			},
		})
	}
	return partsMetadata
}

// toPosix - fetches *posix object from StorageAPI.
func toPosix(disk StorageAPI) *posix {
	retryDisk, ok := disk.(*retryStorage)
	if !ok {
		return nil
	}
	pDisk, ok := retryDisk.remoteStorage.(*posix)
	if !ok {
		return nil
	}
	return pDisk

}

// TestListOnlineDisks - checks if listOnlineDisks and outDatedDisks
// are consistent with each other.
func TestListOnlineDisks(t *testing.T) {
	rootPath, err := newTestConfig(globalMinioDefaultRegion)
	if err != nil {
		t.Fatalf("Failed to initialize config - %v", err)
	}
	defer removeAll(rootPath)

	obj, disks, err := prepareXL()
	if err != nil {
		t.Fatalf("Prepare XL backend failed - %v", err)
	}
	defer removeRoots(disks)

	type tamperKind int
	const (
		noTamper    tamperKind = iota
		deletePart  tamperKind = iota
		corruptPart tamperKind = iota
	)
	threeNanoSecs := time.Unix(0, 3).UTC()
	fourNanoSecs := time.Unix(0, 4).UTC()
	modTimesThreeNone := []time.Time{
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		timeSentinel,
		timeSentinel,
		timeSentinel,
		timeSentinel,
		timeSentinel,
		timeSentinel,
		timeSentinel,
		timeSentinel,
		timeSentinel,
	}
	modTimesThreeFour := []time.Time{
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		threeNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
		fourNanoSecs,
	}
	testCases := []struct {
		modTimes       []time.Time
		expectedTime   time.Time
		errs           []error
		_tamperBackend tamperKind
	}{
		{
			modTimes:     modTimesThreeFour,
			expectedTime: fourNanoSecs,
			errs: []error{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
			_tamperBackend: noTamper,
		},
		{
			modTimes:     modTimesThreeNone,
			expectedTime: threeNanoSecs,
			errs: []error{
				// Disks that have a valid xl.json.
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				// Majority of disks don't have xl.json.
				errFileNotFound,
				errFileNotFound,
				errFileNotFound,
				errFileNotFound,
				errFileNotFound,
				errDiskAccessDenied,
				errDiskNotFound,
				errFileNotFound,
				errFileNotFound,
			},
			_tamperBackend: deletePart,
		},
		{
			modTimes:     modTimesThreeNone,
			expectedTime: threeNanoSecs,
			errs: []error{
				// Disks that have a valid xl.json.
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				// Majority of disks don't have xl.json.
				errFileNotFound,
				errFileNotFound,
				errFileNotFound,
				errFileNotFound,
				errFileNotFound,
				errDiskAccessDenied,
				errDiskNotFound,
				errFileNotFound,
				errFileNotFound,
			},
			_tamperBackend: corruptPart,
		},
	}

	bucket := "bucket"
	object := "object"
	data := bytes.Repeat([]byte("a"), 1024)
	blakeHash := newHash(blake2bAlgo)
	_, bErr := blakeHash.Write(data)
	if bErr != nil {
		t.Fatalf("Failed to compute blakeHash %v", bErr)
	}

	xlDisks := obj.(*xlObjects).storageDisks
	// Calling erasureCreateFile on a temporary path to compute
	// blake2b checksums.
	_, checkSums, err := erasureCreateFile(xlDisks, minioMetaTmpBucket, object, bytes.NewReader(data), true, 1024*1024, 8, 8, blake2bAlgo, 9)
	if err != nil {
		t.Fatalf("Unable to create a temporary file to compute blake2b sum %v", err)
	}

	for i, test := range testCases {
		// Prepare bucket/object backend for the tests below.

		// Cleanup from previous test.
		obj.DeleteObject(bucket, object)
		obj.DeleteBucket(bucket)

		err = obj.MakeBucket("bucket")
		if err != nil {
			t.Fatalf("Failed to make a bucket %v", err)
		}

		_, err = obj.PutObject(bucket, object, int64(len(data)), bytes.NewReader(data), nil, "")
		if err != nil {
			t.Fatalf("Failed to putObject %v", err)
		}

		tamperedIndex := -1
		switch test._tamperBackend {
		case deletePart:
			for index, err := range test.errs {
				if err != nil {
					continue
				}
				// Remove a part from a disk
				// which has a valid xl.json,
				// and check if that disk
				// appears in outDatedDisks.
				tamperedIndex = index
				dErr := xlDisks[index].DeleteFile(bucket, filepath.Join(object, "part.1"))
				if dErr != nil {
					t.Fatalf("Test %d: Failed to delete %s - %v", i+1,
						filepath.Join(object, "part.1"), dErr)
				}
				break
			}
		case corruptPart:
			for index, err := range test.errs {
				if err != nil {
					continue
				}
				// Corrupt a part from a disk
				// which has a valid xl.json,
				// and check if that disk
				// appears in outDatedDisks.
				tamperedIndex = index
				dErr := xlDisks[index].AppendFile(bucket, filepath.Join(object, "part.1"), []byte("corruption"))
				if dErr != nil {
					t.Fatalf("Test %d: Failed to append corrupting data at the end of file %s - %v",
						i+1, filepath.Join(object, "part.1"), dErr)
				}
				break
			}

		}

		partsMetadata := partsMetaFromModTimes(test.modTimes, checkSums)

		onlineDisks, modTime := listOnlineDisks(xlDisks, partsMetadata, test.errs)
		outdatedDisks := outDatedDisks(xlDisks, onlineDisks, partsMetadata, bucket, object)
		if modTime.Equal(timeSentinel) {
			t.Fatalf("Test %d: modTime should never be equal to timeSentinel, but found equal",
				i+1)
		}

		if test._tamperBackend != noTamper {
			if tamperedIndex != -1 && outdatedDisks[tamperedIndex] == nil {
				t.Fatalf("Test %d: disk (%v) with part.1 missing or is an outdated disk, but wasn't listed by outDatedDisks",
					i+1, onlineDisks[tamperedIndex])
			}

		}

		if !modTime.Equal(test.expectedTime) {
			t.Fatalf("Test %d: Expected modTime to be equal to %v but was found to be %v",
				i+1, test.expectedTime, modTime)
		}

		// Check if a disk is considered both online and outdated,
		// which is a contradiction, except if parts are missing.
		overlappingDisks := make(map[string]*posix)
		for _, onlineDisk := range onlineDisks {
			if onlineDisk == nil {
				continue
			}
			pDisk := toPosix(onlineDisk)
			overlappingDisks[pDisk.diskPath] = pDisk
		}

		for index, outdatedDisk := range outdatedDisks {
			// ignore the intentionally tampered disk,
			// this is expected to appear as outdated
			// disk, since it doesn't have all the parts.
			if index == tamperedIndex {
				continue
			}

			if outdatedDisk == nil {
				continue
			}

			pDisk := toPosix(outdatedDisk)
			if _, ok := overlappingDisks[pDisk.diskPath]; ok {
				t.Errorf("Test %d: Outdated disk %v was also detected as an online disk - %v %v",
					i+1, pDisk, onlineDisks, outdatedDisks)
			}
		}
	}
}
