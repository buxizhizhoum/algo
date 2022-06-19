package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 2
	fast := 2
	for ; fast < len(nums); {
		if nums[slow - 2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}


func main() {
	testNums := []int{0,0,1,1,1,1,2,3,3}
	//testNums := []int{1,2,2}
	k := removeDuplicates(testNums)
	fmt.Println(testNums[:k])
}
