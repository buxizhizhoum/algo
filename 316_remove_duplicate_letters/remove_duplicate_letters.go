package main

import "fmt"

func removeDuplicateLetters(s string) string {
	// 用于对字符串中出现的字符进行计数
	counter := make(map[byte]int, 0)
	for i:=0;i<len(s);i++{
		counter[s[i]]++
	}
	stack := make([]byte, 0)
	inStack := make(map[byte]bool, 0)

	for i:=0;i<len(s);i++{
		// 计数调整
		counter[s[i]]--
		// 如果这个字符在栈中已经存在了，跳过
		if val, ok := inStack[s[i]]; ok && val {
			continue
		}
		// 栈非空，当前元素小于栈顶元素，尝试弹出栈顶元素，这样可以使字典序最小
		for ;len(stack) > 0 && s[i] < stack[len(stack) - 1]; {
			count := counter[stack[len(stack) - 1]]
			// 后面没有栈顶元素了
			if count == 0 {
				break
			} else {
				// 存在标志位修改
				delete(inStack, stack[len(stack) - 1])
				stack = stack[0:len(stack) - 1]
			}
		}
		stack = append(stack, s[i])
		inStack[s[i]] = true
	}
	//reverse(&stack)
	return string(stack)
}


//func reverse(s *[]byte) {
//	for i, j := 0, len(*s) - 1; i<j;i, j = i+1, j-1 {
//		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
//	}
//}

func main() {
	//res := removeDuplicateLetters("bcabc")
	res := removeDuplicateLetters("cbacdcb")
	fmt.Println(res)
}
