package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {
	Spliter string
	EmptyNode string
}

func Constructor() Codec {
	return Codec{Spliter: ",", EmptyNode: "#"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := make([]string, 0)
	this._serialize(root, &res)
	return strings.Join(res, ",")
}

func (this *Codec) _serialize(root *TreeNode, res *[]string) {
	if root == nil {
		*res = append(*res, this.EmptyNode)
		return
	}
	*res = append(*res, strconv.Itoa(root.Val))

	this._serialize(root.Left, res)
	this._serialize(root.Right, res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, this.Spliter)
	return this._deserialize(&nodes)
}

func(this *Codec) _deserialize(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == this.EmptyNode {
		return nil
	}

	left := this._deserialize(nodes)
	right := this._deserialize(nodes)
	val, _ := strconv.Atoi(first)
	return &TreeNode{
		Val: val,
		Left: left,
		Right: right,
	}
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	fmt.Println(data)
	ans := deser.deserialize(data)
	fmt.Println(ans)
}
