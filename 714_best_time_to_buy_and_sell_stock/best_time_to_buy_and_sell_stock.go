package best_time_to_buy_and_sell_stock

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// todo: 这部分跑不过
	//dp := make([]int, n)
	//for i:=1;i<n;i++{
	//	dp[i] = max(dp[i-1], dp[i-1] - prices[i-1] + prices[i] - fee)
	//}
	//fmt.Println(dp)
	//return dp[n-1]
	dp0 := make([]int, n)
	dp1 := make([]int, n)
	// 这里统一在买入的时候扣减手续费，也可以在卖出的时候扣减，等价
	dp1[0] = -prices[0] - fee
	for i:= 1;i<n;i++ {
		// 保持 or 卖出
		// todo: 为什么不是max(dp0[i-1], dp1[i-1] + prices[i]-prices[i-1])
		dp0[i] = max(dp0[i-1], dp1[i-1] + prices[i])
		// 保持 or 买入
		// todo: 为什么不是max(dp1[i-1], dp0[i-1] - prices[i] + prices[i-1]- fee)
		dp1[i] = max(dp1[i-1], dp0[i-1] - prices[i] - fee)
	}
	return dp0[n-1]
}

func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}
