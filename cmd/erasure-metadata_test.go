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
	"strconv"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
)

const ActualSize = 1000

// Test FileInfo.AddObjectPart()
func TestAddObjectPart(t *testing.T) {
	testCases := []struct {
		partNum       int
		expectedIndex int
	}{
		{1, 0},
		{2, 1},
		{4, 2},
		{5, 3},
		{7, 4},
		// Insert part.
		{3, 2},
		// Replace existing part.
		{4, 3},
		// Missing part.
		{6, -1},
	}

	// Setup.
	fi := newFileInfo("test-object", 8, 8)
	fi.Erasure.Index = 1
	if !fi.IsValid() {
		t.Fatalf("unable to get xl meta")
	}

	// Test them.
	for _, testCase := range testCases {
		if testCase.expectedIndex > -1 {
			partNumString := strconv.Itoa(testCase.partNum)
			fi.AddObjectPart(testCase.partNum, "etag."+partNumString, int64(testCase.partNum+humanize.MiByte), ActualSize, UTCNow(), nil, nil, 0)
		}

		if index := objectPartIndex(fi.Parts, testCase.partNum); index != testCase.expectedIndex {
			t.Fatalf("%+v: expected = %d, got: %d", testCase, testCase.expectedIndex, index)
		}
	}
}

// Test objectPartIndex(). generates a sample FileInfo data and asserts
// the output of objectPartIndex() with the expected value.
func TestObjectPartIndex(t *testing.T) {
	testCases := []struct {
		partNum       int
		expectedIndex int
	}{
		{2, 1},
		{1, 0},
		{5, 3},
		{4, 2},
		{7, 4},
	}

	// Setup.
	fi := newFileInfo("test-object", 8, 8)
	fi.Erasure.Index = 1
	if !fi.IsValid() {
		t.Fatalf("unable to get xl meta")
	}

	// Add some parts for testing.
	for _, testCase := range testCases {
		partNumString := strconv.Itoa(testCase.partNum)
		fi.AddObjectPart(testCase.partNum, "etag."+partNumString, int64(testCase.partNum+humanize.MiByte), ActualSize, UTCNow(), nil, nil, 0)
	}

	// Add failure test case.
	testCases = append(testCases, struct {
		partNum       int
		expectedIndex int
	}{6, -1})

	// Test them.
	for _, testCase := range testCases {
		if index := objectPartIndex(fi.Parts, testCase.partNum); index != testCase.expectedIndex {
			t.Fatalf("%+v: expected = %d, got: %d", testCase, testCase.expectedIndex, index)
		}
	}
}

// Test FileInfo.ObjectToPartOffset().
func TestObjectToPartOffset(t *testing.T) {
	// Setup.
	fi := newFileInfo("test-object", 8, 8)
	fi.Erasure.Index = 1
	if !fi.IsValid() {
		t.Fatalf("unable to get xl meta")
	}

	// Add some parts for testing.
	// Total size of all parts is 5,242,899 bytes.
	for _, partNum := range []int{1, 2, 4, 5, 7} {
		partNumString := strconv.Itoa(partNum)
		fi.AddObjectPart(partNum, "etag."+partNumString, int64(partNum+humanize.MiByte), ActualSize, UTCNow(), nil, nil, 0)
	}

	testCases := []struct {
		offset         int64
		expectedIndex  int
		expectedOffset int64
		expectedErr    error
	}{
		{0, 0, 0, nil},
		{1 * humanize.MiByte, 0, 1 * humanize.MiByte, nil},
		{1 + humanize.MiByte, 1, 0, nil},
		{2 + humanize.MiByte, 1, 1, nil},
		// Its valid for zero sized object.
		{-1, 0, -1, nil},
		// Max fffset is always (size - 1).
		{(1 + 2 + 4 + 5 + 7) + (5 * humanize.MiByte) - 1, 4, 1048582, nil},
		// Error if offset is size.
		{(1 + 2 + 4 + 5 + 7) + (5 * humanize.MiByte), 0, 0, InvalidRange{}},
	}

	// Test them.
	for _, testCase := range testCases {
		index, offset, err := fi.ObjectToPartOffset(context.Background(), testCase.offset)
		if err != testCase.expectedErr {
			t.Fatalf("%+v: expected = %s, got: %s", testCase, testCase.expectedErr, err)
		}
		if index != testCase.expectedIndex {
			t.Fatalf("%+v: index: expected = %d, got: %d", testCase, testCase.expectedIndex, index)
		}
		if offset != testCase.expectedOffset {
			t.Fatalf("%+v: offset: expected = %d, got: %d", testCase, testCase.expectedOffset, offset)
		}
	}
}

