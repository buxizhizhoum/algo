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


func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0

	for p1:= headA;p1 != nil; p1 = p1.Next {
		lenA += 1
	}
	for p2:= headB;p2 != nil; p2 = p2.Next {
		lenB += 1
	}
	// 哪个长，先把长的部分走完
	p1 := headA
	p2 := headB
	if lenA > lenB {
		for i := 0; i < lenA - lenB; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < lenB - lenA; i++ {
			p2 = p2.Next
		}
	}

	for ;p1 != p2; {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
