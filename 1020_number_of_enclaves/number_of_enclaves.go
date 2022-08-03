package number_of_enclaves

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i:=0;i < m;i++ {
		visited[i] = make([]int, n)
	}
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	res := 0
	for i:=0 ;i<m;i++{
		dfs(grid, visited, directions, i, 0, m, n, &res)
		dfs(grid, visited, directions, i, n-1, m, n, &res)
	}
	for j:=0;j<n;j++{
		dfs(grid, visited, directions, 0, j, m, n, &res)
		dfs(grid, visited, directions, m-1, j, m, n, &res)
	}
	res = 0
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++{
			if grid[i][j] == 1 && visited[i][j] == 0 {
				dfs(grid, visited, directions, i, j, m, n, &res)
			}
		}
	}
	return res
}

func dfs(grid [][]int, visited [][]int, directions[][]int, i, j, m, n int, res *int) {
	if i<0||j <0||i>=m||j>=n {
		return
	}
	if visited[i][j] == 1 {
		return
	}
	if grid[i][j] == 0 {
		visited[i][j] = 1
		return
	}
	visited[i][j] = 1
	*res += 1
	for _, dir := range directions {
		dfs(grid, visited, directions, i+dir[0], j+dir[1], m, n, res)
	}
	return
}
