package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// 定义：s1[0..i] 和 s2[0..j] 的最小编辑距离是 dp[i+1][j+1]
	// todo: 从后往前呢？
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m + 1)
	for i:= 0; i<m+1;i++ {
		tmp := make([]int, n + 1)
		dp[i] = tmp
	}

	for i := 0; i< m+1;i++ {
		dp[i][0] = i
	}
	for j:= 0;j < n + 1;j++ {
		dp[0][j] = j
	}

	for i:=1;i < m+1; i++ {
		for j:=1;j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minThree(
					dp[i-1][j] + 1,
					dp[i][j-1] + 1,
					dp[i-1][j-1] + 1,
				)
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a< b {
		return a
	}
	return b
}

func minThree(a, b, c int) int {
	return min(a, min(b, c))
}

func main() {
	//res := minDistance("horse", "ros")
	res := minDistance("", "")
	fmt.Println(res)
}
