package leetcode

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := map[string]int{"III": 3, "IV": 4, "IX": 9, "LVIII": 58, "MCMXCIV": 1994}
	for input, expectedOutput := range testCases {
		output := romanToInt(input)

		if output != expectedOutput {
			t.Errorf("Input: %s, Expected: %d, Output: %d", input, expectedOutput, output)
		}
	}
}
