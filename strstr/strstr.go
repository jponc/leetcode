package leetcode

// dynamic programming implementation
// O(m+n); where m = haystack length, n = needle length
// Unfortunately I can't run this in leetcode because memory allocation isn't enough :/
func strStrDynamicProgramming(haystack string, needle string) int {
	needle_length := len(needle)
	haystack_length := len(haystack)

	if needle_length == 0 {
		return 0
	}

	if needle_length > haystack_length {
		return -1
	}

	var val uint16

	// Added padding
	a := make([][]uint16, haystack_length+1)
	for i := range a {
		a[i] = make([]uint16, needle_length+1)
	}

	for i := 1; i <= haystack_length; i++ {
		for j := 1; j <= needle_length; j++ {
			if haystack[i-1] == needle[j-1] {
				val = a[i-1][j-1] + 1
				a[i][j] = val

				if val == uint16(needle_length) {
					return i - needle_length
				}
			}
		}
	}

	return -1
}

// iterative in place solution
// O(m*n); where m = haystack length, n = needle length
func strStr(haystack string, needle string) int {
	needleLength := len(needle)
	haystackLength := len(haystack)
	var tempIdx int

	if needleLength == 0 {
		return 0
	}

	if needleLength > haystackLength {
		return -1
	}

	for i := 0; i < haystackLength; i++ {
		for j := 0; j < needleLength; j++ {
			tempIdx = i + j

			if tempIdx > haystackLength-1 {
				return -1
			}

			if haystack[i+j] != needle[j] {
				break
			}

			if j+1 == needleLength {
				return i
			}
		}
	}

	return -1
}
