package main

// 先对比两个grid，遇到不同的元素，从这个元素触发flood fill, 剩下的岛屿就是子岛屿
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])
	visited := make([][]int, m)
	for i:= 0;i<m;i++{
		visited[i] = make([]int, n)
	}

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	// 把不同的岛屿淹掉
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++{
			if grid1[i][j] != grid2[i][j] && grid2[i][j] == 1 {
				dfs(grid2, visited, directions, i, j, m, n)
			}
		}
	}
	// 剩下的岛屿都是子岛屿，计算个数
	res := 0
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if grid2[i][j] == 1 && visited[i][j] == 0 {
				res += 1
				dfs(grid2, visited, directions, i, j, m, n)
			}
		}
	}
	return res
}

// 给定一个点，flood fill与之相连的所有岛屿
func dfs(grid [][]int, visited [][]int, directions [][]int, i, j, m, n int) {
	if i < 0||j <0 ||i >=m||j>=n {
		return
	}
	if visited[i][j] == 1 {
		return
	}
	// 遇到了海水
	if grid[i][j] == 0 {
		visited[i][j] = 1
		return
	}
	visited[i][j] = 1
	for _, dir := range directions {
		dfs(grid, visited, directions, i + dir[0], j+ dir[1], m, n)
	}
	return
}

func main() {
	testGrid1 := [][]int{
		{1},
	}
	testGrid2 := [][]int{
		{1},
	}
	countSubIslands(testGrid1, testGrid2)
}