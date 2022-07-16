package main

func coinChange(coins []int, amount int) int {
	// dp数组的一个元素表示凑出这个金额的最优解
	dp := make([]int, amount + 1)
	for i:=0;i<len(dp);i++{
		// 这里只是个特殊的初始化技巧，把最大值初始化为了amount + 1，
		// 因为实际上不会有这个值，常见的技巧还有初始华为int.Max
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 穷举amount和coin的组合
	// i代表对amount的穷举
	for i:=1;i<amount + 1;i++{
		// 对于一个amount的取值i，所有coin都穷举一遍
		for _, coin := range coins {
			if i - coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1 + dp[i - coin])
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

