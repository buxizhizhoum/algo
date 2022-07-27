package main

import "fmt"

func maxProfit(prices []int) int {
	return maxProfitGeneral(prices, 2)
}

func maxProfitGeneral(prices []int, k int) int {
	// dp0[i][j]表示不持有股票时，第i天，已有交易次数j时的最大收益
	// dp1[i][j]表示持有股票时，第i天，交易次数j时的最大收益
	dp0 := make([][]int, len(prices))
	dp1 := make([][]int, len(prices))
	for i:=0;i<len(prices);i++{
		dp0[i] = make([]int, k+1)
		dp1[i] = make([]int, k+1)
	}
	//dp0[0][j] = 0
	//dp0[i][0] = 0
	//dp1[0][j] = -prices[i]
	// todo: 为什么是这两个的max,
	// 已经交易了0次且持有股票，要么是前一天保持过来，要么一直没有买，买今天的，看哪个便宜买哪个，
	// 这里必然是负数，因为限制了是第一次买入，所以max取花钱最少的那个
	//dp1[i][0] = max(dp[i-1][0], -prices[i])
	for j:= 0;j<k+1;j++{
		dp0[0][j] = 0
		dp1[0][j] = -prices[0]
	}
	for i:=1;i<len(prices);i++{
		dp0[i][0] = 0
		dp1[i][0] = max(dp1[i-1][0], -prices[i])
	}
	//fmt.Println(dp0)
	//fmt.Println(dp1)

	for i:=1;i<len(prices);i++{
		for j:=1;j<k+1;j++{
			// 保持 or 卖出
			dp0[i][j] = max(dp0[i-1][j], dp1[i-1][j] + prices[i])
			// 保持 or 买入
			// todo: 注意交易次数扣减，应该是在买入时候进行
			dp1[i][j] = max(dp1[i-1][j], dp0[i-1][j-1] - prices[i])
		}
	}
	//fmt.Println(dp0)
	//fmt.Println(dp1)
	return dp0[len(prices)-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//maxProfit([]int{3,3,5,0,0,3,1,4})
	res := maxProfit([]int{1,2,3,4,5})
	fmt.Println(res)
}
