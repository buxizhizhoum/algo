package unique_bst

import "fmt"

func numTrees(n int) int {
	memo := make(map[string]int, 0)
	return count(1, n, memo)
}

func count(left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d:%d", left, right)
	cacheCount, ok := memo[key]
	if ok {
		return cacheCount
	}

	if left >= right {
		memo[key] = 1
		return 1
	}


	res := 0
	for i:= left; i <= right; i++ {
		leftCount := count(left, i-1, memo)
		rightCount := count(i+1, right, memo)
		res += leftCount * rightCount
	}

	memo[key] = res

	return res
}
