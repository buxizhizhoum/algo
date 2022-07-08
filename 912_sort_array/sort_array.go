package main

import "fmt"

func sortArray(nums []int) []int {
	temp := make([]int, len(nums))
	sortSlice(nums, 0, len(nums) - 1, temp)
	return nums
}


func sortSlice(nums []int, left, right int, temp []int) {
	if left == right {
		return
	}
	mid := left + (right - left) / 2

	sortSlice(nums, left, mid, temp)
	sortSlice(nums, mid + 1, right, temp)

	merge(nums, left, mid, right, temp)

}


func merge(nums []int, left, mid, right int, temp []int) {
	for i := left; i<=right;i++ {
		temp[i] = nums[i]
	}

	i := left
	j := mid + 1

	for p := left; p <= right; p++ {
		if i == mid + 1 {
			nums[p] = temp[j]
			j++
		} else if j == right + 1 {
			nums[p] = temp[i]
			i++
		} else if temp[i] > temp[j] {
			nums[p] = temp[j]
			j++
		} else if nums[i] <= temp[j] {
			nums[p] = temp[i]
			i++
		}
	}
}


func main() {
	testArr := []int{5,1,1,2,0,0}
	res := sortArray(testArr)
	fmt.Println(res)
}
