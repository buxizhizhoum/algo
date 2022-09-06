package main

import (
	"fmt"
	"strings"
)

type Element struct {
	Val string
	Count int
}

// 先入栈，入栈后，根据后面一个元素和现在这个元素是否相同，决定是否对栈中元素进行弹出
// ref 1047
func removeDuplicate(s string) string {
	stack := make([]*Element, 0)
	var nextChar string
	var peek *Element
	for i:= 0;i<len(s);i++ {
		peek = getPeek(stack)
		// 入栈的时候，根据栈顶元素判断，现在这个元素相邻的有几个
		if peek != nil && peek.Val == string(s[i]) {
			stack = append(stack, &Element{
				Val:   string(s[i]),
				Count: peek.Count + 1,
			})
		} else {
			stack = append(stack, &Element{
				Val:   string(s[i]),
				Count: 1,
			})
		}
		// 如果next char 越界，用一个字符串中不存在的字符
		if i < len(s) - 1 {
			nextChar = string(s[i+1])
		} else {
			nextChar = ""
		}

		peek = getPeek(stack)
		if nextChar != string(s[i]) && peek != nil && peek.Count >= 3 {
			// 需要进行弹出
			cur := peek
			// 只弹出目前这个重复相邻的字符
			dupChar := string(s[i])
			for ;cur != nil && cur.Val == dupChar; {
				stack = stack[:len(stack) - 1]
				cur = getPeek(stack)
			}
		}
	}
	res := buildString(stack)
	return res
}

func getPeek(stack []*Element) *Element {
	var peek *Element
	if len(stack) > 0 {
		peek = stack[len(stack) - 1]
	} else {
		peek = nil
	}
	return peek
}

func buildString(stack []*Element) string {
	res := make([]string, len(stack))
	for i, val := range stack {
		res[i] = val.Val
	}
	return strings.Join(res, "")
}


func main() {
	res := removeDuplicate("aabbbaaccddddac")
	fmt.Println(res)
}

