package leetcode

import (
	"math"
)

// https://leetcode.com/problems/reverse-integer/
// Given a 32-bit signed integer, reverse digits of an integer.

func ReverseInteger(x int) int {
	values := make([]int, 10)
	output := 0
	idx := 0

	for x/10 != 0 {
		toAdd := x % 10
		values[idx] = toAdd

		x = x / 10
		idx++
	}

	values[idx] = x

	for y := 0; y <= idx; y++ {
		output = output + (values[y] * int(math.Pow(10, float64(idx-y))))
	}

	if output <= -2147483648 || output >= 2147483647 {
		return 0
	}

	return output
}
