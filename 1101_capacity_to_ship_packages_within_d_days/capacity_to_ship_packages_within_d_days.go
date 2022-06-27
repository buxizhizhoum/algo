package capacity_to_ship_packages_within_d_days

//https://www.jianshu.com/p/20eaefda7564
func shipWithinDays(weights []int, days int) int {
	left := 0
	right := 0
	for i:=0;i<len(weights);i++{
		if weights[i] > left {
			left = weights[i]
		}
		right += weights[i]
	}

	for left < right {
		mid := left + (right - left) / 2
		tmp := shipDays(weights, mid)
		if tmp == days {
			right = mid
		} else if tmp < days {
			right = mid
		} else if tmp > days {
			left = mid + 1
		}
	}
	return left
}

func shipDays(weights []int, capacity int) int {
	res := 0
	for i:=0;i<len(weights); {
		tmp := capacity
		// 连续装，直到装不下的时候，运载天数+1
		for i < len(weights) {
			tmp -= weights[i]
			if tmp < 0 {
				break
			}
			i++
		}
		res++
	}
	return res
}
