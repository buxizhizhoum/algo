package main

import "fmt"

// 我想：需要的是什么？窗口里面的再做一个map，两个呈现包含关系的时候，收缩左侧边界，时间复杂度会乘以p的长度
// 有一个diff，开始的时候是需要的map的相反map，随着窗口变动，当这个map是空的时候，找到一个答案，diff包含其它字符的时候怎么判断？
// 题解：随着窗口滑动的时候，维护一个valid计数，表示窗口中的字符有几个是满足条件的（字符本身和出现的次数都满足），当满足的时候，找到一个答案
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	res := make([]int, 0)
	for i:=0;i<len(p);i++{
		need[p[i]]++
	}
	left, right := 0, 0
	for ;right < len(s); {
		ch := s[right]
		window[ch]++
		if mapEqual(window, need) {
			res = append(res, left)
		}
		// 找到的字符次数多于需要的，收缩左侧边界
		for ;mapContains(window, need); {
			leftCh := s[left]
			// 这里因为窗口是滑动过来的，window里面一定有left
			window[leftCh]--
			left++
			if val, flag := window[leftCh]; flag && val == 0 {
				delete(window, leftCh)
			}

			if mapEqual(window, need) {
				res = append(res, left)
			}
		}
		right++
	}
	return res
}

func mapEqual(m map[byte]int, n map[byte]int) bool {
	if len(m) != len(n) {
		return false
	}
	for k, v := range m {
		val, ok := n[k]
		if !ok {
			return false
		}
		if val != v {
			return false
		}
	}
	return len(m) == len(n)
}

// whether m contain n
func mapContains(m, n map[byte]int) bool {
	for k, v := range n {
		val, ok := m[k]
		if !ok {
			return false
		}
		if val < v {
			return false
		}
	}
	return true
}

func main() {
	res := findAnagrams("cbaebabacd", "abc")
	fmt.Println(res)
}
