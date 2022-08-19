package move_zeros

func moveZeroes(nums []int)  {
	slow, fast := 0, 0
	for ;fast < len(nums); {
		if nums[fast] == 0 {
			fast += 1
		} else {
			// 在开始段，fast和slow相等且非零的时候，直接执行交换对值没有影响
			// 这里需要考虑两个点，一个是这种，另一个是大体正常情况下的逻辑怎么写
			// 这里就是遇到0的时候，fast, slow开始分离
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow += 1
			fast += 1
		}
	}
}
