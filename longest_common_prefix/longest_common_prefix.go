package leetcode

type Matrix struct {
	values []string
}

func (m *Matrix) isSimilar(idx int) bool {
	set := make(map[byte]bool)
	for _, value := range m.values {

		if idx >= len(value) {
			return false
		}

		char := value[idx]
		set[char] = true
	}

	return len(set) == 1
}

func longestCommonPrefix(strs []string) string {
	m := Matrix{values: strs}
	output := ""
	similar := true

	if len(strs) == 0 {
		return output
	}

	firstWord := strs[0]

	for i := 0; similar; i++ {
		if similar = m.isSimilar(i); similar {
			output += string(firstWord[i])
		}
	}

	return output
}
