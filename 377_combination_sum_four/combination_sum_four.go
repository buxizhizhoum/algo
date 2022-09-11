package _77_combination_sum_four

import "sort"

// 如果dp初始化为[][], intlen(nums), target + 1大小, 没跑过
func combinationSum4(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dp := make([][]int, len(nums) + 1)
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]int, target + 1)
	}
	// dp[0][..], dp 是len(nums) + 1, dp[0][..]表示不使用任何元素
	for j:=0;j<target + 1;j++ {
		dp[0][j] = 0
	}
	// dp[..][0], 只有背包容量为0的时候，不装才是一种解法
	for i:=0;i<len(dp);i++ {
		dp[i][0] = 1
	}

	for i:=1;i<len(nums) + 1;i++ {
		for j:=1;j<target + 1;j++ {
			if j - nums[i-1] >= 0 {
				// todo: 为什么要有这个for循环
				for k:=0;k < i;k++ {
					// todo: 为什么这里不行
					//dp[i][j] = dp[i-1][j] + dp[i][j-nums[k]]
					dp[i][j] += dp[i][j-nums[k]]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][target]
}
