package leetcode

func removeDuplicates(nums []int) int {
	lastIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[lastIndex] != nums[i] {
			lastIndex++
			nums[lastIndex] = nums[i]
		}
	}

	return lastIndex + 1
}
