package main

import "fmt"

func partition(s string) [][]string {
	track := make([]string, 0)
	res := make([][]string, 0)
	backtrack(s, track, &res, 0)
	return res
}

func backtrack(s string, track []string, res *[][]string, startIndex int) {
	if totalLen(track) == len(s) {
		tmp := make([]string, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
	}

	for i := startIndex; i < len(s);i++ {
		// 切割要求每个字符串都是回文，一旦遇到不是回文的情况，直接返回
		if !isPalindrome(s, startIndex, i) {
			continue
		}
		// 回文字符串，需要startIndex，到 i之间的字符串
		track = append(track, s[startIndex: i+1])
		fmt.Println(track)
		backtrack(s, track, res, i + 1)
		track = track[:len(track) - 1]
	}
}

func totalLen(track []string) int {
	res := 0
	for i := 0;i<len(track);i++ {
		res += len(track[i])
	}
	return res
}

func isPalindrome(s string, i, j int) bool {
	for start, end := i, j;start < end; start, end = start +1, end - 1  {
		if s[start] != s[end] {
			return false
		}
	}
	return true
}

func main() {
	res := partition("aab")
	fmt.Println(res)
}