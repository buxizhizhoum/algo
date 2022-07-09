package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return traverse(root, val)
	// 在root == nil的时候，root是新建的，直接返回原来的root不行
	// return root
}

func traverse(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
			Left: nil,
			Right: nil,
		}
	}

	if val < root.Val {
		root.Left = traverse(root.Left, val)
	} else if val > root.Val {
		root.Right = traverse(root.Right, val)
	}
	return root
}


func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	insertIntoBST(root, 5)
}