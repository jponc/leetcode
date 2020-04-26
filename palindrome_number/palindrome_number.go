package leetcode

import (
	"strconv"
)

type Stack struct {
	values []byte
}

func (s *Stack) push(newVal byte) {
	s.values = append(s.values, newVal)
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) pop() (byte, bool) {
	if s.isEmpty() {
		return 0, false
	} else {
		index := len(s.values) - 1
		topVal := s.values[index]
		s.values = s.values[:index]

		return topVal, true
	}
}

func isPalindrome(x int) bool {
	s := Stack{values: []byte{}}
	xStr := strconv.Itoa(x)
	charLength := len(xStr)

	isOdd := charLength%2 != 0
	midIdx := charLength / 2

	// Push all values til mid index
	for i := 0; i < midIdx; i++ {
		s.push(xStr[i])
	}

	// Handle odd numbers, increment mid index pointer
	if isOdd {
		midIdx++
	}

	// pop and compare
	for i := midIdx; i < charLength; i++ {
		curVal := xStr[i]

		if topVal, ok := s.pop(); !ok || topVal != curVal {
			return false
		}
	}

	return true
}

// iterative and index pointers solution
// this is faster and no need for additional structure, which is good for space complexity
func isPalindrome2(x int) bool {
	xStr := strconv.Itoa(x)
	i := 0
	j := len(xStr) - 1

	for i < j {
		if xStr[i] != xStr[j] {
			return false
		}

		i++
		j--
	}

	return true
}
