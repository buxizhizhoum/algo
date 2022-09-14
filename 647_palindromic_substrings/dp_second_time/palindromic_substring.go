package dp_second_time

func countSubstrings(s string) int {
	//dp[i][j]：表示区间范围[i,j] （注意是左闭右闭）的子串是否是回文子串
	dp := make([][]bool, len(s))
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]bool, len(s))
	}
	// 初始化对角线
	for i:=0;i<len(s);i++ {
		dp[i][i] = true
	}

	for i:=len(s)-2;i>=0;i-- {
		// 因为进行了初始化，这里j从i + 1开始
		for j:=i+1;j<len(s);j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					// 从右下角推出，决定了一定是i从大到小，j从小到大
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}

	res := 0
	for i:=0;i<len(s);i++{
		for j:=i;j<len(s);j++ {
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}
