package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))
	// 必须要排序
	sort.Slice(nums, func (i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrack(nums, used, track, &res)
	return res
}

func backtrack(nums []int, used []bool, track []int, res *[][]int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i:=0;i<len(nums);i++{
		if used[i] {
			continue
		}
		// 同样的数已经用过了
		if i > 0 && nums[i] == nums[i-1] && used[i-1] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, used, track, res)
		track = track[:len(track) - 1]
		used[i] = false
	}
}


func main() {
	permuteUnique([]int{1,1,2})
}
