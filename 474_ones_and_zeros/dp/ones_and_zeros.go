package dp

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs))
	for i:=0;i<len(strs);i++ {
		dp[i] = make([][]int, m + 1)
		for j := 0;j<m+1;j++ {
			dp[i][j] = make([]int, n + 1)
		}
	}
	//dp[0][..][..]
	firstZeroCount, firstOneCount := count(strs[0])
	for j:=0;j<m+1;j++ {
		for k :=0;k<n+1;k++ {
			if j >= firstZeroCount && k >= firstOneCount {
				dp[0][j][k] = 1
			}
		}
	}
	// 这里不用初始化，背包容量为0的情况，直接在循环里面处理

	for i:=1;i<len(strs);i++ {
		zeroCount, oneCount := count(strs[i])
		// 这里直接从0开始，背包容量为0，可以在循环中进行处理
		for j:=0;j<m+1;j++ {
			for k:=0;k<n+1;k++ {
				// 能容纳下这个字符串
				if j - zeroCount >= 0 && k - oneCount >= 0 {
					// 不取和取这个字符串
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeroCount][k-oneCount] + 1)
					// 不能取这个字符串
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}
	return dp[len(strs) - 1][m][n]
}

func count(s string) (int, int) {
	if s == "" {
		return 0, 0
	}
	zeroCount, oneCount := 0, 0
	for i:=0;i<len(s);i++ {
		if s[i] == '0' {
			zeroCount++
		} else if s[i] == '1' {
			oneCount++
		}
	}
	return zeroCount, oneCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
