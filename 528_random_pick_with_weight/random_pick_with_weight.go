// not ac
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	Num []int
	PreSum []int
}


func Constructor(w []int) Solution {
	// 用前缀和实现权重选择
	preSum := make([]int, len(w) + 1)
	preSum[0] = 0
	for i:=1;i<=len(w);i++ {
		preSum[i] = preSum[i-1] + w[i-1]
	}
	return Solution{
		Num: w,
		PreSum: preSum,
	}
}


func (this *Solution) PickIndex() int {
	// 从权重和索引[0, len(preSum) - 1]中随机选择一个数字，然后对应到原始数组的索引
	randNum := randInt(this.PreSum[1], this.PreSum[len(this.PreSum) - 1])
	// 从前缀和数组中找到这个数字对应的索引，如果找不到，找到比它大的最小的数字对应的索引
	preSumIndex := searchLeft(this.PreSum, randNum)
	// 前缀和多一个假元素，减去
	return preSumIndex - 1
}

// 在范围内取一个随机数
func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max - min + 1)
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right - left)
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if left > len(nums) {
		return -1
	}

	return left
}

//func searchLeft(nums []int, target int) int {
//	left := 0
//	right := len(nums)
//	for left < right {
//		mid := left + (right - left) / 2
//		if nums[mid] == target {
//			right = mid
//		} else if nums[mid] < target {
//			left = mid + 1
//		} else if nums[mid] > target {
//			right = mid
//		}
//	}
//	return left
//}


func main() {
	//w := []int{1,3}
	//w := []int{1}
	w := []int{3, 14, 1, 7}
	obj := Constructor(w)
	res:=obj.PickIndex()
	fmt.Println(res)
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
	fmt.Println(obj.PickIndex())
}
