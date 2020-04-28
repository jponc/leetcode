package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	current := l1
	prev := l1

	for current != nil && l2 != nil {
		for current != nil && current.Val <= l2.Val {
			prev = current
			current = current.Next
		}

		prev.Next = l2
		prev = l2

		for l2 != nil && l2.Val <= current.Val {
			prev = l2
			l2 = l2.Next
		}
	}

	return l1
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	mergeTwoLists(l1, l2)
}
