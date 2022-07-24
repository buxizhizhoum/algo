package main

import "fmt"

//https://blog.csdn.net/ylyg050518/article/details/75216437
func rob(nums []int) int {
	// dp[i]表示在0-i之间，可以抢到的组多的钱
	n := len(nums)
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if n == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	for i:= 2;i<n;i++{
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := rob([]int{1,2,3,1})
	fmt.Println(res)
}
