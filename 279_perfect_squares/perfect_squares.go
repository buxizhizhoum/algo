package main

func numSquares(n int) int {
	res := backtrack(n)
	return res
}

func backtrack(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	res := n + 1 // 初始化
	for i := 1; i * i <= n; i++ {
		if n - i * i < 0 {
			continue
		}
		tmp := backtrack(n - i * i)
		res = min(res, tmp + 1)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	numSquares(12)
}
