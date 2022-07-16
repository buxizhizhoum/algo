package main

import "fmt"

type MyQueue struct {
	LeftStack []int
	RightStack []int
}


func Constructor() MyQueue {
	return MyQueue{
		LeftStack: make([]int, 0),
		RightStack: make([]int, 0),
	}
}


func (this *MyQueue) Push(x int)  {
	this.RightStack = append(this.RightStack, x)
}


func (this *MyQueue) Pop() int {
	this.Peek()
	if len(this.LeftStack) > 0{
		value := this.LeftStack[len(this.LeftStack) - 1]
		this.LeftStack = this.LeftStack[:len(this.LeftStack) - 1]
		return value
	}
	return -1
}


func (this *MyQueue) Peek() int {
	if len(this.LeftStack) == 0 {
		for i:=len(this.RightStack) - 1;i>=0;i--{
			value := this.RightStack[i]
			this.LeftStack = append(this.LeftStack, value)
			this.RightStack = this.RightStack[:i]
		}
	}
	if len(this.LeftStack) > 0{
		return this.LeftStack[len(this.LeftStack) - 1]
	} else {
		return -1
	}


}


func (this *MyQueue) Empty() bool {
	return len(this.LeftStack) == 0 && len(this.RightStack) == 0
}


func main() {
	//["MyQueue", "push", "push", "peek", "pop", "empty"]
	//[[], [1], [2], [], [], []]
	//Output
	//[null, null, null, 1, 1, false]
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Peek()
	param_3 := obj.Pop()
	param_3 = obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}