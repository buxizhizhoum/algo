package reverse_linked_list

import "fmt"

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


func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}
	var dummy *ListNode
	dummy = nil
	left := dummy
	cur := head
	right := head.Next
	for ; cur.Next != nil; {
		cur.Next = left
		left = cur

		cur = right
		right = right.Next

	}
	cur.Next = left
	return cur
}


func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy *ListNode
	dummy = nil
	left := dummy
	cur := head
	for ; cur != nil; {
		right := cur.Next
		// 换向
		cur.Next = left
		// 窗口前移
		left = cur
		cur = right
	}
	return left
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := reverseList(head)
	fmt.Println(res)
}
