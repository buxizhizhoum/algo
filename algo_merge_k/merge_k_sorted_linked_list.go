package algo_merge_k

import "container/heap"

type ListNode struct {
     Val int
     Next *ListNode
 }


type ListNodeHeap []ListNode

func(h ListNodeHeap) Len() int {return len(h)}

func(h ListNodeHeap) Less(i, j int) bool {return h[i].Val < h[j].Val}

func(h ListNodeHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func(h *ListNodeHeap) Push(value interface{}) {*h = append(*h, value.(ListNode))}

func(h *ListNodeHeap) Pop() interface{} {
	n := h.Len()
	value := (*h)[n-1]
	*h = (*h)[0:n-1]
	return value
}




func mergeKLists(lists []*ListNode) *ListNode {
	nodeHeap := &ListNodeHeap{}
	heap.Init(nodeHeap)
	dummyNode := ListNode{
		Val:-1,
		Next:nil,
	}
	cur := &dummyNode
	for _, l := range lists {
		if l != nil {
			heap.Push(nodeHeap, *l)
		}
	}
	for ; nodeHeap.Len() > 0; {
		n := heap.Pop(nodeHeap)
		node := n.(ListNode)
		cur.Next = &node
		if &node != nil && node.Next != nil {
			heap.Push(nodeHeap, *(node.Next))
		}
		cur = cur.Next
	}
	return dummyNode.Next
}