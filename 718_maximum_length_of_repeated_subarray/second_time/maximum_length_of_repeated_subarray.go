package second_time

func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
	dp := make([][]int, len(nums1) + 1)
	for i:= 0;i<len(dp);i++ {
		dp[i] = make([]int, len(nums2) + 1)
	}
	// dp[0][..] = 0
	// dp[..][0] = 0
	for i := 1;i<len(nums1) + 1;i++{
		for j:=1;j<len(nums2) + 1;j++{
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	res := 0
	for i:=0;i<len(nums1) + 1;i++ {
		for j:=0;j<len(nums2) + 1;j++{
			res = max(res, dp[i][j])
		}
	}
	//fmt.Println(dp)
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
