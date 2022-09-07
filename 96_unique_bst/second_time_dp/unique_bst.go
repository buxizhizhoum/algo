package second_time_dp

// https://programmercarl.com/0096.%E4%B8%8D%E5%90%8C%E7%9A%84%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.html#%E6%80%9D%E8%B7%AF
func numTrees(n int) int {
	// dp[i] ： i个连续节点组成的二叉搜索树的个数为dp[i]。
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i:=2;i<=n;i++{
		// j代表头结点
		for j:=1;j<=i;j++{
			// j-1 为j为头结点左子树节点数量，i-j 为以j为头结点右子树节点数量
			dp[i] += dp[j-1] * dp[i - j]
		}
	}
	return dp[n]
}
