package main

import "fmt"

func foo(nums []int) []int {
	nums = nums[1:]
	return nums
}

func p(a interface{}) string {
	res := fmt.Sprintf("%v", a)
	return res
}

func main() {
	nums := []int{1,2,3}
	nums = append(nums, 4)
	res := foo(nums)
	nums = nums[:len(nums) - 1]
	nums = append(nums, 5)
	fmt.Println(res)
	a := p(1)
	fmt.Println(a)
}


