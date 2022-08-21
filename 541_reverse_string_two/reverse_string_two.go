package main

import "fmt"

func reverseStr(s string, k int) string {
	chars := []byte(s)
	for i:=0;i < len(s);i += 2*k {
		fmt.Println(i)
		if i + k < len(s) {
			reverse(chars, i, i + k - 1)
		} else {
			reverse(chars, i, len(s) - 1)
		}
	}
	return string(chars)
}

func reverse(chars []byte, start, end int) {
	left, right := start, end
	for ;left < right; {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}

func main() {
	res := reverseStr("abcdefg", 2)
	fmt.Println(res)
}
