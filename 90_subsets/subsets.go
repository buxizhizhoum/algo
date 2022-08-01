package subsets

import "sort"
// 元素有重复，同样个数的子集里面，重复元素不应当当做两个元素
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	// 通过排序，并且在后面for循环中跳过重复元素避免重复
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrack(nums, 0, track, &res)
	return res
}

func backtrack(nums []int, start int, track []int, res *[][]int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)

	for i:= start;i<len(nums);i++{
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, i+1, track, res)
		track = track[:len(track) - 1]
	}
}
