package invert_binary_tree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Left = left
	root.Right = right
	return root
}
