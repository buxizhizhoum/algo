package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	// 初始化标记矩阵
	visited := make([][]bool, m)
	for i:=0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	directionIndex := 0
	cur := []int{0, 0}
	dir := directions[directionIndex]
	var nextPoint []int
	res := make([]int, 0)
	for len(res) < m * n {
		nextPoint = []int{cur[0] + dir[0], cur[1] + dir[1]}
		if !isValid(nextPoint[0], nextPoint[1], m, n, visited) {
			directionIndex = (directionIndex + 1) % 4
			dir = directions[directionIndex]
			nextPoint = []int{cur[0] + dir[0], cur[1] + dir[1]}
		}
		res = append(res, matrix[cur[0]][cur[1]])
		visited[cur[0]][cur[1]] = true
		cur = nextPoint
	}
	return res
}

func isValid(x, y, m, n int, visited [][]bool) bool {
	if x >=0 && x < m && y >= 0 && y < n && !visited[x][y]{
		return true
	}
	return false
}
