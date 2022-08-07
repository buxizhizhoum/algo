package super_pow

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	var last int
	b, last = b[:len(b) - 1], b[len(b) - 1]
	part1 := specialPow(a, last)
	part2 := specialPow(superPow(a, b), 10)
	return (part1 * part2) % 1337

}

// 计算a的k次方，然后和k求模的结果，每次都取模避免溢出
func specialPow(a, k int) int {
	a = a % 1337
	res := 1
	// a的k次方不一次算，一轮一轮算
	for i:=0;i<k;i++{
		res *= a
		res %= 1337
	}
	return res
}

