package main

import (
	"fmt"
	"sort"
)

func combinationSum4(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dp := make([][]int, len(nums))
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]int, target + 1)
	}
	// dp[0][..]
	for j:=0;j<target + 1;j++ {
		if j >= nums[0] {
			dp[0][j] = 1
		}
	}
	// dp[..][0], 只有背包容量为0的时候，不装才是一种解法
	for i:=0;i<len(dp);i++ {
		dp[i][0] = 1
	}

	for i:=1;i<len(nums);i++ {
		for j:=1;j<target + 1;j++ {
			if j - nums[i] >= 0 {
				for k:=0;k <= i; k++ {
					//dp[i][j] = dp[i-1][j] + dp[i][j-nums[k]]
					dp[i][j] += dp[i][j-nums[k]]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums) - 1][target]
}


