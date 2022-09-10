package second_time

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := 0;i<len(dp);i++{
		dp[i] = make([]int, amount + 1)
	}
	for i:= 0;i<len(dp);i++{
		for j:=0;j<amount + 1;j++{
			// 求最小值，初始化到一个不可能取到的最大值
			dp[i][j] = amount + 1
		}
	}

	//dp[0][..]
	for j:=0;j<amount + 1;j++{
		if j % coins[0] == 0 {
			dp[0][j] = j / coins[0]
		}
	}
	// dp[..][0]
	for i:=0;i<len(coins);i++{
		dp[i][0] = 0
	}

	for i:=1;i<len(coins);i++{
		for j:=1;j<amount + 1;j++{
			if j - coins[i] >= 0 {
				// 分清coin的个数和路径的个数
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i]] + 1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[len(coins) - 1][amount] == amount + 1 {
		return -1
	}
	return dp[len(coins) - 1][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
