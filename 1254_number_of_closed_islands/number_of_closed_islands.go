package main

import "fmt"

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i:=0;i<m;i++ {
		visited[i] = make([]int, n)
	}

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for i:=0;i<m;i++ {
		// 左右量列相连的处理了
		dfs(grid, visited, directions, i, 0, m, n)
		dfs(grid, visited, directions, i, n-1, m, n)
	}
	for j:=0;j<n;j++{
		//上下两行的处理了
		dfs(grid, visited, directions, 0, j, m, n)
		dfs(grid, visited, directions, m-1, j, m, n)
	}
	res := 0
	for i:=1;i<m;i++{
		for j:=1;j<n;j++ {
			// 每次遍历一个隔离起来的岛屿
			if grid[i][j] == 0 && visited[i][j] == 0{
				res += 1
				dfs(grid, visited, directions, i, j, m, n)
			}
		}
	}
	return res
}


func dfs(grid[][]int, visited[][]int, directions[][]int, i, j, m, n int) {
	if i <0 ||i >=m || j<0 ||j >=n {
		return
	}
	// 已经到陆地了，返回
	if grid[i][j] == 1 {
		visited[i][j] = 1
		return
	}
	if visited[i][j] == 1 {
		return
	}
	visited[i][j] = 1
	for _, dir := range directions {
		dfs(grid, visited, directions, i+dir[0], j+dir[1], m, n)
	}
	return
}


func printGrid(grid [][]int) {
	for _, tmp := range grid {
		fmt.Println(tmp)
	}
	fmt.Println()
}



func main() {
//测试用例:[
	//[0,1,1,1,1,1,1,1],
	//[1,0,1,0,0,0,0,1],
	//[1,0,1,0,0,1,0,1],
	//[1,0,0,0,0,1,0,1],
	//[1,0,0,1,0,1,0,1],
	//[1,1,0,1,0,0,0,1],
	//[1,0,0,0,0,0,0,1],
	//[0,1,1,1,1,1,1,1]]
	//测试结果:2
	//期望结果:1

	//[0,1,1,1,1,1,1,1],
	//[1,1,1,1,1,1,1,1],
	//[1,1,1,1,1,1,1,1],
	//[1,1,1,1,1,1,1,1],
	//[1,0,0,1,1,1,1,1],
	//[1,1,0,1,1,1,1,1],
	//[1,0,0,0,0,1,1,1],
	//[0,1,1,1,1,1,1,1]]
	testGrid := [][]int{
		{0,1,1,1,1,1,1,1},
		{1,0,1,0,0,0,0,1},
		{1,0,1,0,0,1,0,1},
		{1,0,0,0,0,1,0,1},
		{1,0,0,1,0,1,0,1},
		{1,1,0,1,0,0,0,1},
		{1,0,0,0,0,0,0,1},
		{0,1,1,1,1,1,1,1},
	}
	res := closedIsland(testGrid)
	fmt.Println(res)
}
