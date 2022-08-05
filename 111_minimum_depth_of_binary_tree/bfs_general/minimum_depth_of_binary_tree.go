package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


type Element struct {
	Node *TreeNode
	Level int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*Element, 0)
	queue = append(queue, &Element{
		Node:  root,
		Level: 0,
	})
	step := 1
	// 需要逐层弹出，可以入队的时候附带一个层信息，也可以把这一次层节点个数记住，
	// 先用加上层级信息的方式
	var element *Element
	for ;len(queue) > 0;{
		element, queue = queue[0], queue[1:]
		if element.Node.Left != nil {
			queue = append(queue, &Element{
				Node:  element.Node.Left,
				Level: element.Level + 1,
			})
		}
		if element.Node.Right != nil {
			queue = append(queue, &Element{
				Node:  element.Node.Right,
				Level: element.Level + 1,
			})
		}
		// 如果已经到了下一层了，层数加1
		if element.Level > step {
			step += 1
		}
		// 表明有一个子树到底了
		if element.Node.Left == nil && element.Node.Right == nil {
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
