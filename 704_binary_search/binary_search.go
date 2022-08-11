package main

import "fmt"

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
	// 一直找到了最右边的情况
	if left == len(nums) {
		return -1
	}
	// left < right区间[left, right)，区间最后nums[left]还没有判断
	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	//testNums, target := []int{-1,0,3,5,9,12}, 13
	testNums, target := []int{5}, -5
	res := search(testNums, target)
	fmt.Println(res)
}
