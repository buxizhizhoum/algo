package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	for i:=1;i<len(prices);i++{
		// 昨天的最大利润，加上今天的最大利润
		dp[i] = dp[i-1] + max(0, prices[i] - prices[i-1])
	}
	res := 0
	for _, tmp := range dp {
		res = max(res, tmp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

