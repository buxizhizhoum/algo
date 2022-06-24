package permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	for i:=0;i<len(s1);i++{
		need[s1[i]]++
	}

	left := 0
	right := 0
	unique := 0
	for right < len(s2) {
		window[s2[right]]++

		if need[s2[right]] > 0 && window[s2[right]] == need[s2[right]] {
			unique++
		}
		// 这里进循环的判断条件就是找到了一个可能解
		// 循环中对这个解进行检查或者调整左边界再检查
		for left <= right && len(need) == unique {
			// 检查是否满足条件
			// 和循环的条件结合起来，有效字符相等，且字符串长度和s1相等
			if right - left + 1== len(s1) {
				return true
			}

			window[s2[left]]--
			if need[s2[left]] > 0 && window[s2[left]] < need[s2[left]] {
				unique--
			}
			left++
		}
		right++
	}
	return false
}