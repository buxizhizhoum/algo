package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	step := 1
	root.Val = 1
	// 需要逐层弹出，可以入队的时候附带一个层信息，也可以把这一次层节点个数记住，
	// 这里我们投机取巧，用root.Val记录层信息
	var node *TreeNode
	for ;len(queue) > 0;{
		node, queue = queue[0], queue[1:]
		if node.Left != nil {
			node.Left.Val = node.Val + 1
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Val = node.Val + 1
			queue = append(queue, node.Right)
		}
		// 如果已经到了下一层了，层数加1
		if node.Val > step {
			step += 1
		}
		if node.Left == nil && node.Right == nil {
			break
		}
	}
	return step
}


func main() {
	//root := &TreeNode{
	//	Val:   13,
	//	Left:  &TreeNode{
	//		Val:   9,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   20,
	//		Left:  &TreeNode{
	//			Val:   15,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	root := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	res := minDepth(root)
	fmt.Println(res)
}
