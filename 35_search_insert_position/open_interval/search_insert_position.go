package open_interval

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for ;left < right; {
		mid := left + (right - left) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	// 分别处理如下四种情况
	// 目标值在数组所有元素之前 [0,0)
	// 目标值等于数组中某一个元素 return middle
	// 目标值插入数组中的位置 [left, right) ，return right 即可
	// 目标值在数组所有元素之后的情况 [left, right)，因为是右开区间，所以 return right
	return right
}
// 0, 3, 5
// mid = 1 + (2 - 1) / 2
// [2, 2]
