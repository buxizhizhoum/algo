package third_time

func longestCommonSubsequence(text1 string, text2 string) int {
	// 这里在第一行和第一列放了一个空字符，并认为dp[0][..] = 0, dp[..][0] = 0,
	// 这样在后面初始化和for循环的时候写代码容易一些
	dp := make([][]int, len(text1) + 1)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int, len(text2) + 1)
	}

	//dp[0][..] = 0
	// dp[..][0] = 0
	for i:=1;i<len(text1) + 1;i++{
		for j:=1;j<len(text2) + 1;j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}
