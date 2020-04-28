package leetcode

import (
	"testing"
)

func TestRemoveDuplicaes(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := removeDuplicates(input)

	if output != 5 {
		t.Errorf("Expected: %d, Output: %d", 5, output)
	}

	if input[0] != 0 {
		t.Errorf("Error!")
	}

	if input[1] != 1 {
		t.Errorf("Error!")
	}

	if input[2] != 2 {
		t.Errorf("Error!")
	}

	if input[3] != 3 {
		t.Errorf("Error!")
	}

	if input[4] != 4 {

		t.Errorf("Error!")
	}
}
