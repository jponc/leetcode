package leetcode

import "fmt"

// https://leetcode.com/problems/two-sum/
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

func TwoSum(nums []int, target int) []int {
	complementNums := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		fmt.Println(complementNums)
		num := nums[i]

		if complementIdx, ok := complementNums[num]; ok {
			return []int{complementIdx, i}
		} else {
			complementNums[target-num] = i
		}
	}

	return []int{}
}
