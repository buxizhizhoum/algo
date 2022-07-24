package main

import "fmt"
//这道题的本质相当于在一列数组中取出一个或多个不相邻数，使其和最大。
//那么对于这类求极值的问题首先考虑动态规划 Dynamic Programming 来解，
//维护一个一位数组 dp，其中 dp[i] 表示 [0, i] 区间可以抢夺的最大值，
//对当前i来说，有抢和不抢两种互斥的选择，不抢即为 dp[i-1]（等价于去掉 nums[i]
//只抢 [0, i-1] 区间最大值），抢即为 dp[i-2] + nums[i]（等价于去掉 nums[i-1]）。
//再举一个简单的例子来说明一下吧，比如说 nums为{3, 2, 1, 5}，
//那么来看 dp 数组应该是什么样的，首先 dp[0]=3 没啥疑问，再看 dp[1] 是多少呢，
//由于3比2大，所以抢第一个房子的3，当前房子的2不抢，则dp[1]=3，那么再来看 dp[2]，
//由于不能抢相邻的，所以可以用再前面的一个的 dp 值加上当前的房间值，
//和当前房间的前面一个 dp 值比较，取较大值当做当前 dp 值，
//这样就可以得到状态转移方程 dp[i] = max(num[i] + dp[i - 2], dp[i - 1]),
//且需要初始化 dp[0] 和 dp[1]，其中 dp[0] 即为 num[0]，dp[1]
//此时应该为 max(num[0], num[1])，代码如下：


//https://blog.csdn.net/ylyg050518/article/details/75216437
func rob(nums []int) int {
	// dp[i]表示在0-i之间，可以抢到的组多的钱
	n := len(nums)
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if n == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	for i:= 2;i<n;i++{
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := rob([]int{1,2,3,1})
	fmt.Println(res)
}
