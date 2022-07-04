package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func buildTree(inorder []int, postorder []int) *TreeNode {
	val2index := make(map[int]int, 0)
	for i, val := range inorder {
		val2index[val] = i
	}

	return build(inorder, 0, len(inorder) - 1, postorder, 0, len(postorder) - 1, val2index)


}

func build(inOrder []int, inOrderStart, inOrderEnd int,
	postOrder []int, postOrderStart, postOrderEnd int, val2index map[int]int) *TreeNode {
	if inOrderStart > inOrderEnd {
		return nil
	}

	rootVal := postOrder[postOrderEnd]
	index, _ := val2index[rootVal]

	leftSize := index - inOrderStart

	left := build(inOrder, inOrderStart, index-1,
		postOrder, postOrderStart, postOrderStart + leftSize - 1, val2index)
	right := build(inOrder, index+1, inOrderEnd,
		postOrder, postOrderStart + leftSize, postOrderEnd - 1, val2index)

	return &TreeNode{
		Val: rootVal,
		Left: left,
		Right: right,
	}
}


func main() {
	testInorder := []int{9,3,15,20,7}
	testPostOrder := []int{9,15,7,20,3}
	res := buildTree(testInorder, testPostOrder)
	fmt.Println(res)

}
