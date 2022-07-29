package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	res := make([][]string, 0)
	for i := 0;i<n;i++{
		board[i] = make([]string, n)
	}
	// 初始化棋盘
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			board[i][j] = "."
		}
	}
	backtrack(board, n, 0, &res)
	return res
}

func backtrack(board [][]string, n, row int, res *[][]string) {
	if row == n {
		ans := make([]string, n)
		for i, item := range board {
			str := strings.Join(item, "")
			ans[i] = str
		}
		*res = append(*res, ans)
		return
	}

	// 方阵，所以列的数量可以直接用len(board)
	colNums := len(board)
	for j:=0;j<colNums;j++{
		if !isValid(board, n, row, j) {
			continue
		}

		board[row][j] = "Q"
		backtrack(board, n, row+1, res)
		board[row][j] = "."
	}
}

// 在这个位置放，是否合理，因为是从上往下放的，放的时候只需要判断列，右上，左上就可以了
func isValid(board [][]string, n, row, col int) bool {
	// 同一列上面的部分不能有
	for i:=0;i<row;i++ {
		if board[i][col] == "Q"{
			return false
		}
	}

	// 右上角不能有
	for i,j:=row-1,col+1; i>=0 && j<n ; i,j=i-1,j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// 左上角不能有
	for i,j:=row-1,col-1;i>=0 && j>=0;i,j=i-1,j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func main() {
	res := solveNQueens(4)
	fmt.Println(res)
}
