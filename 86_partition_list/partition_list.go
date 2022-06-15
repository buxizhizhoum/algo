package partition_list


type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummySmaller := &ListNode{
		Val: -1,
		Next: nil,
	}
	dummyLarger := &ListNode{
		Val: -1,
		Next: nil,
	}

	p1 := dummySmaller
	p2 := dummyLarger
	p := head
	for ;p != nil; {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 断开原链表每个节点的连接
		tmp := p.Next
		p.Next = nil
		p = tmp
	}

	p1.Next = dummyLarger.Next
	return dummySmaller.Next

}