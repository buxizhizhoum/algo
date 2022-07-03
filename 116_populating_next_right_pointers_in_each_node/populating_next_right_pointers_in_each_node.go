package populating_next_right_pointers_in_each_node

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}


func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

func traverse(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	// 连接两个节点
	node1.Next = node2
	// 两个节点下一层节点连接
	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	traverse(node1.Right, node2.Left)
}
