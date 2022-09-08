package backtrack

import "sort"

// https://leetcode.cn/problems/target-sum/solution/mu-biao-he-by-leetcode-solution-o0cp/
// x 是添加负号组成的数的和
// (sum - x) - x = target => x = (sum - target) / 2

// 这样转化成找到和为(sum - target) / 2 的问题
func findTargetSumWays(nums []int, target int) int {
	total := sum(nums)
	diff := total - target
	if diff % 2 != 0 {
		return 0
	}
	if diff < 0 {
		return 0
	}
	res := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrack(nums, diff / 2, 0, 0, &res)
	return res
}

func backtrack(nums []int, target int, startIndex int, sum int, count *int) {
	if target == sum {
		*count++
	}


	for i:=startIndex;i<len(nums);i++ {
		backtrack(nums, target, i+1, sum + nums[i], count)
	}
}

func sum(nums []int) int {
	res := 0
	for i:=0;i<len(nums);i++ {
		res += nums[i]
	}
	return res
}
