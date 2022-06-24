package main


import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i:=0;i<n;i++ {
		matrix[i] = make([]int, n)
	}

	leftBound := 0
	rightBound := n - 1
	upBound := 0
	downBound := n - 1

	num := 1

	for num <= n * n{
		if upBound <= downBound {
			for i := leftBound;i <= rightBound;i++ {
				matrix[upBound][i] = num
				num++
			}
			upBound++
		}

		if leftBound <= rightBound {
			for i:=upBound;i<=downBound;i++{
				matrix[i][rightBound] = num
				num++
			}
			rightBound--
		}

		if upBound <= downBound {
			for i:= rightBound; i >=leftBound; i-- {
				matrix[downBound][i] = num
				num++
			}
			downBound--
		}

		if leftBound <= rightBound {
			for i:= downBound;i>=upBound;i--{
				matrix[i][leftBound] = num
				num++
			}
			leftBound++
		}
	}
	return matrix
}

func main() {
	res := generateMatrix(3)
	fmt.Println(res)
}
