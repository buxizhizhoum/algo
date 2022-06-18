package move_zeros

// 27一起看

func moveZeroes(nums []int)  {
	slow := 0
	fast := 0
	for ; fast < len(nums); {
		if nums[fast] == 0 {
			fast += 1
		} else {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow += 1
			fast += 1
		}

	}

}
