package binary_tree_paths

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}

func binaryTreePaths(root *TreeNode) []string {
	res := make([][]int, 0)
	track := make([]int, 0)
	treePath(root, track, &res)
	result := make([]string, len(res))
	for i, val := range res {
		tmp := convertString(val)
		result[i] = tmp
	}
	return result
}

func treePath(root *TreeNode, track []int, result *[][]int) {
	if root.Left == nil && root.Right == nil {
		track = append(track, root.Val)
		tmp := copySlice(track)
		*result = append(*result, tmp)
	}
	if root.Left != nil {
		// 回溯
		track = append(track, root.Val)
		treePath(root.Left, track, result)
		track = track[:len(track) - 1]
	}
	if root.Right != nil {
		track = append(track, root.Val)
		treePath(root.Right, track, result)
		track = track[:len(track) - 1]
	}
}

func convertString(track []int) string {
	tmp := make([]string, len(track))
	for i, val := range track {
		tmp[i] = strconv.Itoa(val)
	}
	return strings.Join(tmp, "->")
}

func copySlice(track []int) []int {
	res := make([]int, len(track))
	for i, val := range track {
		res[i] = val
	}
	return res
}
