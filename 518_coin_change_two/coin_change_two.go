package coin_change_two

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n + 1)
	for i:=0;i<n+1;i++ {
		dp[i] = make([]int, amount + 1)
	}

	for i:= 0;i<n+1;i++{
		dp[i][0] = 1
	}
	for j:= 1;j<amount+1;j++{
		dp[0][j] = 0
	}

	for i:=1;i<n+1;i++{
		for j:=1;j<amount+1;j++{
			if j - coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				//dp[i-1][j]表示在不取第i个的时候的组合数
				//dp[i][j-coins[i-1]]表示在取i个的时候，那么就应该关注如何凑出金额 j - coins[i-1]
				// todo: 这里第一个索引是i，和01背包以及分割子集和里面的不一样，这里是完全背包
				dp[i][j] = dp[i-1][j] + dp[i][j - coins[i-1]]
			}
		}
	}
	return dp[n][amount]

}
