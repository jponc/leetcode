package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getNum(l *ListNode) int {
	if l != nil {
		return l.Val
	} else {
		return 0
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := l1
	out := l1
	carry := 0
	sum := 0

	for l1 != nil || l2 != nil {
		sum = getNum(l1) + getNum(l2) + carry
		newVal := sum % 10
		carry = sum / 10

		if l1 != nil {
			l1.Val = newVal
			prev = l1
			l1 = l1.Next
		} else {
			prev.Next = &ListNode{Val: newVal}
			prev = prev.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		prev.Next = &ListNode{Val: carry}
	}

	return out
}
