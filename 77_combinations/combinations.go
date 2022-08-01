package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	// 为什么track不用地址？
	backtrack(n, k, 1, track, &res)
	return res
}

func backtrack(n int, k int, start int, track []int, res *[][]int) {
	// 取子集树里面第二层的节点就可以
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i:= start;i <= n;i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track, res)
		track = track[:len(track) - 1]
	}
}
