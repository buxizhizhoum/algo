package reverse_nodes_in_k_group

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a := head
	b := head
	for i:=0;i< k;i++{
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead


}


func reverse(a *ListNode, b *ListNode) *ListNode {
	var left *ListNode
	left = nil
	cur := a
	right := a
	for ;cur != b; {
		right = cur.Next

		cur.Next = left
		left = cur
		cur = right
	}
	// 为什么这里返回left
	return left
}