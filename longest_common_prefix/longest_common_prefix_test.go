package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	output := longestCommonPrefix(input)

	if output != "fl" {
		t.Errorf("Wrong output: %s", output)
	}

	input = []string{"dog", "racecar", "car"}
	output = longestCommonPrefix(input)

	if output != "" {
		t.Error("Wrong output")
	}
}
