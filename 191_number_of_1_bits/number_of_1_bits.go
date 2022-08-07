package number_of_1_bits

//核心逻辑就是，n - 1 一定可以消除最后一个 1，同时把其后的 0 都变成 1，这样再和 n 做一次 & 运算，就可以仅仅把最后一个 1 变成 0 了
func hammingWeight(num uint32) int {
	res := 0
	for ; num != 0; {
		num = num & (num - 1)
		res += 1
	}
	return res
}

