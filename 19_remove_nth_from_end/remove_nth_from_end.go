package remove_nth_from_end

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val: -1,
		Next: nil,
	}
	dummy.Next = head

	node := findFromEnd(dummy, n+1)
	node.Next = node.Next.Next
	return dummy.Next
}


func findFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for ;n > 0; n-- {
		fast = fast.Next
	}

	for ;fast != nil; {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
