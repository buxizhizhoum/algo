package main

func lengthOfLIS(nums []int) int {
	// todo: 另一种思路，dp[i]代表的含义不一样
	// dp代表以nums[i]开头的最大递增子序列
	dp := make([]int, len(nums))
	for i:=0;i<len(dp);i++{
		dp[i] = 1
	}
	for i:=len(dp) - 1; i>=0; i--{
		for j:=len(dp) - 1; j>i; j--{
			// 如果开头的小于了子问题，有可能会有更长的序列
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	res := 1
	for i :=0; i< len(dp);i++{
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
