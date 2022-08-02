package main

import (
	"sort"
	"strconv"
	"strings"
)

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	// 为什么track不用地址？
	used := make([]bool, n+1)
	resCache := make(map[string]int, 0)
	backtrack(n, k, used, resCache, track, &res)
	return res
}

func backtrack(n int, k int, used []bool, resCache map[string]int, track []int, res *[][]int) {
	// 取子集树里面第二层的节点就可以
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		// 对结果去重
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		key := convertStr(tmp)
		if count, _ := resCache[key]; count >= 1 {
			return
		} else {
			*res = append(*res, tmp)
			resCache[key] += 1
			return
		}
	}
	for i:= 1;i <= n;i++ {
		// 组合问题最好通过修改i的其实未知去重，通过used去重，结果中还是会有重复，还需要去重
		if used[i] {
			continue
		}
		used[i] = true
		track = append(track, i)
		backtrack(n, k, used, resCache, track, res)
		track = track[:len(track) - 1]
		used[i] = false
	}
}

func convertStr(s []int) string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = strconv.Itoa(v)
	}
	return strings.Join(res, ":")
}

func main() {
	combine(13, 13)
}
