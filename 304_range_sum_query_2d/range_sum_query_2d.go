package main

import "fmt"

type NumMatrix struct {
	PreSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{
			PreSum: nil,
		}
	}
	n := len(matrix[0])

	preSum := make([][]int, m+1)
	for i:=range preSum {
		preSum[i] = make([]int, n+1)
	}
	// 初始化0行，0列
	for i:=0; i < m; i++ {
		preSum[i][0] = 0
	}
	for j:=0; j < n; j++ {
		preSum[0][j] = 0
	}

	for i:=1; i < m + 1; i++ {
		for j:=1; j < n + 1 ; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{
		PreSum: preSum,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.PreSum == nil {
		return 0
	}
	return this.PreSum[row2 + 1][col2 + 1] - this.PreSum[row2+1][col1] - this.PreSum[row1][col2+1] + this.PreSum[row1][col1]
}


func main() {
	testMatrix := [][]int{
	{3,0,1,4,2},{5,6,3,2,1},{1,2,0,1,5},{4,1,0,1,7},{1,0,3,0,5},
	}
	obj := Constructor(testMatrix)
	res := obj.SumRegion(2,1,4,3)
	fmt.Println(res)
}
