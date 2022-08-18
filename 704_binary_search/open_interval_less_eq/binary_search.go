package main

// 这里还是用开区间，但是right初始化为len(nums), 这样当不满足left < right的时候，已经搜索
// 了所有索引
func search(nums []int, target int) int {
	left, right := 0, len(nums)
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
	// 终止时left = right, = len(nums), 已经搜索完了所有的数据
	return -1
}
