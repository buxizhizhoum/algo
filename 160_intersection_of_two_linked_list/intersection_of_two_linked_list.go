package intersection_of_two_linked_list

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for ;p1 != p2; {
		// p1, p2跳到另一个链表的head, 不是p1.Next = headB，这样会把链表连起来，是不对的
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}

		if p1 == p2 {
			return p1
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return p1
}
