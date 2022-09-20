package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	return recursion(word1, m-1, word2, n-1)
}

func recursion(word1 string, i int, word2 string, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}


	if word1[i] == word2[j] {
		return recursion(word1, i-1, word2, j-1)
	} else {
		del := recursion(word1, i-1, word2, j)
		insert := recursion(word1, i, word2, j-1)
		replace := recursion(word1, i-1, word2, j-1)
		// 需要加1
		return min(min(del, insert), replace) + 1
	}
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minDistance("a", "ab")
}
