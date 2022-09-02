package combination_sum_two

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	track := make([]int, 0)
	res := make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	backtrack(candidates, track, target, 0, &res)
	return res
}

func backtrack(candidates []int, track []int, target int, startIndex int, res *[][]int) {
	if target == 0 {
		//tmp := copy(track)
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
	}
	if target < 0 {
		return
	}

	for i:= startIndex; i < len(candidates); i++ {
		// 树层去重，重复的元素不再选取
		// 这里去重i 是 > startIndex, 如果i 只限制 >0, 会过度去重，上一层的也会被考虑进来
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		track = append(track, candidates[i])
		// 元素不能重复选，这里startIndex = i + 1
		backtrack(candidates, track, target - candidates[i], i + 1, res)
		track = track[:len(track) - 1]
	}
}

//func copy(track []int) []int {
//	res := make([]int, len(track))
//	for i, v := range track {
//		res[i] = v
//	}
//	return res
//}
