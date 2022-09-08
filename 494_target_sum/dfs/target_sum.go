package dfs


func findTargetSumWays(nums []int, target int) int {
	res := 0
	backtrack(nums, 0, target, 0, &res)
	return res
}

func backtrack(nums []int, startIndex int, target int, sum int, count *int) {
	if startIndex == len(nums) {
		if target == sum {
			*count++
		}
		return
	}

	backtrack(nums, startIndex+1, target, sum + nums[startIndex], count)
	backtrack(nums, startIndex+1, target, sum - nums[startIndex], count)
}
