package main

func maxProfit(prices []int) int {
	// dp[i]定义为i天获得的最大收益
	dp := make([]int, len(prices))
	dp[0] = 0
	for i:= 1;i<len(prices);i++ {
		// 保持不动 or 前一天不卖，今天卖
		dp[i] = max(0, dp[i-1] - prices[i-1] + prices[i])
	}
	res := 0
	for i:= 0;i<len(dp);i++{
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
