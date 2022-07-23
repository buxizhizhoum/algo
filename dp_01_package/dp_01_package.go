package main

import "fmt"

func maximumValue(weight []int, value []int, w int, n int) int {
	//n: 物品数量
	//w: 背包容量
	//weight: 物品重量
	//value: 物品价值
	// 这里补上没有物品，或者背包没有容量这两种base case，dp数组多加一行一列
	// 没有物品这一行，也可以不加，这样物品这一行初始化的时候，需要判断第一个物品和背包大小，
	// 第一个物品重量大于背包容量，初始华为0，反之，初始化为物品价值
	//dp[i][j]前i个物品，容量为j的最大价值
	dp := make([][]int, n+1)
	for i:=0;i<n+1;i++ {
		dp[i] = make([]int, w+1)
	}
	for i:= 0; i< n+1;i++{
		dp[i][0] = 0
	}
	for j:= 0;j< w+1;j++{
		dp[0][j] = 0
	}

	for i:= 1; i< n+1; i++ {
		// j表示容量
		for j:= 1; j< w+1;j++ {
			// 注意这里因为加了一行没物品的初始化条件，dp数组的i索引和weight数组差1
			// 如果当前物品装不下
			if j-weight[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 第二个维度表示容量
				// dp[i-1][j]表示这件物品不取
				// dp[i-1][j-weight[j]]表示这件物品取
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i-1]] + value[i-1])
			}
		}
	}
	//fmt.Println(dp)
	return dp[n][w]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	N := 3
	W := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	res := maximumValue(wt, val, W, N)
	fmt.Println(res)
}
