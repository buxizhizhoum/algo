package remove_element

func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	for ;fast < len(nums); {
		// fast和目标不相等，slow是需要进行替换的元素
		if nums[fast] != val {
			nums[slow] = nums[fast]
			fast += 1
			slow += 1
		} else {
			// 往前走，找不相等的元素，这个时候，slow停留在需要替换的元素位置，
			// slow, fast中间的就是可以替换的部分，也是重复元素的长度
			fast += 1
		}
	}
	return slow
}
