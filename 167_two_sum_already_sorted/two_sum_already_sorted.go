package two_sum_already_sorted


func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	// 不重复
	for ; left < right; {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{-1, -1}
}
