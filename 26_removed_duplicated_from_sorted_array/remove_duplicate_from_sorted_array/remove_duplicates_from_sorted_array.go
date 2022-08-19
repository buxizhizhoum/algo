package remove_duplicate_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	// fast找到第一个不重复的数字，slow是上一个数字，slow + 1到fast之间是需要替换的重复数字
	slow, fast := 0, 1
	for ;fast < len(nums); {
		if nums[fast] == nums[slow] {
			fast += 1
		} else {
			// 需要前进一步，留一个，否则把所有元素都替换了
			slow += 1
			nums[slow] = nums[fast]
			fast += 1
		}
	}
	return slow + 1
}
