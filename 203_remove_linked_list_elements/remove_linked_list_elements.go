package main

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// todo: 什么时候用nil, 什么时候不用
	// 翻转链表的时候，dummyNode是nil
	dummyNode := &ListNode{
		Val: 0,
		Next: head,
	}
	pre := dummyNode
	cur := head
	for ;cur != nil; {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			// 要先更新前面的
			pre = cur
			cur = cur.Next
		}
	}
	return dummyNode.Next
}

func printLinkedList(head *ListNode) {
	cur := head
	for ;cur != nil; cur = cur.Next {
		//fmt.Println(cur.Val)
	}
}


func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	res := removeElements(head, 6)
	printLinkedList(res)
}