package main

import "fmt"

// 拿糖果，和house robber类似，相邻的不能拿
func maxCandy(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i:= 2;i<len(nums);i++{
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}
	return dp[len(nums) - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testNums := []int{2,5,1,1,9,100,99}
	res := maxCandy(testNums)
	fmt.Println(res)
}


