package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	// 索引0处补了一个空字符含义
	dp := make([][]int, m + 1)
	for i:=0;i< m + 1;i++ {
		tmp := make([]int, n + 1)
		dp[i] = tmp
	}
	for i:=0;i< m+ 1;i++ {
		dp[i][0] = 0
	}
	for j:= 0;j < n+ 1;j++ {
		dp[0][j] = 0
	}
	for i:=1;i< m + 1;i++ {
		for j := 1;j< n+ 1;j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(max(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}

func main() {
	longestCommonSubsequence("abcde", "ace")
}
