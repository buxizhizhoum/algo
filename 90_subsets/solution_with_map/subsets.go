package subsets

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	// 这里和491寻找所有上升自序列不同之处在于需要排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrack(nums, track, &res, 0)
	return res
}

func backtrack(nums []int, track []int, res *[][]int, startIndex int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)

	// 使用used数组进行去重
	used := make(map[int]int, 0)
	for i := startIndex ; i< len(nums); i++ {
		if count, _ := used[nums[i]]; count > 0 {
			continue
		}

		used[nums[i]]++
		track = append(track, nums[i])
		backtrack(nums, track, res, i + 1)
		track = track[:len(track) - 1]
	}
}
