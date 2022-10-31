package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res) - 1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[len(res) -1][1] = max(res[len(res) - 1][1], interval[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
