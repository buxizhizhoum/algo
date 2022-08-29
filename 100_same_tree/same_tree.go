package main

import (
	"fmt"
	"reflect"
)

//Definition for a binary tree node.
type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 注意，这不能直接用==判断，需要根据val相等判断
	if (p != nil && q == nil) || (p == nil && q != nil ) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Left == nil && q.Left != nil {
		return false
	}
	if p.Left != nil && q.Left == nil {
		return false
	}
	if p.Right == nil && q.Right != nil {
		return false
	}
	if p.Right != nil && q.Right == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right
}

func main() {
	p := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	q := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	res := isSameTree(p, q)
	fmt.Println(res)
	//tmp := TreeNode{
	//	Val:   0,
	//	Left:  nil,
	//	Right: nil,
	//}
	fmt.Println(TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	} == TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(p == q)
	fmt.Println(reflect.DeepEqual(p, q))
}
