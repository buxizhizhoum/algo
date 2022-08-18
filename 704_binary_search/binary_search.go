package main

import "fmt"

func search(nums []int, target int) int {
	// 如果用小于，区间最好是初始化为[0, len(nums)], 这样跳出循环后left = right = len(nums)
	// 已经搜索完了所有的数据
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
	return -1
}

func main() {
	//testNums, target := []int{-1,0,3,5,9,12}, 13
	testNums, target := []int{5}, -5
	res := search(testNums, target)
	fmt.Println(res)
}
