package main

import (
	"fmt"
	"strings"
)

func numDistinctIslands(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i:= 0;i<m;i++ {
		visited[i] = make([]int, n)
	}
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	res := make(map[string]int, 0)
	sb := make([]string, 0)
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				tmp := dfs(grid, visited, directions, i, j, m, n, "", sb)
				res[tmp] += 1
			}
		}
	}
	return len(res)
}

// 记录岛屿便利的顺序，当做对岛屿的标识，然后通过字符串比对判断岛屿是否相同
func dfs(grid [][]int, visited [][]int, directions [][]int, i, j, m, n int, direction string, sb []string) string {
	if i<0||j<0||i>=m||j>=n{
		return ""
	}
	if visited[i][j] == 1 {
		return ""
	}
	if grid[i][j] == 0 {
		visited[i][j] = 1
		return ""
	}
	sb = append(sb, direction)

	visited[i][j] = 1
	for _, dir := range directions {
		newDirection := fmt.Sprintf("%d:%d", dir[0], dir[1])
		tmp := dfs(grid, visited, directions, i + dir[0], j + dir[1], m, n, newDirection, sb)
		if tmp != "" {
			sb = append(sb, tmp)
		}
	}
	rollbackDirection := fmt.Sprintf("%s%s", "-", direction)
	sb = append(sb, rollbackDirection)
	return strings.Join(sb, ",")
}

func main() {
	testGrid := [][]int{
		{1,1,0,1,1},
		{1,0,0,0,0},
		{0,0,0,0,1},
		{1,1,0,1,1},
	}
	res := numDistinctIslands(testGrid)
	fmt.Println(res)
}
