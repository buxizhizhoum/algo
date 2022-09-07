package second_time

func numTrees(n int) int {
	res := count(1, n)
	return res
}

func count(l, r int) int {
	if l >= r {
		return 1
	}
	// 计算[l, r]区间能产生的bst
	res := 0
	// 在l, r 之间，每个值都可以将区间分为两部分，然后分别算这两部分可以组成的bst
	for i:=l; i<=r;i++ {
		left := count(l, i-1)
		right := count(i+1, r)
		res += left * right
	}
	return res
}
