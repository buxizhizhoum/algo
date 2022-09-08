package last_stone_weight
// https://programmercarl.com/1049.%E6%9C%80%E5%90%8E%E4%B8%80%E5%9D%97%E7%9F%B3%E5%A4%B4%E7%9A%84%E9%87%8D%E9%87%8FII.html#_1049-%E6%9C%80%E5%90%8E%E4%B8%80%E5%9D%97%E7%9F%B3%E5%A4%B4%E7%9A%84%E9%87%8D%E9%87%8F-ii
// 本题其实就是尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了
func lastStoneWeightII(stones []int) int {
	total := sum(stones)
	dp := make([][]int, len(stones))
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, total / 2 + 1)
	}

	// dp[0][j]
	for j:=0;j<total / 2 + 1;j++ {
		if j >= stones[0] {
			dp[0][j] = stones[0]
		}
	}

	// dp[i][0] = 0
	// 容量为0，初始化为0
	for i:=0;i<len(stones);i++{
		dp[i][0] = 0
	}

	for i:=1;i<len(stones);i++ {
		for j:=1;j<total /2 + 1;j++ {
			if j - stones[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]] + stones[i])
			}
		}
	}

	partA := dp[len(stones) - 1][total / 2]
	partB := total - partA
	if partA > partB {
		return partA - partB
	} else {
		return partB - partA
	}
}

func sum(nums []int) int {
	res := 0
	for i:=0;i<len(nums);i++ {
		res += nums[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
