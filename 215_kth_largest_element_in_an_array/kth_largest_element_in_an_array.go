package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func(h IntHeap) Len() int {return len(h)}

func(h IntHeap) Less(i, j int) bool {return h[i] < h[j]}

func(h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func(h *IntHeap) Push(value interface{}) {*h = append(*h, value.(int))}

func(h *IntHeap) Pop() interface{} {
	n := h.Len()
	value := (*h)[n-1]
	*h = (*h)[0:n-1]
	return value
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func findKthLargest(nums []int, k int) int {
	tmpHeap := &IntHeap{}
	heap.Init(tmpHeap)

	for _, v := range nums {
		heap.Push(tmpHeap, v)
		if tmpHeap.Len() > k {
			heap.Pop(tmpHeap)
		}
	}
	res := (*tmpHeap)[0]
	return res
	//return res.(int)
}


func main() {
	nums := []int{3,2,1,5,6,4}
	k := 2
	res := findKthLargest(nums, k)
	fmt.Println(res)
}


