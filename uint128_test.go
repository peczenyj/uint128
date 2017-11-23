// Copyright 2017 Weborama. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uint128_test

import (
	"testing"

	"github.com/Weborama/uint128"
)

func TestUint128Operations(t *testing.T) {
	testCases := []struct {
		input                 uint128.Uint128
		expectedLen           int
		expectedLeadingZeros  int
		expectedOnesCount     int
		expectedTrailingZeros int
		expectedReverse       uint128.Uint128
		expectedReverseBytes  uint128.Uint128
		expectedIncr          uint128.Uint128
		expectedDecr          uint128.Uint128
	}{
		{
			uint128.Uint128{H: 0x0, L: 0x456},
			11,
			117,
			5,
			1,
			uint128.Uint128{H: 0x6a20000000000000, L: 0x0},
			uint128.Uint128{H: 0x5604000000000000, L: 0x0},
			uint128.Uint128{H: 0x0, L: 0x456},
			uint128.Uint128{H: 0x0, L: 0x456},
		},
		{
			uint128.Uint128{H: 0x1, L: 0x456},
			65,
			63,
			6,
			1,
			uint128.Uint128{H: 0x6a20000000000000, L: 0x8000000000000000},
			uint128.Uint128{H: 0x5604000000000000, L: 0x100000000000000},
			uint128.Uint128{H: 0x1, L: 0x456},
			uint128.Uint128{H: 0x1, L: 0x456},
		},
	}

	// NOTE: binary representation: fmt.Sprintf("%0b%064b", input.H, input.L)

	for _, testCase := range testCases {
		var rval int
		var ruint uint128.Uint128
		t.Run(testCase.input.String(), func(t *testing.T) {
			rval = uint128.Len(testCase.input)
			if rval != testCase.expectedLen {
				t.Fatalf("Len - Expected:%d Got:%d", testCase.expectedLen, rval)
			}
			rval = uint128.LeadingZeros(testCase.input)
			if rval != testCase.expectedLeadingZeros {
				t.Fatalf("LeadingZeros - Expected:%d Got:%d", testCase.expectedLeadingZeros, rval)
			}
			rval = uint128.OnesCount(testCase.input)
			if rval != testCase.expectedOnesCount {
				t.Fatalf("OnesCount - Expected:%d Got:%d", testCase.expectedOnesCount, rval)
			}
			rval = uint128.TrailingZeros(testCase.input)
			if rval != testCase.expectedTrailingZeros {
				t.Fatalf("TrailingZeros - Expected:%d Got:%d", testCase.expectedTrailingZeros, rval)
			}
			ruint = uint128.Reverse(testCase.input)
			if ruint != testCase.expectedReverse {
				t.Fatalf("Reverse - Expected:%v Got:%v", testCase.expectedReverse, ruint)
			}
			ruint = uint128.ReverseBytes(testCase.input)
			if ruint != testCase.expectedReverseBytes {
				t.Fatalf("ReverseBytes - Expected:%v Got:%v", testCase.expectedReverseBytes, ruint)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	testCases := []struct {
		input uint128.Uint128
	}{
		{
			uint128.Uint128{H: 0x0, L: 0x456},
		},
		{
			uint128.Uint128{H: 0x1, L: 0x456},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input.String(), func(t *testing.T) {
			t.Logf("%v", testCase.input)
			t.Logf("%+v", testCase.input)
			t.Logf("%#v", testCase.input)
			t.Logf("%T", testCase.input)
			t.Logf("%b", testCase.input)
			t.Logf("%d", testCase.input)
			t.Logf("%o", testCase.input)
			t.Logf("%x", testCase.input)
			t.Logf("%X", testCase.input)
		})
	}
}
