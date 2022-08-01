package main


func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, 0, track, &res)
	return res
}

// 想想成对子集树的前序遍历，在开始的时候，就把结果添加到最终结果中
func backtrack(nums []int, start int, track []int, res *[][]int) {
	// 这里不能直接放track，track在后续会发生变化
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)
	// 每次从start开始而不是从0开始，确保了不会重复
	for i:= start;i<len(nums);i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, track, res)
		track = track[:len(track) - 1]
	}
}

//解答失败: 测试用例:[1,2,3]
//测试结果:[[],[3],[1,3],[1,2,3],[1,3],[3],[2,3],[3]]
//期望结果:[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]] stdout:
