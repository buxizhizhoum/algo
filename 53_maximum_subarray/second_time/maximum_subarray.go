package second_time

func maxSubArray(nums []int) int {
	//dp[i]：包括下标i之前的最大连续子序列和为dp[i]。
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i:=1;i<len(nums);i++{
		dp[i] = max(nums[i], dp[i-1] + nums[i])
	}

	res := dp[0]
	for i:=0;i<len(dp);i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
