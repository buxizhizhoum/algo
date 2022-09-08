package main

import "fmt"
// 能不能分隔成等和子集，转化成能否找到元素，装满sum / 2的背包
func canPartition(nums []int) bool {
	total := sum(nums)
	if total % 2 != 0 {
		return false
	}
	dp := make([][]int, len(nums))
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]int, total / 2 + 1)
	}

	//dp[0][..]
	// 取第0个元素，如果第j列的容量能放下，则是第0个元素的大小，否则是0
	for j:= 0;j<total /2 + 1;j++ {
		if j >= nums[0] {
			dp[0][j] = nums[0]
		}
	}
	//dp[..][0] == 0
	// 容量为0,初始化为0
	for i:=0;i<len(nums);i++{
		dp[i][0] = 0
	}

	for i:=1;i<len(nums);i++ {
		for j:=1;j< total / 2 + 1; j++ {
			// 放不下
			if j - nums[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]] + nums[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums) - 1][total / 2] == total / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(nums []int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}

func main() {
	testNums := []int{1,5,11,5}
	res := canPartition(testNums)
	fmt.Println(res)
}
