package leetcode

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	output := strStr("hello", "ll")

	if output != 2 {
		t.Errorf("Expected: %d, Output: %d", 2, output)
	}

	output = strStr("aaaaa", "bba")

	if output != -1 {
		t.Errorf("Expected: %d, Output: %d", -1, output)
	}
}
