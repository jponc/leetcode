package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	output := addTwoNumbers(l1, l2)

	if output.Val != 7 {
		t.Errorf("Expected: %d, Output: %d", 7, output.Val)
	}

	output = output.Next
	if output.Val != 0 {
		t.Errorf("Expected: %d, Output: %d", 0, output.Val)
	}

	output = output.Next
	if output.Val != 8 {
		t.Errorf("Expected: %d, Output: %d", 8, output.Val)
	}
}
