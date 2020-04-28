package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	output := l1
	temp := l2

	if l2.Val < l1.Val {
		output = l2
		temp = l1
	}

	curOutput := output
	prevOutput := output

	for curOutput != nil && temp != nil {
		if curOutput.Val <= temp.Val {
			prevOutput = curOutput
			curOutput = curOutput.Next
		} else {

			for temp != nil && temp.Val <= curOutput.Val {
				prevOutput.Next = temp
				prevOutput = prevOutput.Next
				temp = temp.Next
			}

			prevOutput.Next = curOutput
		}
	}

	if temp != nil {
		prevOutput.Next = temp
	}

	return output
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	mergeTwoLists(l1, l2)
}
