package leetcode

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	output := searchInsert([]int{1, 3, 5, 6}, 5)
	if output != 2 {
		t.Errorf("Expected: %d, Output: %d", 2, output)
	}

	output = searchInsert([]int{1, 3, 5, 6}, 2)
	if output != 1 {
		t.Errorf("Expected: %d, Output: %d", 1, output)
	}

	output = searchInsert([]int{1, 3, 5, 6}, 7)
	if output != 4 {
		t.Errorf("Expected: %d, Output: %d", 4, output)
	}

	output = searchInsert([]int{1, 3, 5, 6}, 0)
	if output != 0 {
		t.Errorf("Expected: %d, Output: %d", 0, output)
	}
}
