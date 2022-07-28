package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make(map[int]bool, 0)
	backtrack(nums, used, []int{}, &res)
	return res
}

func backtrack(nums []int, used map[int]bool, track []int, res *[][]int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i:=0;i<len(nums);i++ {
		if flag, ok := used[i]; flag && ok {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, used, track, res)
		track = track[:len(track) - 1]
		used[i] = false
	}
	return
}

func main() {
	res := permute([]int{5,4,2,6})
	fmt.Println(res)
}
