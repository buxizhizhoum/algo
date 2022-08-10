package main

import "fmt"

// 求子数组下标，要求子数组的和小于等于target，数组全部为正数

func maxSubArrayLen(nums []int, target int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	// 左闭右开区间
	left, right := 0, 0
	windowSum := 0
	maxLen := 0
	resLeft, resRight := left, right
	for ;right < len(nums); {
		windowSum += nums[right]
		if windowSum <= target {
			if right - left > maxLen {
				resLeft, resRight = left, right
				maxLen = right - left
			}
		}
		// 左闭右开，上界在处理完长度后再修改
		right++

		for ;windowSum > target; {
			windowSum -= nums[left]
			left++
			if windowSum <= target {
				if right - left > maxLen {
					resLeft, resRight = left, right
					maxLen = right - left
				}
			}
		}
	}
	return resLeft, resRight
}

func main() {
	testNums, target := []int{3,2,2,1,1,1,4}, 5
	//testNums, target := []int{1,2,6,6,2}, 5
	start, end := maxSubArrayLen(testNums, target)
	fmt.Println(start, end)
}


