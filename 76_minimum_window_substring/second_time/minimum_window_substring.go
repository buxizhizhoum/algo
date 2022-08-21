package main

import "fmt"

// 这个解法更好懂一点，时间复杂度o(M*N)
// 窗口包含的时候，收缩左侧边界，优化结果
func minWindow(s string, t string) string {
	need := make(map[byte]int, 0)
	for i:=0;i<len(t);i++ {
		// 只需要包含，设置成1即可
		need[t[i]]++
	}
	window := make(map[byte]int, 0)
	left, right := 0, 0
	res := s + ":"
	for ;right < len(s); {
		window[s[right]]++
		if contain(window, need) {
			if right - left + 1 < len(res) {
				res = s[left:right+1]
			}
		}

		for ;contain(window, need); {
			window[s[left]]--
			// 如果已经不在字典中，直接删除键
			if val, ok := window[s[left]]; ok && val == 0{
				delete(window, s[left])
			}
			left++
			if contain(window, need) {
				if right - left + 1 < len(res) {
					res = s[left:right+1]
				}
			}
		}
		right++
	}
	if res == s + ":" {
		return ""
	}
	return res
}

func contain(window, need map[byte]int) bool {
	for k, v := range need {
		val, ok := window[k]
		if !ok || val < v {
			return false
		}
	}
	return true
}

func main() {
	//res := minWindow("ADOBECODEBANC","ABC")
	//res := minWindow("A","A")
	res := minWindow("A","AA")
	fmt.Println(res)
}
