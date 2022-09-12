package longest_continuous_increasing_subquences

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i:=0;i<len(dp);i++ {
		dp[i] = 1
	}

	for i:=1;i<len(nums);i++ {
		// 这道题是求连续上升子序列，和上升自序列不同之处在于连续，所以这道题只需要一层for循环
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}

	res := 0
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