func TestFindFileInfoInQuorum(t *testing.T) {
	getNFInfo := func(n int, quorum int, t int64, dataDir string, succModTimes []time.Time) []FileInfo {
		fi := newFileInfo("test", 8, 8)
		fi.AddObjectPart(1, "etag", 100, 100, UTCNow(), nil, nil, 0)
		fi.ModTime = time.Unix(t, 0)
		fi.DataDir = dataDir
		fis := make([]FileInfo, n)
		for i := range fis {
			fis[i] = fi
			fis[i].Erasure.Index = i + 1
			if succModTimes != nil {
				fis[i].SuccessorModTime = succModTimes[i]
				fis[i].IsLatest = succModTimes[i].IsZero()
			}
			quorum--
			if quorum == 0 {
				break
			}
		}
		return fis
	}

	commonSuccModTime := time.Date(2023, time.August, 25, 0, 0, 0, 0, time.UTC)
	succModTimesInQuorum := make([]time.Time, 16)
	succModTimesNoQuorum := make([]time.Time, 16)
	for i := 0; i < 16; i++ {
		if i < 4 {
			continue
		}
		succModTimesInQuorum[i] = commonSuccModTime
		if i < 9 {
			continue
		}
		succModTimesNoQuorum[i] = commonSuccModTime
	}
	tests := []struct {
		fis                 []FileInfo
		modTime             time.Time
		succmodTimes        []time.Time
		expectedErr         error
		expectedQuorum      int
		expectedSuccModTime time.Time
		expectedIsLatest    bool
	}{
		{
			fis:            getNFInfo(16, 16, 1603863445, "36a21454-a2ca-11eb-bbaa-93a81c686f21", nil),
			modTime:        time.Unix(1603863445, 0),
			expectedErr:    nil,
			expectedQuorum: 8,
		},
		{
			fis:            getNFInfo(16, 7, 1603863445, "36a21454-a2ca-11eb-bbaa-93a81c686f21", nil),
			modTime:        time.Unix(1603863445, 0),
			expectedErr:    errErasureReadQuorum,
			expectedQuorum: 8,
		},
		{
			fis:            getNFInfo(16, 16, 1603863445, "36a21454-a2ca-11eb-bbaa-93a81c686f21", nil),
			modTime:        time.Unix(1603863445, 0),
			expectedErr:    errErasureReadQuorum,
			expectedQuorum: 0,
		},
		{
			fis:                 getNFInfo(16, 16, 1603863445, "36a21454-a2ca-11eb-bbaa-93a81c686f21", succModTimesInQuorum),
			modTime:             time.Unix(1603863445, 0),
			expectedErr:         nil,
			expectedQuorum:      12,
			expectedSuccModTime: commonSuccModTime,
			expectedIsLatest:    false,
		},
		{
			fis:                 getNFInfo(16, 16, 1603863445, "36a21454-a2ca-11eb-bbaa-93a81c686f21", succModTimesNoQuorum),
			modTime:             time.Unix(1603863445, 0),
			expectedErr:         nil,
			expectedQuorum:      12,
			expectedSuccModTime: time.Time{},
			expectedIsLatest:    true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run("", func(t *testing.T) {
			fi, err := findFileInfoInQuorum(context.Background(), test.fis, test.modTime, "", test.expectedQuorum)
			if err != test.expectedErr {
				t.Errorf("Expected %s, got %s", test.expectedErr, err)
			}
			if test.succmodTimes != nil {
				if !test.expectedSuccModTime.Equal(fi.SuccessorModTime) {
					t.Errorf("Expected successor mod time to be %v but got %v", test.expectedSuccModTime, fi.SuccessorModTime)
				}
				if test.expectedIsLatest != fi.IsLatest {
					t.Errorf("Expected IsLatest to be %v but got %v", test.expectedIsLatest, fi.IsLatest)
				}
			}
		})
	}
}

func TestTransitionInfoEquals(t *testing.T) {
	inputs := []struct {
		tier            string
		remoteObjName   string
		remoteVersionID string
		status          string
	}{
		{
			tier:            "S3TIER-1",
			remoteObjName:   mustGetUUID(),
			remoteVersionID: mustGetUUID(),
			status:          "complete",
		},
		{
			tier:            "S3TIER-2",
			remoteObjName:   mustGetUUID(),
			remoteVersionID: mustGetUUID(),
			status:          "complete",
		},
	}

	var i uint
	for i = 0; i < 8; i++ {
		fi := FileInfo{
			TransitionTier:      inputs[0].tier,
			TransitionedObjName: inputs[0].remoteObjName,
			TransitionVersionID: inputs[0].remoteVersionID,
			TransitionStatus:    inputs[0].status,
		}
		ofi := fi
		if i&(1<<0) != 0 {
			ofi.TransitionTier = inputs[1].tier
		}
		if i&(1<<1) != 0 {
			ofi.TransitionedObjName = inputs[1].remoteObjName
		}
		if i&(1<<2) != 0 {
			ofi.TransitionVersionID = inputs[1].remoteVersionID
		}
		actual := fi.TransitionInfoEquals(ofi)
		if i == 0 && !actual {
			t.Fatalf("Test %d: Expected FileInfo's transition info to be equal: fi %v ofi %v", i, fi, ofi)
		}
		if i != 0 && actual {
			t.Fatalf("Test %d: Expected FileInfo's transition info to be inequal: fi %v ofi %v", i, fi, ofi)
		}
	}
	fi := FileInfo{
		TransitionTier:      inputs[0].tier,
		TransitionedObjName: inputs[0].remoteObjName,
		TransitionVersionID: inputs[0].remoteVersionID,
		TransitionStatus:    inputs[0].status,
	}
	ofi := FileInfo{}
	if fi.TransitionInfoEquals(ofi) {
		t.Fatalf("Expected to be inequal: fi %v ofi %v", fi, ofi)
	}
}

func TestSkipTierFreeVersion(t *testing.T) {
	fi := newFileInfo("object", 8, 8)
	fi.SetSkipTierFreeVersion()
	if ok := fi.SkipTierFreeVersion(); !ok {
		t.Fatal("Expected SkipTierFreeVersion to be set on FileInfo but wasn't")
	}
}
