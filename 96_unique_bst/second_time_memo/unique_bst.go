package second_time_memo

import "fmt"

func numTrees(n int) int {
	memo := make(map[string]int, 0)
	res := count(1, n, memo)
	return res
}

func count(l, r int, memo map[string]int) int {
	if l >= r {
		return 1
	}
	key := fmt.Sprintf("%s:%s", l, r)
	if val, ok := memo[key]; ok {
		return val
	}
	// 计算[l, r]区间能产生的bst
	res := 0
	// 在l, r 之间，每个值都可以将区间分为两部分，然后分别算这两部分可以组成的bst
	for i:=l; i<=r;i++ {
		left := count(l, i-1, memo)
		right := count(i+1, r, memo)
		res += left * right
	}
	memo[key] = res
	return res
}
