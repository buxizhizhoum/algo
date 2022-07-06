package level_traversal

import (
	"strconv"
	"strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {
	Sep string
	Empty string

}

func Constructor() Codec {
	return Codec{Sep: ",", Empty: "#"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.Empty
	}

	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for ; len(queue) > 0; {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			res = append(res, this.Empty)
		}
	}
	return strings.Join(res, this.Sep)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, this.Sep)
	if len(nodes) == 0 || nodes[0] == this.Empty{
		return nil
	}

	val, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: val}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i:=1; i < len(nodes); {
		node := queue[0]
		queue = queue[1:]

		left := nodes[i]
		i++
		right := nodes[i]
		i++

		if left != this.Empty {
			leftVal, _ := strconv.Atoi(left)
			leftNode := &TreeNode{
				Val: leftVal,
				Left: nil,
				Right: nil,
			}
			node.Left = leftNode
			queue = append(queue, leftNode)
		} else {
			node.Left = nil
		}

		if right != this.Empty {
			rightVal,_ := strconv.Atoi(right)
			rightNode := &TreeNode {
				Val: rightVal,
				Left: nil,
				Right: nil,
			}
			node.Right = rightNode

			queue = append(queue, rightNode)
		} else {
			node.Right = nil
		}
	}
	return root
}

