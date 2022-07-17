package main

import "fmt"

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	memo := make([][]int, len(matrix))
	for i, _ := range memo {
		memo[i] = make([]int, len(matrix[0]))
	}
	for i:=0;i < len(memo);i++ {
		for j:=0;j < len(memo[0]);j++{
			// 设置一个特殊值
			memo[i][j] = 66666
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		memo[0][j] = matrix[0][j]
	}

	for i:= 0; i< len(matrix);i++{
		for j:=0;j< len(matrix[0]);j++ {
			dp(matrix, memo, i, j)
		}
	}

	res := 99999
	for j := 0;j < len(memo[0]);j++{
		res = min(res, memo[len(memo)-1][j])
	}
	return res
}


func dp(matrix [][]int, memo [][]int, i, j int) int {
	if i < 0 || j < 0 || i > len(matrix) - 1 || j > len(matrix[0]) - 1 {
		return 99999
	}

	if memo[i][j] != 66666 {
		return memo[i][j]
	}

	res := min(
		min(dp(matrix, memo, i-1, j), dp(matrix, memo, i-1, j-1)),
		dp(matrix, memo, i-1, j+1)) + matrix[i][j]
	memo[i][j] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {
	testMatrix := [][]int{
		{2,1,3},
		{6,5,4},
		{7,8,9},
	}
	res := minFallingPathSum(testMatrix)
	fmt.Println(res)
}

