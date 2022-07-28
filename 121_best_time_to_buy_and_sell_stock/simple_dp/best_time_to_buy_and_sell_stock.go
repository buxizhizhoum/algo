package main

import "fmt"

func maxProfit(prices []int) int {
	// dp[i]定义为在第i天卖出时的最大获益
	dp := make([]int, len(prices))
	dp[0] = 0
	for i:= 1;i<len(prices);i++ {
		// 保持不动 or 前一天不卖，今天卖
		// 注意只能交易一次
		// 根据在哪一天买入我们可以分成两种情况：
		//
		// 在第i天买入在第i天卖出；
		// 在第i天前某一天买入在第i天卖出；
		// 对于第1种情况，即在同一天买入卖出，获益0；
		// 对于第2种情况，因为dp[i-1]代表在第i-1天卖出时的最大获益，
		// 如果在第i-1天不卖而是在第i天卖就是这种情况下的最大获益
		dp[i] = max(0, dp[i-1] - prices[i-1] + prices[i])
		// 下面这样不行
		//dp[i] = max(0, max(dp[i-1], dp[i-1] - prices[i-1] + prices[i]))
		// 可以交易无限次
		//dp[i] = dp[i-1] + max(0, prices[i] - prices[i-1])
	}
	fmt.Println(dp)
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

func main() {
	//测试用例:[7,1,5,3,6,4] 测试结果:7 期望结果:5 stdout: [0 0 4 4 7 7]
	res := maxProfit([]int{7,1,5,3,6,4})
	fmt.Println(res)
}
