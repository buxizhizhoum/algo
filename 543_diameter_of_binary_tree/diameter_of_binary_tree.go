package diameter_of_binary_tree


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	diameter(root, &res)
	return res
}

// 返回节点最大深度
func diameter(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	// 左右节点最大深度
	leftDepth := diameter(root.Left, res)
	rightDepth := diameter(root.Right, res)
	// 当前节点的直径是左子树和右子树的和
	// 最大直径是节点左子树最大深度 + 右子树之中最大深度
	nodeDiameter := leftDepth + rightDepth
	// 取最大的diameter
	*res = max(*res, nodeDiameter)
	// 返回节点最大深度，后续算直径
	return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
