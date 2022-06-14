package reverse_linked_list


type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rest := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}


func reverseN(head *ListNode, n int) *ListNode {
	reversed, _ := _reverseN(head.Next, n - 1)
	return reversed
}


func _reverseN(head *ListNode, n int) (*ListNode, *ListNode) {
	if n == 1 {
		successor := head.Next
		return head, successor
	}
	reversed, successor := _reverseN(head.Next, n - 1)
	head.Next.Next = head
	head.Next = successor
	return reversed, successor
}
