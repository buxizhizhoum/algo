package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	resCount := make(map[string]int, 0)
	backtrack(candidates, target, 0, resCount, track, &res)
	return res
}


func backtrack(candidates []int, target int, start int, resCount map[string]int, track []int, res *[][]int) {
	fmt.Println(target)
	//if sum(track) > target {
	//	return
	//}
	if target == 0 {
		tmp := make([]int, len(track))
		copy(tmp, track)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
 		})
		tmpStr := convertStr(tmp)
		key := strings.Join(tmpStr, ":")
		if count, _ := resCount[key]; count >=1 {
			return
		} else {
			*res = append(*res, tmp)
			resCount[key] += 1
			return
		}

	} else if target < 0 {
		return
	}

	for i:=start;i<len(candidates);i++{
		track = append(track, candidates[i])
		backtrack(candidates, target - candidates[i], i, resCount, track, res)
		track = track[:len(track) - 1]
	}
	return
}

func convertStr(s []int) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = strconv.Itoa(v)
	}
	return res
}

func sum(s []int) int {
	res := 0
	for _, v := range s {
		res += v
	}
	return res
}

func main() {
//	res := combinationSum([]int{2,3,6,7}, 7)
	res := combinationSum([]int{100,200,4,12}, 400)
	fmt.Println(res)
}
