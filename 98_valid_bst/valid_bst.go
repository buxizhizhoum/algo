package valid_bst

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}


func traverse(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	left := traverse(root.Left, min, root)
	right := traverse(root.Right, root, max)
	return left && right
}
