package leetcode

func lengthOfLongestSubstring(s string) int {
	max := 0
	i := 0
	j := 0
	lookup := make(map[byte]bool)

	for j < len(s) {
		if _, found := lookup[s[j]]; !found {
			lookup[s[j]] = true
			j++

			if len(lookup) > max {
				max = len(lookup)
			}
		} else {
			delete(lookup, s[i])
			i++
		}

	}

	return max
}
