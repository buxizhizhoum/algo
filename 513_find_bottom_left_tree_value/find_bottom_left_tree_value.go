package find_bottom_left_tree_value

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var result int = 0
	var top *TreeNode
	for ;len(queue) > 0; {
		size := len(queue)
		for i:= 0;i<size;i++ {
			top = queue[0]
			queue = queue[1:]
			if i == 0{
				result = top.Val
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return result
}
