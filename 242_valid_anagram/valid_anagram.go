package valid_anagram

func isAnagram(s string, t string) bool {
	cache := make(map[byte]int, 0)
	for i:=0;i<len(s);i++ {
		cache[s[i]]++
	}
	for i:=0;i<len(t);i++ {
		cache[t[i]]--
	}
	for _, v := range cache {
		if v != 0 {
			return false
		}
	}
	return true
}
