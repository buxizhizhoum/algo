package single_number

//一个数和它本身做异或运算结果为 0，即 a ^ a = 0；一个数和 0 做异或运算的结果为它本身，即 a ^ 0 = a
//只要把所有数字进行异或，成对儿的数字就会变成 0，落单的数字和 0 做异或还是它本身，所以最后异或的结果就是只出现一次的元素
func singleNumber(nums []int) int {
	res := nums[0]
	for i:= 1;i<len(nums);i++ {
		res = res ^ nums[i]
	}
	return res
}
