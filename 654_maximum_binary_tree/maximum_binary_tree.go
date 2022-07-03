package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums) - 1)
}

func build(nums []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}
	index := maxIndex(nums, i, j)

	left := build(nums, i, index-1)
	right := build(nums, index+1, j)
	root := TreeNode{
		Val: nums[index],
		Left: left,
		Right: right,
	}
	return &root
}

func maxIndex(nums []int, start, end int) int {
	if start > end {
		return -1
	}
	index := start
	for i:=start; i<=end; i++{
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}


func main() {
	testNums := []int{3,2,1,6,0,5}
	constructMaximumBinaryTree(testNums)
}