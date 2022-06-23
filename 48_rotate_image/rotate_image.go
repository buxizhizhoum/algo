package rotate_image

func rotate(matrix [][]int)  {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	for i:=0; i < m; i++ {
		for j:=i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i:=0; i < m; i++ {
		reverse(matrix[i])
	}
}

func reverse(nums []int) {
	left := 0
	right := len(nums) - 1
	for ; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
