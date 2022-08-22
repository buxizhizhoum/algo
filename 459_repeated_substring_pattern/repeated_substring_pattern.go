package main

import "strings"

// https://programmercarl.com/0459.%E9%87%8D%E5%A4%8D%E7%9A%84%E5%AD%90%E5%AD%97%E7%AC%A6%E4%B8%B2.html#%E7%A7%BB%E5%8A%A8%E5%8C%B9%E9%85%8D
// 那么既然前面有相同的子串，后面有相同的子串，用 s + s，
//这样组成的字符串中，后面的子串做前串，前后的子串做后串，就一定还能组成一个s
//所以判断字符串s是否有重复子串组成，只要两个s拼接在一起，里面还出现一个s的话，就说明是又重复子串组成。
//当然，我们在判断 s + s 拼接的字符串里是否出现一个s的的时候，
//要刨除 s + s 的首字符和尾字符，这样避免在s+s中搜索出原来的s，我们要搜索的是中间拼接出来的s
func repeatedSubstringPattern(s string) bool {
	tmp := s + s
	tmp = tmp[1:len(tmp)-1]
	return strings.Contains(tmp, s)
}
