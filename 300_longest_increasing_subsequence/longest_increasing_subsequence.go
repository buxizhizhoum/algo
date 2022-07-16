package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	// dp代表以nums[i]结尾的最大递增子序列
	dp := make([]int, len(nums))
	for i:=0;i<len(dp);i++{
		dp[i] = 1
	}
	for i:=0;i<len(dp);i++{
		for j:=0;j<i;j++{
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	res := 1
	for i :=0; i< len(dp);i++{
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
