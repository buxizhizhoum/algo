package main

import "fmt"

func solveSudoku(board [][]byte)  {
	m := len(board)
	n := len(board[0])
	backtrack(board, 0, 0, m, n)
}

func backtrack(board [][]byte, rowStart, colStart int, m, n int) bool {
	if rowStart == m - 1 && colStart == n - 1 {
		return true
	}
	// 处理换行，也可以不传行列，直接判断board，不满足continue, 会慢一点
	if colStart >=n {
		rowStart += 1
		colStart = 0
	}
	for i:=rowStart;i<m;i++ {
		for j:= colStart;j<n;j++ {
			if board[i][j] != '.' {
				continue
			}
			//fmt.Println("start i, j", i, j)
			for number:=byte('1');number<=byte('9');number++ {
				flag := isValid(board, i, j, m, n, number)
				if !flag {
					continue
				}
				board[i][j] = number
				//fmt.Println("i, j ", i, j)
				// todo: 这里传i，j为什么不行？
				res := backtrack(board, 0, 0, m, n)
				if res {
					return true
				}
				board[i][j] = '.'
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, m, n int, number byte) bool {
	// 判断列有没有相同的
	for i := 0; i<m;i++ {
		if board[i][col] == number {
			return false
		}
	}
	// 判断行有没有相同的
	for j:=0;j<n;j++ {
		if board[row][j] == number {
			return false
		}
	}
	// 判断九宫格
	rowStart := row / 3 * 3
	colStart := col / 3 * 3
	for i:= rowStart;i<rowStart + 3;i++ {
		for j:=colStart;j<colStart+3;j++ {
			if board[i][j] == number {
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
//测试用例:[
	//["5","3",".",".","7",".",".",".","."],
	//["6",".",".","1","9","5",".",".","."],
	//[".","9","8",".",".",".",".","6","."],
	//["8",".",".",".","6",".",".",".","3"],
	//["4",".",".","8",".","3",".",".","1"],
	//["7",".",".",".","2",".",".",".","6"],
	//[".","6",".",".",".",".","2","8","."],
	//[".",".",".","4","1","9",".",".","5"],
	//[".",".",".",".","8",".",".","7","9"]]
//测试结果:[
	//["1","2","3","6","4","7","5","9","8"],
	//["6","4","7","5","9","8","1","3","2"],
	//["5","9","8","1","3","2","4","6","7"],
	//["2","1","6","9","7","4","8","5","3"],
	//["4","8","9","3","5","6","7","2","1"],
	//["7","3","5","8","2","1","9","4","6"],
	//["8","5","2","7","6","9","3","1","4"],
	//["9","7","4","2","1","3","6","8","5"],
	//["3","6","1","4","8","5","2","7","9"]]
//期望结果:[
	//["5","3","4","6","7","8","9","1","2"],
	//["6","7","2","1","9","5","3","4","8"],
	//["1","9","8","3","4","2","5","6","7"],
	//["8","5","9","7","6","1","4","2","3"],
	//["4","2","6","8","5","3","7","9","1"],
	//["7","1","3","9","2","4","8","5","6"],
	//["9","6","1","5","3","7","2","8","4"],
	//["2","8","7","4","1","9","6","3","5"],
	//["3","4","5","2","8","6","1","7","9"]]
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
