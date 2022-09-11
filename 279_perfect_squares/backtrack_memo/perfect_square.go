package backtrack_memo

func numSquares(n int) int {
	memo := make(map[int]int, 0)
	res := backtrack(n, memo)
	return res
}

func backtrack(n int, memo map[int]int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	res := n + 1 // 初始化
	for i := 1; i * i <= n; i++ {
		if n - i * i < 0 {
			continue
		}
		tmp := backtrack(n - i * i, memo)
		res = min(res, tmp + 1)
	}
	memo[n] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
