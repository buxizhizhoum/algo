package more_general_solution

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i:= 1;i<len(prices);i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])
	}

	return dp[len(prices) - 1][1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
