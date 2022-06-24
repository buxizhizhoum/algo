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
	// 唯一字符个数
	valid := 0

	// 长度，区间，存长度为了后续判断方便
	ans := []int{math.MaxInt32, 0, 0}
	// 发现某个字符在 window 的数量满足了 need 的需要，就要更新 valid，表示有一个字符已经满足要求
	for right < len(s){
		window[s[right]]++// 支持吗？
		// 当右边指针指向的字符属于需要的字符，且在上面扩展右边界后，窗口内字符数和需要的字符数相等
		// 有效长度加1
		// valid表示有几个字符已经满足需求
		if need[s[right]] > 0 && window[s[right]] == need[s[right]] {
			valid++
		}
		// valid == len(need) 表示window中包含了可以组成字符串t的字符，即已经找到了一个可行的解
		for left <= right && valid == len(need) {
			// 更新结果
			// right - left + 1 is the length of the string
			if right - left + 1 < ans[0] {
				ans = []int{right - left + 1, left, right + 1}
			}

			window[s[left]]--
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
