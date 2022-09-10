package backtrack

// todo: 记忆化搜索，超时了
func findMaxForm(strs []string, m int, n int) int {
	res := backtrack(strs, m, n, 0)
	return res
}

func backtrack(strs []string, m, n int, startIndex int) int {
	//if m < 0 || n < 0 {
	if m <= 0 && n <= 0 {
		return 0
	}

	res := 0
	for i:= startIndex; i<len(strs);i++ {
		zeroCount, oneCount := count(strs[i])
		if m - zeroCount >= 0 && n - oneCount >= 0 {
			tmp := backtrack(strs, m - zeroCount, n - oneCount, i + 1)
			res = max(res, tmp + 1)
		}
	}
	return res
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

