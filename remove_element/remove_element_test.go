package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	input := []int{3, 2, 2, 3}
	output := removeElement(input, 3)

	if output != 2 {
		t.Errorf("Expected: %d, Output: %d", 2, output)
	}

	if input[0] != 2 || input[1] != 2 {
		t.Errorf("Error!")
	}
}
