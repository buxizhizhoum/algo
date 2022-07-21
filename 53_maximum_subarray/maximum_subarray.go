package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i:= 1;i<len(nums);i++ {
		dp[i] = max(nums[i], dp[i-1] + nums[i])
	}
	// 最小整数
	res := ^int(^uint(0) >> 1)
	for i:=0;i<len(dp);i++{
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
