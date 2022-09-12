package maximum_length_of_repeated_subarray

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]int, len(nums2))
	}
	// dp[0][..]
	for j:=0;j<len(nums2);j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
		}
	}
	// dp[..][0]
	for i:=0;i<len(nums1);i++{
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
	}

	for i := 1;i<len(nums1);i++{
		for j:=1;j<len(nums2);j++{
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	res := 0
	for i:=0;i<len(nums1);i++ {
		for j:=0;j<len(nums2);j++{
			res = max(res, dp[i][j])
		}
	}
	fmt.Println(dp)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//[1,2,3,2,1]
//[3,2,1,4,7]

//[0 0 1 0 0]
//[0 1 0 0 0]
//[1 0 0 0 0]
//[0 2 0 0 0] dp[3][0]
//[0 0 3 0 0]
