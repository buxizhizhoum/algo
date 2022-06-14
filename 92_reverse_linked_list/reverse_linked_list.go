package reverse_linked_list


type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	// 从next节点，翻转left - 1, right -1 之间的点
	// 这部分拿到的时候就已经是翻转好的
	reversed := reverseBetween(head.Next, left - 1, right - 1)
	head.Next = reversed
	return head
}

var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1{
		successor = head.Next
		return head
	}
	// 从next节点，翻转n - 1个节点
	// 这部分拿到的时候，只返回了翻转后的头结点，还需进行拼接
	last := reverseN(head.Next, n - 1)
	// 拼接
	head.Next.Next = head
	head.Next = successor
	return last
}
