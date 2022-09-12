package third_time

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的lcs为dp[i][j]
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
