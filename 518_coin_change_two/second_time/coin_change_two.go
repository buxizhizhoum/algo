package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))
	for i := 0;i<len(dp);i++{
		dp[i] = make([]int, amount + 1)
	}
	// 背包容量为0的时候，不放任何东西就是一个解法，所以初始化为0
	dp[0][0] = 1
	// dp[0][..]
	for j := 1;j<amount + 1;j++{
		// 背包容量能装下这个硬币的时候，找到了一种组合
		if j % coins[0] == 0 {
			dp[0][j] = 1
		}
	}
	// dp[..][0] = 0, 这个不用初始化，在for循环中，让j从0开始即可


	for i:=1;i<len(coins);i++{
		for j:=0;j<amount + 1;j++ {
			if j - coins[i] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	//fmt.Println(dp)
	return dp[len(coins)-1][amount]
}

func main() {
	res := change(5, []int{1,2,5})
	//res := change(3, []int{2})
	fmt.Println(res)
}



