package longest_palindrome_substring

func longestPalindrome(s string) string {
	res := ""
	for i:=0; i < len(s); i++ {
		subStr := findPalindrome(s, i, i)
		subStr2 := findPalindrome(s, i, i + 1)
		if len(subStr) > len(res) {
			res = subStr
		}
		if len(subStr2) > len(res) {
			res = subStr2
		}
	}
	return res
}

// 从中心向两端找回文字符串
func findPalindrome(s string, l, r int) string {
	strLength := len(s)
	for ; r < strLength && l >= 0; {
		if s[l] == s[r] {
			r++
			l--
		} else {
			break
		}
	}
	// 循环中，先操作，再--， 这里要l + 1, 因为r本身为开区间，不用再修改
	return s[l + 1: r]
}
