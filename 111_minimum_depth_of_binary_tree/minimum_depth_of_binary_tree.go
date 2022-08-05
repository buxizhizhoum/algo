package minimum_depth_of_binary_tree

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	// 空的子树不算高度
	if root.Left == nil {
		return rightDepth + 1
	}
	if root.Right == nil {
		return leftDepth + 1
	}
	return min(leftDepth, rightDepth) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
