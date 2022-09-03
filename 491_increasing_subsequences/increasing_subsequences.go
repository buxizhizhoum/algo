package increasing_subsequences

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, track, &res, 0)
	return res
}

func backtrack(nums []int, track []int, res *[][]int, startIndex int) {
	if len(track) > 1 {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
	}
	// 层间去重，所以每层都需要重新初始化
	used := make(map[int]int, 0)
	for i := startIndex; i < len(nums); i++ {
		// 如果使用过，或者不能构成递增子序列
		if used[nums[i]] > 0 || (len(track) > 0 && nums[i] < track[len(track) - 1]) {
			continue
		}
		// 只用来行间去重，不用再pop, 看看90是否一样？
		used[nums[i]]++
		track = append(track, nums[i])
		backtrack(nums, track, res, i + 1)
		track = track[:len(track) - 1]
	}

}
