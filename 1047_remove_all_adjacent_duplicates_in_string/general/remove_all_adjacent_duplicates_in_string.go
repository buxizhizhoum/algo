package main

import "fmt"

type Element struct {
	Char byte
	Count int
}

// 这个写法是开心消消了，本意就是所有靠在一起的重复的都删除
// 先往栈里面添加元素，当下一个字符和栈顶元素不相等的时候，判断栈是否需要弹出
func removeDuplicates(s string) string {
	stack := make([]*Element, 0)
	for i:=0;i<len(s);i++ {
		top := getTop(stack)
		// 根据是否于top相同，判断添加字符的计数
		if top != nil && top.Char == s[i] {
			stack = append(stack, &Element{Char: s[i], Count: stack[len(stack) - 1].Count + 1})
		} else {
			stack = append(stack, &Element{Char: s[i], Count: 1})
		}
		// 下一个字符是什么，后面根据下一个字符判断是继续添加还是开始判断弹出
		var nextChar byte
		if i < len(s) - 1 {
			nextChar = s[i+1]
		} else {
			// 越界以后，nextChar应该是一个字符串中不存在的数
			nextChar = 0
		}

		top = getTop(stack)
		if top != nil && top.Char != nextChar && top.Count >= 2 {
			// 满足条件的都弹出
			top = stack[len(stack)-1]
			// 目前的字符，与之相等的都需要弹出
			curChar := top.Char
			for ; top != nil && top.Count > 0 && top.Char == curChar; {
				stack = stack[:len(stack)-1]
				top = getTop(stack)
			}
		}
	}

	res := make([]byte, 0)
	// 把stack里面的转换成字符串
	for _, item := range stack {
		tmp := make([]byte, item.Count)
		for i:=0;i<item.Count;i++ {
			tmp[i] = item.Char
		}
		res = append(res, tmp...)
	}
	return string(res)
}


func getTop(stack []*Element) *Element {
	var res *Element
	if len(stack) > 0 {
		res = stack[len(stack)-1]
	} else {
		res = nil
	}
	return res
}

func main() {
	res := removeDuplicates("abbaca")
	//res := removeDuplicates("aaaaaaaa")
	fmt.Println("res: ", res)
}
