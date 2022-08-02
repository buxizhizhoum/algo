package main

import "fmt"

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
	visited := make([][]int, m)
	for i:=0;i<m;i++ {
		visited[i] = make([]int, n)
	}
	for i:= 0;i<m;i++ {
		for j:=0;j<n;j++ {
			if grid[i][j] == '1' && visited[i][j] != 1 {
				res += 1
				dfs(grid, directions, visited, i, j, m, n)
			}
		}
	}
	return res
}

// 每次dfs都把和陆地相连的标记为已访问
func dfs(grid [][]byte, directions [][]int, visited [][]int, i, j int, m, n int) {
	if i <0 || i >= m || j < 0 || j >= n {
		return
	}
	// 已经是海水了或者已经标记过了
	if grid[i][j] == '0' || visited[i][j] == 1{
		return
	}
	// 把和陆地相连的都标记为已访问
	visited[i][j] = 1
	for _, dir := range directions {
		dfs(grid, directions, visited, i+dir[0], j+dir[1], m, n)
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
	res := numIslands(grid)
	fmt.Println("res: ", res)
}
