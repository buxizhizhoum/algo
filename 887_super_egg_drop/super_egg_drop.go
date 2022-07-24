package main

const INT_MAX = int(^uint(0) >> 1)

func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i:=0;i<k+1;i++{
		memo[i] = make([]int, n + 1)
	}
	for i:=0;i<k+1;i++{
		for j:=0;j<n+1;j++ {
			memo[i][j] = -1
		}
	}
	return dp(k, n, memo)
}

func dp(k, n int, memo [][]int) int {
	if k==1 {
		return n
	}
	if n == 0 {
		return 0
	}
	tmp := memo[k][n]
	if tmp != -1 {
		return tmp
	}
	res := INT_MAX
	for i:=1;i<n+1;i++{
		res = min(
			res,
			max(
				// 碎了
				dp(k-1, i-1, memo),
				// 没碎
				dp(k, n-i, memo),
			) + 1,
		)
	}
	memo[k][n] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	superEggDrop(6, 5000)
}
