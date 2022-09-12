package uncrossed_lines

// 求不想交的直线最大的条数，实际就是求最长公共子序列
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	//dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
	dp := make([][]int, len(nums1) + 1)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int, len(nums2) + 1)
	}
	for i:=1;i<len(nums1) + 1;i++ {
		for j:=1;j<len(nums2) + 1;j++{
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
