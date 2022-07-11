package flatten_nested_list_iterator


// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {return false}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {return 0}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {return nil}


type NestedIterator struct {
	Stack *[]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	// stack := make([]*NestedInteger, 0)
	// for i:= len(nestedList)-1;i>=0;i--{
	//     stack = append(stack, nestedList[i])
	// }
	reverse(&nestedList)
	return &NestedIterator{Stack: &nestedList}
}

func (this *NestedIterator) Next() int {
	element := (*this.Stack)[len(*this.Stack) - 1]
	*this.Stack = (*this.Stack)[:len(*this.Stack) - 1]
	return element.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	// 如果最后一个元素不是整数，弹出来，继续展开最后一个元素，直到展开为整数
	for ;len(*this.Stack) > 0 && !(*this.Stack)[len(*this.Stack) - 1].IsInteger(); {
		cur := (*this.Stack)[len(*this.Stack) - 1]
		*this.Stack = (*this.Stack)[:len(*this.Stack) - 1]
		cur_l := cur.GetList()
		reverse(&cur_l)
		for _, item := range cur_l {
			*this.Stack = append(*this.Stack, item)
		}
	}
	return len(*this.Stack) > 0
}


func reverse(nestedList *[]*NestedInteger) {
	for i,j:=0,len(*nestedList) - 1; i <= j; i,j=i+1, j-1 {
		(*nestedList)[i], (*nestedList)[j] = (*nestedList)[j], (*nestedList)[i]
	}

}
