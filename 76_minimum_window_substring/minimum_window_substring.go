package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	window := make(map[byte]int, 0)
	need := make(map[byte]int, 0)
	for i:=0;i< len(t);i++ {
		need[t[i]] += 1
	}

	left := 0
	right := 0
	valid := 0

	// 长度，区间，存长度为了后续判断方便
	ans := []int{math.MaxInt32, 0, 0}

	for right < len(s){
		window[s[right]] += 1 // 支持吗？

		if need[s[right]] > 0 && window[s[right]] == need[s[right]] {
			valid++
		}

		for left <= right && valid == len(need) {
			// 更新结果
			// right - left + 1 is the length of the string
			if right - left + 1 < ans[0] {
				ans = []int{right - left + 1, left, right + 1}
			}

			window[s[left]] -= 1
			if need[s[left]] > 0 && window[s[left]] < need[s[left]] {
				valid--
			}
			left++
		}
		right++
	}

	if ans[0] == math.MaxInt32 {
		return ""
	} else {
		return s[ans[1]: ans[2]]
	}
}


func main() {
	res := minWindow("a", "a")
	fmt.Println("result: ", res)
}
