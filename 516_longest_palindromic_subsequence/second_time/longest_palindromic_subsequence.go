package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i:=0;i<len(s);i++ {
		dp[i] = make([]int, len(s))
	}

	// 对角线
	for i:= 0;i<len(s);i++ {
		dp[i][i] = 1
	}

	for i:=len(s) - 2;i>=0;i-- {
		for j:= i+1;j < len(s);j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[0][len(s) - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//[1 2 3 3 4]
	//[0 1 2 2 3]
	//[0 0 1 1 3]
	//[0 0 0 1 1]
	//[0 0 0 0 1]
	longestPalindromeSubseq("bbbab")
}
