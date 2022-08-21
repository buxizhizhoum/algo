package easier_solution

type ListNode struct {
	Val int
	Next *ListNode
}

// 从末尾算长度，根据最短的链表，把长的链表的指针移动到距离末尾相同的长度，然后开始判断节点是否相等
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for ;curA != nil; {
		curA = curA.Next
		lenA++
	}
	for ;curB != nil; {
		curB = curB.Next
		lenB++
	}
	gap := lenA - lenB
	curA, curB = headA, headB
	if gap > 0 {
		return FindIntersectionNode(curA, curB, gap)
	} else {
		return FindIntersectionNode(curB, curA, -gap)
	}

}

func FindIntersectionNode(curA *ListNode, curB *ListNode, gap int) *ListNode {
	for ;gap > 0; {
		curA = curA.Next
		gap--
	}

	for ;curA != nil; {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}