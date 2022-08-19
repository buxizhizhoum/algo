package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	i, j := len(s) - 1, len(t) - 1
	skipS, skipJ := 0, 0
	for ;i >=0 || j >=0; {
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
		if i >=0 && j >=0 && s[i] != t[j] {
			return false
		}
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
