package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	val2index := make(map[int]int, 0)
	for i, item := range inorder {
		val2index[item] = i
	}
	return build(
		preorder, 0, len(preorder) - 1,
		inorder, 0, len(inorder) - 1, val2index)
}

func build(preOrder []int, preOrderStart, preOrderEnd int,
	inOrder []int, inOrderStart, inOrderEnd int, val2index map[int]int) *TreeNode {

	if preOrderStart > preOrderEnd {
		return nil
	}
	rootVal := preOrder[preOrderStart]
	rootIndex, _ := val2index[rootVal]
	leftTreeSize := rootIndex - inOrderStart

	left := build(
		preOrder, preOrderStart+1, preOrderStart + leftTreeSize,
		inOrder, inOrderStart, rootIndex - 1, val2index)
	right := build(
		preOrder, preOrderStart + leftTreeSize + 1, preOrderEnd,
		inOrder, rootIndex + 1, inOrderEnd, val2index)

	root := TreeNode{
		Val: rootVal,
		Left: left,
		Right: right,
	}
	return &root
}
