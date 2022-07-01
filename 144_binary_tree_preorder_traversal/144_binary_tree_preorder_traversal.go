package binary_tree_preorder_traversal


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	traverse(root.Left, res)
	traverse(root.Right, res)
}
