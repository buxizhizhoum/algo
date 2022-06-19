package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
		Next: nil,
	}
	dummy.Next = head
	slow := dummy
	fast := head
	for ; fast!= nil; {
		for ; fast.Next != nil && fast.Val == fast.Next.Val; {
			fast = fast.Next
		}

		if slow.Next == fast {
			slow = slow.Next
			fast = fast.Next
		} else {
			slow.Next = fast.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}
