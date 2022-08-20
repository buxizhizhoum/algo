package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	// 求最小的，设置一个不可能达到的长度做初始值
	res := len(nums) + 1
	var window int
	for ;right < len(nums); right++ {
		sum += nums[right]
		window = right - left + 1
		if sum >= target {
			res = min(res, window)
		} else {
			continue
		}

		for ;left < right; {
			sum -= nums[left]
			left++
			window = right - left + 1

			if sum >= target {
				res = min(res, window)
			} else {
				break
			}
		}
	}
	// 找不到的情况
	if res == len(nums) + 1 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//target, nums := 7, []int{2,3,1,2,4,3}
	target, nums := 11, []int{1,2,3,4,5}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
