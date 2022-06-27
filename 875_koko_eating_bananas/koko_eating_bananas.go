package koko_eating_bananas


// https://www.jianshu.com/p/20eaefda7564
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0
	for i:=0;i<len(piles);i++{
		if piles[i] > right {
			right = piles[i]
		}
	}

	for left < right {
		mid := left + (right - left) / 2
		hour := eatTime(piles, mid)
		if hour == h {
			right = mid
		} else if hour < h {
			right = mid
		} else if hour > h {
			left = mid + 1
		}
	}
	return left
}

func eatTime(piles []int, speed int) int {
	res := 0
	for i:=0;i<len(piles);i++{
		res += piles[i] / speed
		if piles[i] % speed > 0 {
			res++
		}
	}
	return res
}
