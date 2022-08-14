package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
	//fmt.Println("reverse")

	dummy := &ListNode{
		Val: 0,
		Next: nil,
	}
	cur := dummy
	val, carry := 0, 0
	for ;l1 != nil && l2 != nil; {
		//fmt.Println("for add two")
		tmpSum := l1.Val + l2.Val + carry
		val, carry = tmpSum % 10, tmpSum / 10
		tmpNode := &ListNode{
			Val: val,
			Next: nil,
		}
		cur.Next = tmpNode
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for ;l1 != nil; {
		tmpSum := l1.Val + carry
		val, carry = tmpSum % 10, tmpSum / 10
		tmpNode := &ListNode{
			Val: val,
			Next: nil,
		}
		cur.Next = tmpNode
		cur = cur.Next
		l1 = l1.Next
	}
	for ;l2 != nil; {
		tmpSum := l2.Val + carry
		val, carry = tmpSum % 10, tmpSum / 10
		tmpNode := &ListNode{
			Val: val,
			Next: nil,
		}
		cur.Next = tmpNode
		cur = cur.Next
		l2 = l2.Next
	}
	// 最后的进位处理
	if carry != 0 {
		tmpNode := &ListNode{
			Val:  carry,
			Next: nil,
		}
		cur.Next = tmpNode
		cur = cur.Next
	}
	return reverse(dummy.Next)
}


func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	pre = nil
	cur := head
	for ;cur != nil ; {
		next := cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func printLinkedList(head *ListNode) {
	cur := head
	for ;cur != nil; {
		fmt.Print(cur.Val)
		cur = cur.Next
	}
}

func main() {
//[7,2,4,3]
//	[5,6,4]
//	l1 := &ListNode{
//		Val:  7,
//		Next: &ListNode{
//			Val:  2,
//			Next: &ListNode{
//				Val:  4,
//				Next: &ListNode{
//					Val:  3,
//					Next: nil,
//				},
//			},
//		},
//	}
//	l2:=&ListNode{
//		Val:  5,
//		Next: &ListNode{
//			Val:  6,
//			Next: &ListNode{
//				Val:  4,
//				Next: nil,
//			},
//		},
//	}
	l1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	res := addTwoNumbers(l1, l2)
	printLinkedList(res)
}

