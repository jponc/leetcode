package leetcode

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := map[string]bool{"()": true, "()[]{}": true, "(]": false, "([)]": false, "{[]}": true}
	for input, expectedOutput := range testCases {
		output := isValidParentheses(input)

		if output != expectedOutput {
			t.Errorf("Input: %s, Expected: %t, Output: %t", input, expectedOutput, output)
		}
	}
}
