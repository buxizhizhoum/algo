package three_sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	cache := make(map[string]int, 0)

	res := make([][]int, 0)
	for i:=0;i<len(nums)-2;i++{
		if nums[i] > 0 {
			return res
		}

		left, right := i+1, len(nums) - 1
		for ;left < right; {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp > 0 {
				right--
			} else if tmp < 0 {
				left++
			} else if tmp == 0 {
				key := fmt.Sprintf("%d:%d:%d", nums[i], nums[left], nums[right])
				_, ok := cache[key]
				if !ok {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					cache[key] = 1
				}
				// 满足条件双指针同时收缩
				left++
				right--
			}
		}
	}
	return res
}
