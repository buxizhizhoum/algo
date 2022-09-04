package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	grid := make([][]string, n)
	for i := 0;i<n;i++ {
		grid[i] = make([]string, n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			grid[i][j] = "."
		}
	}
	res := make([][]string, 0)
	backtrack(n, grid, &res, 0)
	return res
}


func backtrack(n int, grid [][]string, res *[][]string, row int) {
	if row >= n {
		tmp := prepareResult(grid)
		*res = append(*res, tmp)
		return
	}

	for i := 0 ;i < n; i++ {
		valid := isValid(grid, row, i, n)
		//printGrid(grid)
		//fmt.Println("valid, i, j", valid, row, i)
		if !valid {
			continue
		}
		grid[row][i] = "Q"
		backtrack(n, grid, res, row + 1)
		grid[row][i] = "."
	}
}

func isValid(grid [][]string, row, col int, n int) bool {
	// 这一行有没有
	for i := 0; i < n; i++ {
		if grid[row][i] == "Q" && i != col {
			return false
		}
	}
	// 这一例有没有
	for j := 0;j< n; j++ {
		if grid[j][col] == "Q" && j != row {
			return false
		}
	}
	// 左上到右下斜对角平行的线，只需要判断上面的部分就够了，下面还没放
	for i, j := row, col;i >= 0 && j >= 0 ; i, j = i-1, j-1 {
		//fmt.Println(i, j)
		if grid[i][j] == "Q" && i != row && j != col {
			return false
		}
	}
	// 另一条对角线平行的线，只需要判断上面的线就行，下面不用判断
	for i, j := row, col; i >= 0  && j < n ; i, j = i - 1, j + 1 {
		if grid[i][j] == "Q" && i != row && j != col {
			return false
		}
	}
	return true
}

func prepareResult(grid [][]string) []string {
	res := make([]string, len(grid))
	for i, row := range grid {
		tmp := strings.Join(row, "")
		res[i] = tmp
	}
	return res
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}

}

func main() {
	res := solveNQueens(4)
	fmt.Println(res)
}
