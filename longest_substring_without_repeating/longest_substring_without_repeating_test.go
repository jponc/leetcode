package leetcode

import (
	"testing"
)

func TestLongestSubstringWithoutRepeating(t *testing.T) {
	testCases := map[string]int{"abcabcbb": 3, "bbbbb": 1, "pwwkew": 3, "dvdf": 3}

	for input, expectedOutput := range testCases {
		output := lengthOfLongestSubstring(input)

		if output != expectedOutput {
			t.Errorf("Input: %s, Expected: %d, Output: %d", input, expectedOutput, output)
		}
	}
}
