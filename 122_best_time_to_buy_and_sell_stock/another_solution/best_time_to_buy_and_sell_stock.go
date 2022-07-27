package another_solution

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	for i:=1;i<len(prices);i++{
		// 昨天的最大利润，或者昨天不卖，今天卖的最大利润
		dp[i] = max(dp[i-1], dp[i-1] + prices[i] - prices[i-1])
		// 昨天的最大利润，加上今天的最大利润
		// todo: 对比下714
		//dp[i] = dp[i-1] + max(0, prices[i] - prices[i-1])
	}

	return dp[len(prices) -1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
