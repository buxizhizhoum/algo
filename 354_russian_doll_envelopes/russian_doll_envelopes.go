package main
// time limit exceeded
import (
	"fmt"
	"sort"
)
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			// 如果第一列相等，按照第二列逆序排
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})
	fmt.Println(envelopes)

	heights := make([]int, len(envelopes))
	for i, _ := range envelopes {
		heights[i] = envelopes[i][1]
	}

	return lengthOfLIS(heights)
}

func lengthOfLIS(nums []int) int {
	// dp代表以nums[i]结尾的最大递增子序列
	dp := make([]int, len(nums))
	for i:=0;i<len(dp);i++{
		dp[i] = 1
	}
	for i:=0;i<len(dp);i++{
		for j:=0;j<i;j++{
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	res := 1
	for i :=0; i< len(dp);i++{
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	//testNums := [][]int{
	//	{5,4},
	//	{6,4},
	//	{6,7},
	//	{2,3},
	//}
	testNums := [][]int{
		{30,50},{12,2},{3,4},{12,15},
	}
	maxEnvelopes(testNums)
}
