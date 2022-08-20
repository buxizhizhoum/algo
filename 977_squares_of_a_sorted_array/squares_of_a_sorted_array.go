package squares_of_a_sorted_array

// 使用双指针来做，用两个变量分别指向开头和结尾，然后比较，每次将绝对值较大的那个数的平方值先加入数组的末尾，
// 然后依次往前更新，最后得到的就是所求的顺序
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums) - 1
	res := make([]int, len(nums))
	for cur :=len(nums) - 1;cur >=0; cur-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare > rightSquare {
			res[cur] = leftSquare
			left++
		} else {
			res[cur] = rightSquare
			right--
		}
	}
	return res
}
