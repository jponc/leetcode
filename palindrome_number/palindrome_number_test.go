package leetcode

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	testCases := map[int]bool{121: true, -123: false, 242: true}
	for input, expectedOutput := range testCases {
		output := isPalindrome(input)
		output2 := isPalindrome2(input)

		if output != expectedOutput {
			t.Errorf("Input: %d, Expected: %t, Output: %t", input, expectedOutput, output)
		}

		if output2 != expectedOutput {
			t.Errorf("Input: %d, Expected: %t, Output: %t", input, expectedOutput, output2)
		}
	}
}
