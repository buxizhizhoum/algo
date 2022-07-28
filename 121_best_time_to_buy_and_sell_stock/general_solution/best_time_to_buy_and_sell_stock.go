package general_solution

// 买卖一次
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp0 := make([][]int, len(prices))
	dp1 := make([][]int, len(prices))
	for i:= 0;i<len(prices);i++{
		// k = 1
		dp0[i] = make([]int, 2)
		dp1[i] = make([]int, 2)
	}
	//dp0[0][i] = 0
	//dp0[i][0] = 0
	dp1[0][1] = -prices[0]
	dp1[1][1] = max(0, dp1[0][0] - prices[1] + prices[0])
	for i:= 1;i<len(prices);i++ {
		for j:=1;j<2;j++ {
			dp0[i][j] = max(dp0[i-1][j], dp1[i-1][j] + prices[i])
			dp1[i][j] = max(dp1[i-1][j], dp0[i-1][j-1] - prices[i])
		}
	}
	return dp0[len(prices) - 1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
