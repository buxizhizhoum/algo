package recursion_memo

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make(map[string]int, 0)
	return recursion(word1, m-1, word2, n-1, memo)
}

func recursion(word1 string, i int, word2 string, j int, memo map[string]int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	key := fmt.Sprintf("%s:%s", i, j)
	if val, ok := memo[key]; ok {
		return val
	}

	var res int
	if word1[i] == word2[j] {
		res = recursion(word1, i-1, word2, j-1, memo)
		memo[key] = res
		return res
	} else {
		del := recursion(word1, i-1, word2, j, memo)
		insert := recursion(word1, i, word2, j-1, memo)
		replace := recursion(word1, i-1, word2, j-1, memo)
		// 需要加1
		res = min(min(del, insert), replace) + 1
		memo[key] = res
		return res
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
