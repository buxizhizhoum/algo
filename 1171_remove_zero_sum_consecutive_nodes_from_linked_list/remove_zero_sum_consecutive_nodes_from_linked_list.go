package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
		Next: head,
	}
	curSum := 0
	preSum2Node := make(map[int]*ListNode, 0)
	preSum2Node[0] = dummy
	cur := dummy
	for ;cur != nil; {
		curSum += cur.Val
		// 这里后面的前缀和会覆盖前面的，最终会存储最后一次出现的前缀和
		preSum2Node[curSum] = cur
		cur = cur.Next
	}
	cur = dummy
	curSum = 0
	for ;cur != nil; {
		curSum += cur.Val
		endNode, ok := preSum2Node[curSum]
		// 前缀和前面的接到后面，实际上这里每次都是ok, 对于相同的节点，这样操作相当于没有操作
		if ok {
			cur.Next = endNode.Next
		}
		cur = cur.Next
	}

	return dummy.Next
}

func printLinkedList(head *ListNode) {
	cur := head
	for ;cur != nil; {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func main() {
	//[1,2,-3,3,1]
	//head := &ListNode{
	//	Val:1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  -3,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: &ListNode{
	//					Val:  1,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	head := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	res := removeZeroSumSublists(head)
	printLinkedList(res)
}
