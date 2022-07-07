package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	dupTrees := make([]*TreeNode, 0)
	memo := make(map[string]int, 0)
	traverse(root, memo, &dupTrees)
	return dupTrees
}


func traverse(root *TreeNode, memo map[string]int, dupTrees *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left, memo, dupTrees)
	right := traverse(root.Right, memo, dupTrees)

	subTree := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
	//fmt.Println(subTree)
	count, _ := memo[subTree]
	if count == 1{
		*dupTrees = append(*dupTrees, root)
		memo[subTree] += 1
	} else {
		memo[subTree] += 1
	}
	return subTree
}
