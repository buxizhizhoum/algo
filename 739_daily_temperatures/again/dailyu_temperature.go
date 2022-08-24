package again

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stk := make([]int, 0)
	for i:=len(temperatures) -1;i>=0 ;i-- {
		// 比当前这个温度小的，都弹出，用不到了，后续都有当前的温度挡着，对再前面的都不可见
		// temperatures[stk[len(stk) - 1]表示栈顶元素代表的温度，这里比温度
		for ;len(stk) > 0 && temperatures[stk[len(stk) - 1]] <= temperatures[i]; {
			stk = stk[:len(stk) - 1]
		}
		// 这里在栈里面推进了索引，为了计算等待天数
		// 栈里面的元素已经是比当前这个元素更大的元素的索引了，
		if len(stk) == 0 {
			res[i] = 0
		} else {
			// 栈顶元素是下一个更大元素的索引，相减算间距
			res[i] = stk[len(stk) - 1] - i
		}
		// 当前这个元素代表了这个时候发现的最大值，入栈
		stk = append(stk, i)
	}
	return res
}
