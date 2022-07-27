package best_time_to_buy_and_sell_stock

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 当天买，当天卖，没有意义，当天买，次日卖，操作数最大n/2，超过相当于不限操作数
	if k > n / 2 {
		return maxProfitKInfinit(prices)
	}
	// dp0[i][j] 不持有股票，第i天，交易次数j的时候的最大利润
	// dp1[i][j] 持有股票，第i天，交易次数j的时候的最大利润
	dp0 := make([][]int, n)
	dp1 := make([][]int, n)
	for i:= 0;i<n;i++{
		dp0[i] = make([]int, k+1)
		dp1[i] = make([]int, k+1)
	}
	//dp0[0][j] = 0
	//dp0[i][0] = 0
	//dp1[0][j] = -prices[0]
	//// 可以在任何一天买
	//dp1[i][0] = max(dp[i-1], -prices[i])
	dp1[0][0] = 0
	for i:=1;i<n;i++{
		dp0[i][0] = 0
		dp1[i][0] = max(dp1[i-1][0], -prices[i])
	}
	for j:= 0;j<k+1;j++{
		dp0[0][j] = 0
		dp1[0][j] = -prices[0]
	}
	for i:= 1;i<n;i++{
		for j:= 1;j<k+1;j++{
			// 保持 or 卖出
			dp0[i][j] = max(dp0[i-1][j], dp1[i-1][j] + prices[i])
			// 保持 or 买入
			dp1[i][j] = max(dp1[i-1][j], dp0[i-1][j-1] - prices[i])
		}
	}
	return dp0[n-1][k]
}


func maxProfitKInfinit(prices []int) int {
	// dp[i]表示第i天的最大利润
	dp := make([]int, len(prices))
	dp[0] = 0
	for i:=1;i<len(prices);i++{
		dp[i] = max(dp[i-1], dp[i-1]-prices[i-1] + prices[i])
		//dp[i] = dp[i-1] + max(0, prices[i] - prices[i-1])
	}
	res := 0
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
