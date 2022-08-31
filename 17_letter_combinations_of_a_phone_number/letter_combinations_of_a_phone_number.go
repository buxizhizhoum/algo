package main

import "fmt"

func letterCombinations(digits string) []string {
	num2char := map[byte]string{
		//0: "",
		//1: "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	track := make([]byte, 0)
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	backtrack(num2char, track, &res, 0, digits)
	return res
}

func backtrack(num2char map[byte]string, track []byte, res *[]string, index int, digits string) {
	if len(track) == len(digits) {
		tmp := byte2string(track)
		*res = append(*res, tmp)
		return
	}
	digit := digits[index]
	letters, ok := num2char[digit]
	if !ok {
		return
	}
	for i:=0;i<len(letters);i++ {
		track = append(track, letters[i])
		// 根据index去取下一个数字对应的字符串
		backtrack(num2char, track, res, index + 1, digits)
		track = track[:len(track) - 1]
	}
	return
}

func byte2string(track []byte) string {
	res := make([]byte, len(track))
	for i:=0;i<len(track);i++ {
		res[i] = track[i]
	}
	return string(res)
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
