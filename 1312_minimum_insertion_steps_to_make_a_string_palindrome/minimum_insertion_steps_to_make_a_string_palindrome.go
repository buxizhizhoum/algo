package main


func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i:= 0; i< n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i< n; i++ {
		dp[i][i] = 0
	}

	for i:=n-2;i>=0;i--{
		// j的取值是怎么考虑的?
		for j:=i+1;j<n;j++ {
			//如果s[i] == s[j] 那么此时要求的就是 s[i+1…j-1] 变成回文串的最少插入次数了，即 dp[i+1][j-1]
			//如果s[i] != s[j] 此时我们可以考虑 在 最右边插入一个 s[i] 那么就变成了 s[i+1…j] 的 最少插入次数 再 加上一次， 或者在最左边插入一个s[j] ，即 s[i…j-1] 的 最少插入次数 + 1
			//此时的dp[i][j] = min(dp[i+1][j],dp[i][j-1]) +1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minInsertions("zzazz")
}
