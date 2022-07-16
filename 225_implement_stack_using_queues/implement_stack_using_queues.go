package implement_stack_using_queues

type MyStack struct {
	Queue []int
}


func Constructor() MyStack {
	return MyStack{
		Queue: make([]int, 0),
	}
}

// 每次push，都把栈顶元素挪到队头
func (this *MyStack) Push(x int)  {
	this.Queue = append(this.Queue, x)
	length := len(this.Queue)
	for ;length > 1; length--{
		value := this.Queue[0]
		this.Queue = append(this.Queue, value)
		this.Queue = this.Queue[1:]
	}
}


func (this *MyStack) Pop() int {
	value := this.Queue[0]
	this.Queue = this.Queue[1:]
	return value
}


func (this *MyStack) Top() int {
	return this.Queue[0]
}


func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}
