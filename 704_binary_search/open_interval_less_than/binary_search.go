package main

// 不推荐这种写法，如果用 < 则区间选[0, len(nums)]更好，否则当left = right跳出循环的时候
// 还有一个数字nums[left]没有在循环中搜索到，需要额外判断
func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for ;left < right; {
		mid := left + (right - left) / 2
		if nums[mid] == target{
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	// 终止条件left = right, left还没有判断
	if nums[left] == target {
		return left
	}
	return -1
}
