package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	backtrack(candidates, target, 0, track, &res)
	return res
}


func backtrack(candidates []int, target int, start int, track []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	} else if target < 0 {
		return
	}

	for i:=start;i<len(candidates);i++{
		track = append(track, candidates[i])
		// 注意这里start传入i很关键，当处理到第i个元素的时候，
		// i不需要再考虑之前的情形，因为同样的组合已经在i之前的元素中包含了
		// 这里就相当于不在这里进行剪枝
		backtrack(candidates, target - candidates[i], i, track, res)
		track = track[:len(track) - 1]
	}
	return
}
