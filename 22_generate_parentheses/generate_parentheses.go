package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	track := make([]string, 0)
	backtrack(n, n, track, &res)
	return res
}

func backtrack(left, right int, track []string, res *[]string) {
	// 左括号剩余多
	if left > right {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		tmp := strings.Join(track, "")
		*res = append(*res, tmp)
		return
	}

	choices := []string{"(", ")"}
	for _, choice := range choices {
		if choice == "(" {
			track = append(track, choice)
			backtrack(left -1, right, track, res)
			track = track[:len(track) - 1]
		} else if choice == ")" {
			track = append(track, ")")
			backtrack(left, right -1, track, res)
			track = track[:len(track) - 1]
		}
	}
	return
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
