package main

import (
	"container/list"
	"fmt"
)


type LRUCache struct {
	Capacity int
	Cache map[int]*list.Element
	List *list.List
}

type KeyPair struct {
	Key int
	Value int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Cache: make(map[int]*list.Element, capacity),
		List: new(list.List),
	}
}


func (this *LRUCache) Get(key int) int {
	if element, ok := this.Cache[key]; ok {
		value := element.Value.(*list.Element).Value.(KeyPair).Value
		this.List.MoveToFront(element)
		return value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if element, ok := this.Cache[key]; ok {
		this.List.MoveToFront(element)
		element.Value.(*list.Element).Value = KeyPair{Key: key, Value: value}
		return
	} else {
		if this.List.Len() == this.Capacity {
			// 获取最后一个节点索引，删除
			lastElement := this.List.Back()
			lastElementKey := lastElement.Value.(*list.Element).Value.(KeyPair).Key
			delete(this.Cache, lastElementKey)
			this.List.Remove(this.List.Back())
		}
	}

	element := &list.Element{
		Value: KeyPair{Key: key, Value: value},
	}
	ele := this.List.PushFront(element)
	this.Cache[key] = ele
	//this.List.PushBack(element)
	//this.List.MoveToFront(element)
}


func main() {
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	//cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(1))
	//cache.Put(3, 3)
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 4)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(4))

	// ["LRUCache","put","put","put","put","get","get"]
	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	//测试结果:[null,null,null,null,null,1,3]
	//期望结果:[null,null,null,null,null,-1,3]
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}