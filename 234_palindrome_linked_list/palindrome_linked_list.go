package palindrome_linked_list


type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	middle := findMiddle(head)
	left := head
	right := reverse(middle)

	for ; right != nil; {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func findMiddle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for ; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var left *ListNode
	left = nil
	cur := head
	for ; cur != nil; {
		right := cur.Next

		cur.Next = left
		left = cur
		cur = right
	}
	return left
}
