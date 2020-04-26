package leetcode

type Tuple struct {
	first  byte
	second byte
}

func romanToInt(s string) int {
	lastIdx := len(s) - 1
	valuesMapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	edgesMapping := map[Tuple]int{
		Tuple{'I', 'V'}: 4,
		Tuple{'I', 'X'}: 9,
		Tuple{'X', 'L'}: 40,
		Tuple{'X', 'C'}: 90,
		Tuple{'C', 'D'}: 400,
		Tuple{'C', 'M'}: 900,
	}

	total := 0

	for i := 0; i <= lastIdx; i++ {
		curChar := s[i]
		val := valuesMapping[curChar]

		if i != lastIdx {
			nextChar := s[i+1]

			if edgeVal, ok := edgesMapping[Tuple{curChar, nextChar}]; ok {
				val = edgeVal
				i++
			}
		}

		total += val
	}

	return total
}
