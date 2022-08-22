package remove_all_adjacent_duplicates_in_string

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i:=0;i<len(s);i++ {
		if len(stack) != 0 && stack[len(stack) - 1] == s[i]{
			// 删除最后一个元素
			stack = stack[0:len(stack) - 1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
