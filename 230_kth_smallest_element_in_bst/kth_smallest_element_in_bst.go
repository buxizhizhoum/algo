package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	rank := 0
	traverse(root, k, &res, &rank)
	return res
}


func traverse(root *TreeNode, k int, res *int, rank *int) {
	if root == nil{
		return
	}

	traverse(root.Left, k, res, rank)
	(*rank)++
	fmt.Println("rank: %s", *rank)
	if *rank == k {
		*res = root.Val
		fmt.Println(root.Val)
		return
	}
	traverse(root.Right, k, res, rank)
}

