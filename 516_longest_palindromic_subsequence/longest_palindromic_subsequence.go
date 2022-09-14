package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i:= 0; i< n ; i++ {
		dp[i] = make([]int, n)
	}
	// 每个字符串相等的位置，初始化为1，对一个字符而言，最长回文字符串就是自己
	for i := 0;i < n ; i ++ {
		dp[i][i] = 1
	}
	//dp[i][j] = max(max(dp[i+1][j], dp[i][j-1]), dp[i-1][j-1])
	// 决定了从下往上，从左往右遍历, 最后一行对角线上方，已经算好了，从倒数第二行开始
	for i := n - 2; i >=0 ;i-- {
		for j := i+1; j < n; j++ {
			if s[i] == s[j] {
				// 如果两个字符串相等，则是[i+1, j-1]之间最长回文字符串长度+2
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// todo: dp[i][j] = max(max(dp[i+1][j], dp[i][j-1]), dp[i+1][j-1])
				//dp[i+1][j-1]一定是比dp[i][j-1], dp[i+1][j]要小的
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	longestPalindromeSubseq("bbbab")
}
