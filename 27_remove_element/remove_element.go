package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) < 1{
		return 0
	}
	slow := 0
	fast := 0
	for ;fast < len(nums); {
		if nums[fast] != val {
			// 这里直接不用判断也行，只是会存在冗余赋值
			if fast != slow {
				nums[slow] = nums[fast]
			}
			fast += 1
			slow += 1
		} else {
			// 跳过相等的
			fast += 1
		}
	}
	return slow
}
