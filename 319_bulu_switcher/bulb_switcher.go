package main

import "math"

// 只有完全开方的数，会被按奇数次，否则都会按偶数次
// 比sqrt(n)小的几个完全平方数都是被按奇数次
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
