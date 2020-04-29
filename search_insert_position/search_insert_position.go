package leetcode

func searchInsert(nums []int, target int) int {
	var i int

	for i = 0; i < len(nums) && nums[i] < target; i++ {
		if nums[i] > target {
			return i - 1
		}
	}

	return i
}
