package best_time_to_buy_and_sell_stock

const INT_MAX = int(^uint(0) >> 1)

func maxProfit(prices []int) int {
	minPrices := INT_MAX
	res := 0
	for i:= 0;i<len(prices);i++ {
		minPrices = min(minPrices, prices[i])
		res = max(res, prices[i] - minPrices)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
