package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	for i:=0;i<len(p);i++{
		need[p[i]]++
	}
	left := 0
	right := 0
	unique := 0
	for right < len(s) {
		window[s[right]]++
		// 判断这个字符是否影响有效长度
		if need[s[right]] >0 && window[s[right]] == need[s[right]] {
			unique++
		}

		for left <= right && len(need) == unique {
			// 判断窗口中的字符串是否满足条件
			if right - left + 1 == len(p) {
				res = append(res, left)
			}
			// 缩减左边界
			window[s[left]]--
			if need[s[left]] > 0 && window[s[left]] < need[s[left]] {
				unique--
			}
			left++
		}
		right++
	}
	return res
}