package partition_euqal_subset_sum

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	// 数组和无法分隔
	if sum % 2 != 0 {
		return false
	}
	//dp[i][j]表示用前i个物品，能否凑满和为j的情况
	dp := make([][]bool, n+1)
	for i:=0;i<n+1;i++{
		dp[i] = make([]bool, sum/2+1)
	}
	// 没有空间的情况，任何数都不选，就可以满足
	for i:=0;i< n+1;i++{
		dp[i][0] = true
	}
	// 没有物品可选，装不满
	for j:=1;j< sum/2+1;j++{
		dp[0][j] = false
	}
	for i:= 1;i<n+1;i++{
		for j:=1;j< sum/2+1;j++{
			// dp[i-1][j]表示用前i-1个数凑满和为j是否可行
			// dp[i-1][j-nums[i-1]]表示在取第i个数的情况下，凑满和为j是否可行，
			// j-nums[i-1]表示取了第i个物品（i在dp数组从1开始，在nums从0开始，索引查2）前面要预留的背包容量
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	// fmt.Println(dp)
	return dp[n][sum/2]
}
