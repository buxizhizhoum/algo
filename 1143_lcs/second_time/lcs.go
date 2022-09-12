package second_time

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int, len(text2))
	}

	//dp[0][..]
	for j:=0;j<len(text2);j++ {
		if text2[j] == text1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	// dp[..][0]
	for i:=0;i<len(text1);i++{
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i:=1;i<len(text1);i++{
		for j:=1;j<len(text2);j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	res := 0
	for i:= 0;i<len(text1);i++{
		for j:= 0;j<len(text2);j++ {
			res = max(res, dp[i][j])
		}
	}
	fmt.Println(dp)
	return res
}

func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}
//  a c e
//a[1 1 1]
//b[1 1 1]
//c[1 2 2]
//d[1 2 2]
//e[1 2 3]
