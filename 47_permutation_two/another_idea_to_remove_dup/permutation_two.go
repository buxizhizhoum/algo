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

	for i := 0;i<len(nums);i++ {
		if used[i] {
			continue
		}
		// 遇到重复元素，在同一层后面的都不取了，
		// 之所以加used[i-1] = false是为了在树枝去重中不影响，可以继续向下递归，
		// 否则这一层的一个元素取了，下一层遇到一个重复元素，就跳过了，实际上是不用的
		if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, track, res, used)
		track = track[:len(track) - 1]
		used[i] = false
	}
}
