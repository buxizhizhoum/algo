package balanced_binary_tree

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res := maxHeight(root)
	return res != -1
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxHeight(root.Left)
	right := maxHeight(root.Right)
	if left == -1 || right == -1 {
		return -1
	}
	if int(math.Abs(float64(left) - float64(right))) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


