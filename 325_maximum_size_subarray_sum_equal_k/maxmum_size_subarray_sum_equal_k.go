package main

import "fmt"

//https://juejin.cn/post/6972181565593878564
//func maxSubarrayLen(nums []int, k int) int {
//	cache := make(map[int]int, 0)
//	cache[0] = -1
//	res, curSum := 0, 0
//	for i, _ := range nums {
//		curSum += nums[i]
//		if index, ok := cache[k - curSum]; ok {
//			res = max(res, i - index)
//		} else if _, flag := cache[curSum]; !flag {
//			cache[curSum] = i
//		}
//	}
//	return res
//}


func maxSubarrayLen(nums []int, k int) int {
	preSum := make([]int, len(nums) + 1)
	for i:=1;i<len(nums) + 1;i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	cache := make(map[int]int, 0)
	for i, val := range preSum {
		// 这里要求数组的最长长度，前缀和里面只存放最开始出现的和和索引映射，这样减去的数组长度是最短的
		if _, ok := cache[val]; !ok {
			cache[val] = i
		}
	}

	res := 0
	for i, _ := range nums {
		curSum := preSum[i+1]
		needSum := curSum - k
		// 截止目前的前缀和是sum，看看有没有k-sum这个前缀和项目，如果有，两者相减就是数组长度
		if index, ok := cache[needSum]; ok {
			res = max(res, i - index + 1)
		}
	}
	return res
}

func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testNums, target := []int{1, -1, 5, -2, 3}, 3
	res := maxSubarrayLen(testNums, target)
	fmt.Println(res)
}
