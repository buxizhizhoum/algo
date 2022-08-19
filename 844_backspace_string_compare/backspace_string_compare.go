package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	// 从前面不好判断现在这个字符是否存在最终的字符串，所以从后面开始
	i, j := len(s) - 1, len(t) - 1
	skipS, skipJ := 0, 0
	// 这里只要有一个没到头，就继续，存在一个遍历完了，另一个还没开始的情况，要都遍历完才能判断
	for ;i >=0 || j >=0; {
		// 计算要回退多少个字符，如果还能回退，遇到字符的时候，实现回退
		for ;i>=0; {
			if s[i] == '#' {
				i--
				skipS++
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for ;j>=0; {
			if t[j] == '#' {
				j--
				skipJ++
			} else if skipJ > 0 {
				skipJ--
				j--
			} else {
				break
			}
		}
		// 已经回退完的情况下，看如果不相等，直接返回false
		if i >=0 && j >=0 && s[i] != t[j] {
			return false
		}
		// 如果一个已经遍历完了，另一个还有，不相等
		if (i < 0) != (j < 0) {
			return false
		}
		i--
		j--
	}
	return true
}

func main() {
	//res := backspaceCompare("ab#c", "ad#c")
	//res := backspaceCompare("xywrrmp", "xywrrmu#p")
	res := backspaceCompare("bxj##tw", "bxj###tw")
	fmt.Println(res)
}
