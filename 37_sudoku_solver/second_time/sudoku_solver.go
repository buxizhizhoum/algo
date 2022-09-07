package main

import "fmt"

func solveSudoku(board [][]byte)  {
	m := len(board)
	n := len(board[0])
	backtrack(board, 0, 0, m, n)
}


func backtrack(board [][]byte, rowStart, colStart int, m, n int) bool {
	if rowStart == m- 1 && colStart == n- 1 {
		return true
	}
	for i:= rowStart;i< m;i++ {
		for j:= colStart;j< n;j++ {
			if board[i][j] != '.' {
				continue
			}
			printBoard(board)
			fmt.Println("start i, j", i, j)
			for number:=byte('1');number<=byte('9');number++ {
				if !isValid(board, i, j, m, n, number) {
					continue
				}
				//printBoard(board)
				board[i][j] = number
				//flag := backtrack(board, i, j, m, n)
				flag := backtrack(board, 0, 0, m, n)
				if flag {
					return true
				}
				board[i][j] = '.'
			}
			// 所有的数都试过了，此路不通，回去，但是这里的两重for循环不好回，所以上面backtrack每次都是从0开始
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, m, n int, number byte) bool {
	for i :=0; i <m; i++{
		if board[i][col] == number && i != row {
		//if board[i][col] == number {
			return false
		}
	}
	for j :=0; j <n; j++{
		if board[row][j] == number && col != j {
		//if board[row][j] == number {
			return false
		}
	}
	// 判断九宫格
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow + 3; i++ {
		for j := startCol; j < startCol + 3; j++ {
			if board[i][j] == number && i != row && j != col {
			//if board[i][j] == number {
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]byte) {
	for _, row := range board {
		for _, col := range row {
			fmt.Print(string(col))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	testBoard := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	solveSudoku(testBoard)
	printBoard(testBoard)
}
