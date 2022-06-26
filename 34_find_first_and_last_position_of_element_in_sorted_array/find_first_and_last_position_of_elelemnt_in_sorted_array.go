package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right - left) / 2
		if nums[mid] == target {
			// 收缩右边界
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right - left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
