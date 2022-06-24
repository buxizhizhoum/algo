package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i:= 0;i < n;i++ {
		matrix[i] = make([]int, n)
	}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	cur := []int{0, 0}
	directionIndex := 0
	dir := directions[directionIndex]
	num := 1
	var nextPoint []int
	for num <= n * n {
		nextPoint = []int{cur[0] + dir[0], cur[1] + dir[1]}
		if !isValid(nextPoint[0], nextPoint[1], n, n, matrix) {
			directionIndex = (directionIndex + 1) % 4
			dir = directions[directionIndex]
			nextPoint = []int{cur[0] + dir[0], cur[1] + dir[1]}
		}
		matrix[cur[0]][cur[1]] = num
		cur = nextPoint
		num++
	}
	return matrix
}

func isValid(x, y, m, n int, matrix [][]int) bool {
	if x >=0 && x < m && y >= 0 && y < n && matrix[x][y] == 0{
		return true
	}
	return false
}

func main() {
	generateMatrix(3)
}
