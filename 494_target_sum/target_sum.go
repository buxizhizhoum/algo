package target_sum

// https://leetcode.cn/problems/target-sum/solution/mu-biao-he-by-leetcode-solution-o0cp/
// x 是添加负号组成的负数
// (sum - x) - x = S => x = (S - sum) / 2
func findTargetSumWays(nums []int, target int) int {
	total := sum(nums)
	diff := total - target
	if diff % 2 != 0 {
		return 0
	}
	if target > total {
		return 0
	}

	dp := make([][]int, len(nums))
	for i :=0;i<len(dp);i++ {
		dp[i] = make([]int, diff + 1)
	}

	// dp[0][j]
	for j:=0;j<diff + 1;j++ {
		if j >= nums[0] {
			dp[0][j] = 1
		}
	}
	// dp[i][0] = 0, 已经是0 了，可以略过

	for i:= 1;i<len(nums);i++ {
		for j:=1;j< diff + 1;j++ {
			if j - nums[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// todo: 注意这里状态转移不一样，这里是加和
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	//fmt.Println(dp)
	return dp[len(nums) - 1][diff]
}

func sum(nums []int) int {
	res := 0
	for i:=0;i<len(nums);i++ {
		res += nums[i]
	}
	return res
}
