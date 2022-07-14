package next_greater_element

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	stack := make([]int, 0)
	for i:= 2*length - 1; i>=0;i-- {
		for ;len(stack) > 0 && stack[len(stack) - 1] <= nums[i%length]; {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) > 0 {
			// 等到了开始的部分，虚拟出来的后半部分，实际上会被覆盖一遍，栈的数据却积累了，这里就解决了最后一位的循环问题
			res[i%length] = stack[len(stack) - 1]
		} else {
			res[i%length] = -1
		}
		stack = append(stack, nums[i%length])
	}
	return res
}
