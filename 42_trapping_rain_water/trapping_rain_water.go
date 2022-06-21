package main

import "fmt"

func trap(height []int) int {
	maxL := maxLeft(height)
	maxR := maxRight(height)
	fmt.Println(maxL)
	fmt.Println(maxR)
	sum := 0
	for i:= 1; i < len(height) - 1; i++ {
		maxH := min(maxL[i], maxR[i])
		tmp := maxH - height[i]
		if tmp > 0 {
			sum += tmp
		}
	}
	return sum
}


func maxLeft(height []int) []int {
	res := make([]int, len(height))
	res[0] = 0
	for i:=1; i < len(height); i++ {
		res[i] = max(height[i-1], res[i-1])
		//if height[i-1] > tmp {
		//    res[i] = height[i-1]
		//    tmp = height[i-1]
		//} else {
		//    res[i] = tmp
		//}
	}
	return res
}

func maxRight(height []int) []int {
	res := make([]int, len(height))
	res[len(height) - 1] = 0
	for i:=(len(height) - 2); i >= 0; i-- {
		res[i] = max(height[i+1], res[i+1])
		//if height[i+1] > tmp {
		//   res[i] = height[i+1]
		//    tmp = height[i+1]
		//} else {
		//    res[i] = tmp
		//}
	}
	return res
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left := 0
	right := len(height) - 1
	maxLeft := height[0]
	maxRight := height[len(height) - 1]
	sum := 0
	for ; left <= right; {
		maxLeft = max(maxLeft, height[left])
		maxRight = max(maxRight, height[right])

		tmp := 0
		// 两侧低点出现在左侧，右侧不需要考虑了，直接根据左侧低点计算容量
		if maxLeft < maxRight {
			tmp = maxLeft - height[left]
			left++
		} else {
			tmp = maxRight - height[right]
			right--
		}
		sum += tmp
	}
	return sum
}


func main() {
	//height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	res := trap1(height)
	fmt.Println(res)
}
