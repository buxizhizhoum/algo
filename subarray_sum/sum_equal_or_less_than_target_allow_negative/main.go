package main

import (
	"algo/utils"
	"fmt"
)

// 和小于target的最长连续数组下标，如果输入全部为正数，可以用滑动窗口实现

func maxSubArray(nums []int, target int) (int, int) {
	n := len(nums)
	// 前缀和
	preSum := make([]int, n + 1)
	for i:=1;i<=n;i++{
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	// dp[i][j]表示 nums[i][j]的最长子数组长度
	dp := make([][]int, n)
	for i:= 0; i<n;i++ {
		dp[i] = make([]int, n)
	}
	// 0结束，标明没有数组长度
	//dp[i][0] = 0
	// 0 开头，j结束，判断target
	//dp[0][j] =0
	if nums[0] <= target {
		dp[0][0] = 1
	}
	sum := 0
	for j:= 1;j<n;j++ {
		sum = preSum[j+1] - preSum[0]
		if sum <= target {
			dp[0][j] = dp[0][j-1] + 1
		} else {
			// 和已经大了，从0开始取，没法取了，长度设置成0
			dp[0][j] = 0
		}
	}
	for i:=1;i<n;i++ {
		for j:=i;j<n;j++ {
			// 取j, dp[i][j-1]
			lengthJ := 0
			sumJ := preSum[j%n] - preSum[i%n]
			if sumJ + nums[j%n] > target {
				// 取了长度就超了，最大子数组长度是0
				lengthJ = 0
			} else {
				lengthJ = dp[i][j-1] + 1
			}
			// 取i, dp[i-1][j]
			lengthI := 0
			sumI := preSum[j%n+1] - preSum[i%n-1]
			if sumI + nums[i%n] > target {
				lengthI = 0
			} else {
				lengthI = dp[i][j-1] + 1
			}
			dp[i][j] = max(lengthI, lengthJ)
		}
	}
	utils.PrintGrid(dp)
	// 寻找最长子数组的起止坐标
	maxLength := 0
	start, end := 0, 0
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if dp[i][j] > maxLength {
				start, end = i, j
				maxLength = dp[i][j]
			}
		}
	}
	return start, end
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	//testNums, target := []int{3,2,2,1,1,1,4}, 5
	//testNums, target := []int{1,2,6,6,2}, 5
	testNums, target := []int{1,2,-6,6,2}, 5
	start, end := maxSubArray(testNums, target)
	fmt.Println(start, end)
}