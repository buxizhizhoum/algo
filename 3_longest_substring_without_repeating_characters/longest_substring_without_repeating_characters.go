package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	window := make(map[byte]int, 0)

	left :=0
	right :=0

	for right < len(s) {
		window[s[right]]++
		// 收缩左边界，直到窗口中因为刚才扩展右边界导致的重复字符的个数减少到1
		for left <= right && window[s[right]] > 1 {
			window[s[left]]--
			left++
		}
		res = max(res, right - left + 1)
		right++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//res := lengthOfLongestSubstring("abcabcbb")
	res := lengthOfLongestSubstring("dvdf")
	fmt.Println(res)
}
