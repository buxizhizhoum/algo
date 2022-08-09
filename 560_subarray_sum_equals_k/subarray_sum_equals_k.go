package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums) + 1)
	for i:=1;i<len(nums) + 1;i++{
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	// 前缀和出现的次数，因为会有负数，所以会存在重复的情况，每个前缀和可能出现多次
	sumCount := make(map[int]int, 0)
	sumCount[0] = 1

	count := 0
	for i:= 0;i<len(nums);i++{
		curSum := preSum[i+1]
		needSum := curSum - k
		val, ok := sumCount[needSum]
		if ok {
			count += val
		}
		sumCount[curSum] += 1
	}
	return count
}

func main() {
	//testNums, k := []int{1,2,3}, 3
	testNums, k := []int{1,1,1}, 2
	//testNums, k := []int{1}, 0
	res := subarraySum(testNums, k)
	fmt.Println(res)

}
