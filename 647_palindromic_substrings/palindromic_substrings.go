package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i:=0;i<len(s);i++ {
		res += palindromicCount(s, i, i)
		res += palindromicCount(s, i, i+1)
		//fmt.Println(res)
	}
	return res
}

func palindromicCount(s string, start, end int) int {
	res := 0
	fmt.Println(s, start, end)
	for ;start >=0 && end < len(s); {
		if s[start] == s[end] {
			res++
		} else {
			break
		}
		start--
		end++
	}
	//fmt.Println(s, start, end, res)
	return res
}

func main() {
	res := countSubstrings("fdsklf")
	fmt.Println(res)
}
