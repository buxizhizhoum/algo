package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	val2index := make(map[int]int, 0)
	for i, val := range postorder {
		val2index[val] = i
	}
	return build(preorder, 0, len(preorder) - 1, postorder, 0, len(postorder) - 1, val2index)

}

func build(preOrder []int, preOrderStart, preOrderEnd int,
	postOrder []int, postOrderStart, postOrderEnd int, val2index map[int]int) *TreeNode {
	if preOrderStart > preOrderEnd {
		return nil
	}
	if preOrderStart == preOrderEnd {
		return &TreeNode{
			Val: preOrder[preOrderStart],
			Left: nil,
			Right: nil,
		}
	}
	rootVal := preOrder[preOrderStart]
	// 这里是要用左子树的根节点去分隔后序遍历数组
	leftSubTreeRootVal := preOrder[preOrderStart + 1]
	index := val2index[leftSubTreeRootVal]

	leftSize := index - postOrderStart + 1

	left := build(preOrder, preOrderStart + 1, preOrderStart + leftSize,
		postOrder, postOrderStart, index, val2index)
	right := build(preOrder, preOrderStart + leftSize + 1, preOrderEnd,
		postOrder, index + 1, postOrderEnd - 1, val2index)

	return &TreeNode{
		Val: rootVal,
		Left: left,
		Right: right,
	}
}


func main() {
	testInorder := []int{1,2,4,5,3,6,7}
	testPostOrder := []int{4,5,2,6,7,3,1}
	constructFromPrePost(testInorder, testPostOrder)
}