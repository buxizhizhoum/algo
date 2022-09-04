package main

import (
	"fmt"
	"sort"
)

// 这里使用两个标识进行去重
// 一个是同层去重的levelUsed, 另一个是对路径元素去重的used
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	// 这里是路径上使用过的
	used := make([]bool, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return  nums[i] < nums[j]
	})
	backtrack(nums, track, &res, used)
	return res
}


func backtrack(nums []int, track []int, res *[][]int, used []bool) {
	if len(nums) == len(track) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	levelUsed := make(map[int]bool, len(nums))
	for i := 0;i<len(nums);i++ {
		if used[i] {
			continue
		}
		// 遇到重复元素，在同一层后面的都不取了
		if i > 0 && nums[i-1] == nums[i] && levelUsed[nums[i-1]] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		levelUsed[nums[i]] = true
		backtrack(nums, track, res, used)
		track = track[:len(track) - 1]
		used[i] = false
		// 这里不需要再对levelUsed进行回溯，这里是对同层去重，下层会重新初始化
	}
}

func main() {
	nums := []int{1,1,2}
	res := permuteUnique(nums)
	fmt.Println(res)
}
