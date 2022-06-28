package main

import (
	"fmt"
	"sort"
)

type Element struct {
	index int
	value int
}

func advantageCount(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums2))
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	tmp := make([]Element, len(nums2))
	for i:=0;i<len(nums2);i++ {
		tmp[i] = Element{
			index: i,
			value: nums2[i],
		}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].value < tmp[j].value
	})

	left := 0
	right := len(tmp) - 1
	for i:=len(tmp) - 1;i>=0;i-- {
		index := tmp[i].index
		maxVal := tmp[i].value
		// 如果目前最大的数可以压过对方最大的数，则压过，否则用最小的数顶替
		if nums1[right] > maxVal {
			res[index] = nums1[right]
			right--
		} else {
			res[index] = nums1[left]
			left++
		}
	}
	return res
}

//[2,7,11,15] [1,10,4,11] 测试结果:[15,11,7,2] 期望结果:[2,11,7,15]
func main() {
	nums1 := []int{2,7,11,15}
	nums2 := []int{1,10,4,11}
	res := advantageCount(nums1, nums2)
	fmt.Println(res)
}
