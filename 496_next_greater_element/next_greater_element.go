package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	greaterElement := nextGreater(nums2)
	fmt.Println(greaterElement)
	greaterElementMap := make(map[int]int, 0)
	// nums2的下一个最大值映射
	for i, _ := range nums2 {
		greaterElementMap[nums2[i]] = greaterElement[i]
	}
	for i, v := range nums1 {
		if greater, ok := greaterElementMap[v]; ok {
			res[i] = greater
		} else {
			res[i] = -1
		}
	}
	return res
}

// 获取一个数组下一个最大值
func nextGreater(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	for i:= len(nums) -1;i>=0;i--{
		// 会被nums[i]遮挡的元素，都不需要在栈里面
		for ; len(stack) > 0 && stack[len(stack) - 1] <= nums[i]; {
			stack = stack[:len(stack) - 1]
		}
		// 如果栈顶元素被遮挡不住，那么下一个能看到的元素就是栈顶元素
		if len(stack) > 0 && stack[len(stack) - 1] > nums[i] {
			res[i] = stack[len(stack) - 1]
		} else {
			res[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return res
}

func main() {
	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println(res)
}
