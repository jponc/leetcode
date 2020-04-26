package leetcode

type Stack struct {
	values []rune
}

func (s *Stack) Push(c rune) {
	s.values = append(s.values, c)
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (rune, bool) {
	if s.isEmpty() {
		return 0, false
	} else {
		index := len(s.values) - 1
		topVal := s.values[index]
		s.values = s.values[:index]

		return topVal, true
	}
}

func isValidParentheses(s string) bool {
	stack := Stack{values: []rune{}}
	inParens := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		if val, ok := inParens[char]; ok {
			stack.Push(val)
		} else {
			if topVal, ok := stack.Pop(); !ok || topVal != char {
				return false
			}

		}
	}

	return stack.isEmpty()
}
