package leetcode

func removeElement(nums []int, val int) int {
	lastIndex := 0

	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[lastIndex] = nums[i]
			lastIndex++
		}
	}

	return lastIndex
}
