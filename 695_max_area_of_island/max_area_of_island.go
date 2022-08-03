package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i:=0;i<m;i++{
		visited[i] = make([]int, n)
	}
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	res := 0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			tmp := dfs(grid, visited, directions, i, j, m, n)
			res = max(res, tmp)
		}
	}
	return res
}

func dfs(grid [][]int, visited [][]int, directions [][]int, i, j, m, n int) int {
	if i<0||j<0||i>=m||j>=n{
		return 0
	}
	if visited[i][j] == 1 {
		return 0
	}
	if grid[i][j] == 0 {
		visited[i][j] = 1
		return 0
	}
	visited[i][j] = 1
	res := 1
	for _, dir := range directions {
		tmp := dfs(grid, visited, directions, i+dir[0], j+dir[1], m, n)
		res += tmp
	}
	return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	testGrid := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	res := maxAreaOfIsland(testGrid)
	fmt.Println(res)
}
