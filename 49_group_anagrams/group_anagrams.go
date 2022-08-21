package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	cache := make(map[string][]string, 0)
	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		cache[key] = append(cache[key], str)
	}
	res := make([][]string, 0)
	for _, v := range cache {
		res = append(res, v)
	}
	return res
}

func main() {
	testStrs := []string{"eat","tea","tan","ate","nat","bat"}
	res := groupAnagrams(testStrs)
	fmt.Println(res)
}
