package minimum_ascii_delete_sum_for_two_strings

func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, m + 1)
	for i := 0; i< m+ 1;i++ {
		dp[i] = make([]int, n + 1)
	}
	// todo: 注意初始化这里需要加起来
	for i := 1; i< m + 1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j:= 1; j < n+ 1; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	for i:= 1; i< m+ 1;i++ {
		for j := 1; j < n + 1; j++ {
			if s1[i-1] == s2[j -1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					min(dp[i-1][j] + int(s1[i-1]),
						dp[i][j-1] + int(s2[j-1])),
					dp[i-1][j-1] + int(s1[i-1]) + int(s2[j-1]))
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
