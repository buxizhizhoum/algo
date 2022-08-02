package main

func numIslands(grid [][]byte) int {
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	res := 0
	m := len(grid)
	n := len(grid[0])
	for i:= 0;i<m;i++ {
		for j:=0;j<n;j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs(grid, directions, i, j, m, n)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, directions [][]int, i, j int, m, n int) {
	if i <0 || i >= m || j < 0 || j >= n {
		return
	}
	// 已经是海水了
	if grid[i][j] == '0' {
		return
	}
	// 先把岛淹了
	grid[i][j] = '0'
	for _, dir := range directions {
		dfs(grid, directions, i+dir[0], j+dir[1], m, n)
	}
	return
}

func main() {
	grid := [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
	numIslands(grid)
}

