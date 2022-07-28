package best_time_to_buy_and_sell_stock


func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	// 不持有股票，第i天的收益
	dp0 := make([]int, n)
	// 持有股票，第i天收益
	dp1 := make([]int, n)

	dp0[0] = 0
	dp1[0] = -prices[0]
	// 不持有股票，第1天，要么在第0天买了又卖了，今天冷却，收益0， 要么第0天买了，第1天卖，收益prices[1] - prices[0]
	dp0[1] = max(0, prices[1]- prices[0])
	// 持有股票，要么第0天持有股票，第1天不动，要么第0天不持有，第1天持有
	dp1[1] = max(-prices[0], -prices[1])
	for i:= 2;i<n;i++{
		dp0[i] = max(dp0[i-1], dp1[i-1] + prices[i])
		// dp0[i-1]相当于冷却的那一天
		dp1[i] = max(dp1[i-1], dp0[i-2] - prices[i])
	}
	return dp0[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}