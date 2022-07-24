package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type Node struct {
	From int
	Price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	indgree := make(map[int][]Node, 0)
	for _, item := range flights {
		from, to, price := item[0], item[1], item[2]
		indgree[to] = append(indgree[to], Node{From: from, Price: price})
	}
	// k+1: 转中转次数为节点个数
	k +=1
	memo := make([][]int, n)
	for i:=0;i<n;i++{
		memo[i] = make([]int, k+1)
	}
	for i:= 0;i<n;i++{
		for j:= 0;j<k+1;j++{
			memo[i][j] = INT_MIN
		}
	}

	return dp(indgree, src, dst, k, memo)
}

func dp(indgree map[int][]Node, src, dst, k int, memo [][]int) int{
	if src == dst {
		return 0
	}
	if k == 0 {
		return -1
	}
	res := INT_MAX
	nodes, ok := indgree[dst]
	if !ok {
		return -1
	}
	if memo[dst][k] != INT_MIN {
		return memo[dst][k]
	}
	for _, upperStreamNode := range nodes {
		subProblem := dp(indgree, src, upperStreamNode.From, k-1, memo)
		if subProblem != -1 {
			res = min(res, subProblem + upperStreamNode.Price)
		}
	}

	if res == INT_MAX {
		return -1
	}
	memo[dst][k] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testFlignts := [][]int{
		{1,0,5},
	}
	res := findCheapestPrice(1, testFlignts,0,1, 1)
	fmt.Println(res)
}

//[[1,0,5]] 0 1 1

