// Copyright 2018 Weborama. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uint128_test

import (
	"math"
	"testing"

	"github.com/weborama/uint128"
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
			uint128.Uint128{H: 0x0, L: 0x456}, 11, 117, 5, 1,
			uint128.Uint128{H: 0x6a20000000000000, L: 0x0},
			uint128.Uint128{H: 0x5604000000000000, L: 0x0},
			uint128.Uint128{H: 0x0, L: 0x456}, uint128.Uint128{H: 0x0, L: 0x456},
		},
		{
			uint128.Uint128{H: 0x1, L: 0x456}, 65, 63, 6, 1,
			uint128.Uint128{H: 0x6a20000000000000, L: 0x8000000000000000},
			uint128.Uint128{H: 0x5604000000000000, L: 0x100000000000000},
			uint128.Uint128{H: 0x1, L: 0x456}, uint128.Uint128{H: 0x1, L: 0x456},
		},
	}

	// NOTE: binary representation: fmt.Sprintf("%0b%064b", input.H, input.L)

	for _, testCase := range testCases {
		var rval int

		var ruint uint128.Uint128

		testCase := testCase
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

func TestAdd128(t *testing.T) {
	testcases := []struct {
		label         string
		x, y, carry   uint128.Uint128
		sum, carryOut uint128.Uint128
	}{
		{
			label: "zero sum should result zero",
			x:     uint128.Zero(),
			y:     uint128.Zero(),
			sum:   uint128.Zero(),
		},
		{
			label: "zero+1 sum should result 1",
			x:     uint128.Zero(),
			y:     uint128.Uint128{H: 0, L: 1},
			sum:   uint128.Uint128{H: 0, L: 1},
		},
		{
			label: "1+ maxuint64 sum should overflow from Low to High",
			x:     uint128.Uint128{H: 0, L: 1},
			y:     uint128.Uint128{H: 0, L: math.MaxUint64},
			sum:   uint128.Uint128{H: 1, L: 0},
		},
		{
			label:    "1+ maxuint128 sum should overflow",
			x:        uint128.Uint128{H: 0, L: 1},
			y:        uint128.MaxUint128(),
			sum:      uint128.Uint128{H: 0, L: 0},
			carryOut: uint128.Uint128{H: 0, L: 1},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.label, func(t *testing.T) {
			sum, carryOut := uint128.Add128(tc.x, tc.y, tc.carry)

			assert(t, sum, tc.sum, "unexpected sum")
			assert(t, carryOut, tc.carryOut, "unexpected carryOut")
		})
	}
}

func TestAdd(t *testing.T) {
	testcases := []struct {
		label string
		x, y  uint128.Uint128
		sum   uint128.Uint128
	}{
		{
			label: "zero sum should result zero",
			x:     uint128.Zero(),
			y:     uint128.Zero(),
			sum:   uint128.Zero(),
		},
		{
			label: "zero+1 sum should result 1",
			x:     uint128.Zero(),
			y:     uint128.Uint128{H: 0, L: 1},
			sum:   uint128.Uint128{H: 0, L: 1},
		},
		{
			label: "1+ maxuint64 sum should overflow from Low to High",
			x:     uint128.Uint128{H: 0, L: 1},
			y:     uint128.Uint128{H: 0, L: math.MaxUint64},
			sum:   uint128.Uint128{H: 1, L: 0},
		},
		{
			label: "1+ maxuint128 sum should overflow",
			x:     uint128.Uint128{H: 0, L: 1},
			y:     uint128.MaxUint128(),
			sum:   uint128.Uint128{H: 0, L: 0},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.label, func(t *testing.T) {
			sum := tc.x.Add(tc.y)

			assert(t, sum, tc.sum, "unexpected sum")
		})
	}
}

func TestIncr(t *testing.T) {
	testcases := []struct {
		label string
		x     uint128.Uint128
		incr  uint128.Uint128
	}{
		{
			label: "zero+1 sum should result 1",
			x:     uint128.Zero(),
			incr:  uint128.Uint128{H: 0, L: 1},
		},
		{
			label: "maxuint64+1 sum should overflow from Low to High",
			x:     uint128.Uint128{H: 0, L: math.MaxUint64},
			incr:  uint128.Uint128{H: 1, L: 0},
		},
		{
			label: "maxuint128+1 sum should overflow",
			x:     uint128.MaxUint128(),
			incr:  uint128.Uint128{H: 0, L: 0},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.label, func(t *testing.T) {
			incr := tc.x.Incr()

			assert(t, incr, tc.incr, "unexpected incr")
		})
	}
}

func assert(t *testing.T, got, expected uint128.Uint128, message string) {
	if got != expected {
		t.Errorf("error: %s (got: %v, expected: %v)", message, got, expected)
	}
}

// func TestFormat(t *testing.T) {
// 	testCases := []struct {
// 		input uint128.Uint128
// 	}{
// 		{
// 			uint128.Uint128{H: 0x0, L: 0x456},
// 		},
// 		{
// 			uint128.Uint128{H: 0x1, L: 0x456},
// 		},
// 	}
//
// 	for _, testCase := range testCases {
// 		t.Run(testCase.input.String(), func(t *testing.T) {
// 			t.Logf("%%v: \t%v", testCase.input)
// 			t.Logf("%%+v:\t%+v", testCase.input)
// 			t.Logf("%%#v:\t%#v", testCase.input)
// 			t.Logf("%%T:\t%T", testCase.input)
// 			t.Logf("%%b:\t%b", testCase.input)
// 			t.Logf("%%d:\t%d", testCase.input)
// 			t.Logf("%%o:\t%o", testCase.input)
// 			t.Logf("%%x:\t%x", testCase.input)
// 			t.Logf("%%X:\t%X", testCase.input)
// 		})
// 	}
// }
