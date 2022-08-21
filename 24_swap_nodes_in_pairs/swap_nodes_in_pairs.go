package swap_nodes_in_pairs

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val: 0,
		Next: head,
	}
	cur := dummyNode
	for ;cur.Next != nil && cur.Next.Next != nil; {
		curNext := cur.Next
		curNextNextNext := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		// cur的下一个节点的下一个节点指向cur的下一个节点，下一个节点已经更新了，所以要提前存起来
		cur.Next.Next = curNext
		cur.Next.Next.Next = curNextNextNext
		// cur dummy 在要转换的节点前面
		cur = cur.Next.Next
	}
	return dummyNode.Next
}
