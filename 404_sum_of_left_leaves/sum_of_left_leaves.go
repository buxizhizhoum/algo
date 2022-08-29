package sum_of_left_leaves

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := sumOfLeftLeaves(root.Left)
	right := sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + left + right
	}
	return left + right
}
