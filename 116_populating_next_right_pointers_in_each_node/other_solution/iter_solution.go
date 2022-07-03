package other_solution

type Node struct {
     Val int
     Left *Node
     Right *Node
     Next *Node
}


type NodeLevel struct {
	Node *Node
	Level int
}

func connect(root *Node) *Node {
	queue := make([]NodeLevel, 0)
	if root == nil {
		return nil
	}

	queue = append(queue, NodeLevel{
		Node: root,
		Level: 0,
	})

	for i:=0;i<len(queue);i++{
		nodeLevel := queue[i]

		if nodeLevel.Node.Left != nil && nodeLevel.Node.Right != nil {
			queue = append(queue, NodeLevel{
				Node: nodeLevel.Node.Left,
				Level: nodeLevel.Level + 1,
			})

			queue = append(queue, NodeLevel{
				Node: nodeLevel.Node.Right,
				Level: nodeLevel.Level + 1,
			})
		}
	}

	for i:=0;i<len(queue) - 1;i++ {
		cur := queue[i]
		next := queue[i+1]
		if cur.Level == next.Level {
			cur.Node.Next = next.Node
		}
	}
	return queue[0].Node
}
