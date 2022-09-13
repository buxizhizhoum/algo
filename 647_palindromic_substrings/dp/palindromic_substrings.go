package dp

//https://programmercarl.com/0647.%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2.html#%E6%9A%B4%E5%8A%9B%E8%A7%A3%E6%B3%95
func countSubstrings(s string) int {
	//dp[i][j]：表示区间范围[i,j] （注意是左闭右闭）的子串是否是回文子串
	dp := make([][]bool, len(s))
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]bool, len(s))
	}

	for i:=len(s)-1;i>=0;i-- {
		for j:=i;j<len(s);j++ {
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
