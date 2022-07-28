package general_solution

func maxProfit(prices []int) int {
	// 不持有，第i天的收益
	dp0 := make([]int, len(prices))
	// 持有，第i天收益
	dp1 := make([]int, len(prices))
	dp0[0] = 0
	// 持有股票，第0天收益
	dp1[0] = -prices[0]
	for i:= 1;i<len(prices);i++ {
		// 保持不动 or 卖出
		dp0[i] = max(dp0[i-1], dp1[i-1] + prices[i])
		// 保持不动 or 买入
		dp1[i] = max(dp1[i-1], dp0[i-1] - prices[i])
	}
	return dp0[len(prices) - 1]
}

func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}
