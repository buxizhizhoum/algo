package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1
	for cur:=root.Left; cur != nil; cur=cur.Left {
		left++
	}
	right:= 1
	for cur:=root.Right; cur != nil; cur=cur.Right{
		right++
	}
	// 满二叉树
	if left == right {
		return int(math.Pow(2, float64(left)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}


