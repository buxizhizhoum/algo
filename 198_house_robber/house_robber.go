package house_robber

// todo: 另一个更好理解，看另一个
func rob(nums []int) int {
	// dp[i]表示从i这个点开始抢，能抢到的钱，dp[0]是最开始的点，实际编号从1开始，所以0就是开头
	// dp[n]是最后一个点，站在n往后看，已经没有可以抢的点了，因为不包含n，dp[n]=0
	dp := make([]int, len(nums) + 2)
	for i:= len(nums)-1;i>=0;i--{
		dp[i] = max(dp[i+1], dp[i+2] + nums[i])
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
