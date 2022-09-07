package integer_break

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3;i<=n;i++ {
		for j:=1;j<i-1;j++ {
			// 可以想 dp[i]最大乘积是怎么得到的呢？
			//其实可以从1遍历j，然后有两种渠道得到dp[i].
			//一个是j * (i - j) 直接相乘。
			//一个是j * dp[i - j]，相当于是拆分(i - j)，对这个拆分不理解的话，可以回想dp数组的定义。

			//j是从1开始遍历，拆分j的情况，在遍历j的过程中其实都计算过了。那么从1遍历j，比较(i - j) * j和dp[i - j] * j 取最大的。递推公式：dp[i] = max(dp[i], max((i - j) * j, dp[i - j] * j));
			//也可以这么理解，j * (i - j) 是单纯的把整数拆分为两个数相乘，而j * dp[i - j]是拆分成两个以及两个以上的个数相乘。
			//如果定义dp[i - j] * dp[j] 也是默认将一个数强制拆成4份以及4份以上了。
			//所以递推公式：dp[i] = max({dp[i], (i - j) * j, dp[i - j] * j});
			//那么在取最大值的时候，为什么还要比较dp[i]呢？
			//因为在递推公式推导的过程中，每次计算dp[i]，取最大的而已。
			dp[i] = max(dp[i], max((i - j) * j, dp[i-j] * j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}