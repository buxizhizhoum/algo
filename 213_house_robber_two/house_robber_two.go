package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
//假设我们有5个房a、b、c、d、e，使用动态规划分别求偷a,b,c,d能得到的最大的钱数，
//以及偷b,c,d,e能得到的最大钱数，从这两个钱数里面取较大者，就是我们能从a,b,c,d,e
//里面可以偷到的最大钱数。

// 另一个思路是用house robber 1的思路对rob n-1, rob n 算两次，取最大值
// todo: 强烈建议用 1 的思路算两遍
func rob(nums []int) int {
	n := len(nums)
	// dp1覆盖0到n-1
	dp1 := make([]int, n-1)
	// dp2 覆盖1-n
	dp2 := make([]int, n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	//覆盖0到n-1
	dp1[0] = nums[0]
	dp1[1] = max(nums[0], nums[1])
	//覆盖1到n
	dp2[1] = nums[1]
	dp2[2] = max(nums[1], nums[2])
	for i:= 2;i<n-1;i++{
		dp1[i] = max(dp1[i-1], dp1[i-2] + nums[i])
	}
	// 2已经算了，这里不能从2开始，会覆盖2的计算结果
	for i:=3;i<n;i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2] + nums[i])
	}
	return max(dp1[n-2], dp2[n-1])
}

func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}

func main() {
	//res := rob([]int{1,2,3})
	//res := rob([]int{1,2,1,1})
	res := rob([]int{200,3,140,20,10})
	fmt.Println(res)
}
