package removed_duplicated_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	fast := 0
	for ; fast < len(nums);{
		if nums[fast] == nums[slow] {
			fast += 1
		} else {
			slow += 1
			nums[slow] = nums[fast]
			fast += 1
		}
	}
	return slow + 1
}

