package main

import "fmt"

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
	// todo: 这里不用初始化，背包容量为0的情况，直接在循环里面处理
	//dp[..][0][..]
	var tmpZeroCount, tmpOneCount int
	for i := 1;i<len(strs);i++ {
		tmpZeroCount, tmpOneCount = count(strs[i])
		for k:=0;k<n+1;k++ {
			if tmpZeroCount == 0 && k >= tmpOneCount {
				dp[i][0][k] = max(dp[i-1][0][k], dp[i-1][0][k-tmpOneCount] + 1)
			} else {
				dp[i][0][k] = dp[i-1][0][k]
			}
		}
	}
	// 这里对背包容量为0的情况做了初始化，实际上这个是不需要进行初始化的，直接在for循环中，背包容量可以从0开始
	//dp[..][..][0]
	//var tmpZeroCount int
	for i:= 1;i<len(strs);i++ {
		tmpZeroCount, tmpOneCount = count(strs[i])
		for j:= 0;j<m+1;j++ {
			if j >= tmpZeroCount && tmpOneCount == 0{
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-tmpZeroCount][0] + 1)
			} else {
				dp[i][j][0] = dp[i-1][j][0]
			}
		}
	}

	for i:=1;i<len(strs);i++ {
		zeroCount, oneCount := count(strs[i])
		// 这里直接从0开始，背包容量为0，可以在循环中进行处理
		for j:=1;j<m+1;j++ {
			for k:=1;k<n+1;k++ {
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

func main() {
	//testStrs := []string{"10","0001","111001","1","0"}
	//m, n := 5, 3
	testStrs := []string{"0","0","1","1"}
	m, n := 2, 2
	//i, j, k := 1, 2, 0
	//[[[0 0 0] [1 1 1] [1 1 1]] [[0 0 0] [1 1 1] [1 2 2]] [[0 1 1] [0 2 2] [0 2 3]] [[0 1 1] [0 2 3] [0 2 3]]]
	//[[[0 0 0] [1 1 1] [1 1 1]] [[0 0 0] [1 1 1] [2 2 2]] [[0 1 1] [1 2 2] [2 3 3]] [[0 1 2] [1 2 3] [2 3 4]]]

	res := findMaxForm(testStrs, m, n)
	fmt.Println(res)
}