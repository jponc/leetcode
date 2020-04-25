package leetcode

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	testCases := map[int]int{123: 321, -123: -321, 240: 42}
	for input, expectedOutput := range testCases {
		output := ReverseInteger(input)

		if output != expectedOutput {
			t.Errorf("Input: %d, Expected: %d, Output: %d", input, expectedOutput, output)
		}
	}
}
