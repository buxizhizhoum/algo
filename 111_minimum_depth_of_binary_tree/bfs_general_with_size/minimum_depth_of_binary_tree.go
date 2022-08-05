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
	// 需要逐层弹出，可以入队的时候附带一个层信息，也可以把这一次层节点个数记住，
	// 先用加上层级信息的方式
	var node *TreeNode
	for ;len(queue) > 0;{
		qSize := len(queue)
		for i:=0;i<qSize;i++ {
			node, queue = queue[0], queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			// 表明有一个子树到底了
			if node.Left == nil && node.Right == nil {
				break
			}
		}
		step += 1
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
