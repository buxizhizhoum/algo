package permutation

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
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
	// 通过这个实现同一层的重复元素去重
	prevNum := -666
	for i := 0;i<len(nums);i++ {
		if used[i] {
			continue
		}
		// 遇到重复元素，在同一层后面的都不取了
		if i > 0 && nums[i-1] == nums[i] && nums[i] == prevNum {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		prevNum = nums[i]
		backtrack(nums, track, res, used)
		track = track[:len(track) - 1]
		used[i] = false
	}
}
