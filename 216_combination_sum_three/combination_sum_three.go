package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	track := make([]int, 0)
	res := make([][]int, 0)
	backtrack(track, &res, k, n, 1)
	return res
}

func backtrack(track []int, res *[][]int, k int, n int, startIndex int) {
	if len(track) == k {
		if sum(track) == n {
			tmp := copy(track)
			*res = append(*res, tmp)
		}
		return
	}
	for i := startIndex;i<=9;i++ {
		// 这里append，里面如果正好满足，添加到res
		track = append(track, i)
		backtrack(track, res, k, n, i + 1)
		// 这里再去掉末尾，然后继续循环，一直是一个slice，后面会覆盖前面
		track = track[:len(track) - 1]
	}
	return
}

func sum(nums []int) int {
	res := 0
	for i:= 0;i<len(nums);i++{
		res += nums[i]
	}
	return res
}

func copy(track []int) []int {
	res := make([]int, len(track))
	for i, val := range track {
		res[i] = val
	}
	return res
}

func main() {
	res := combinationSum3(3, 7)
	fmt.Println(res)
}
