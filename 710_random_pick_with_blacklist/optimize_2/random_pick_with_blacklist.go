package main


import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	Mapping map[int]int
	Size int
}


func Constructor(n int, blacklist []int) Solution {
	// 把黑名单中的数字，都交换到末尾，但是不是真的交换，而是通过mapping把前面在黑名单的数字映射到末尾段
	mapping := make(map[int]int, 0)
	for i:=0;i<len(blacklist);i++ {
		// blacklist item to blacklist item, it is also the index in nums
		mapping[blacklist[i]] = blacklist[i]
	}
	// 数组分两段[0, n - len(blacklist) - 1], 闭区间, [n - len(blacklist), n - 1]
	// 对黑名单中的元素进行检查，然后如果在前段存在黑名单元素，将其映射到数组后端
	// 但是这里交换并没有实际交换元素位置，只是在mapping中进行了虚拟的映射
	last := n - 1
	for _, val := range blacklist {
		// 这里的val在[0, n)之间
		if val > n - len(blacklist) - 1 {
			continue
		}
		// todo: 注意这里的for循环写法
		for _, rok := mapping[last];rok ; _, rok = mapping[last]{
			// 找到末端元素不属于黑名单中的例子
			//_, rok := mapping[last]
			//if rok {
			//	last--
			//	continue
			//}
			//break
			last--
		}
		// 前半段的黑名单元素，映射到后段，val是[0, n)之间的一个数，代表了黑名单元素在前半段的索引
		mapping[val] = last
		last--
	}

	return Solution{
		Mapping: mapping,
		Size: n - len(blacklist) - 1,
	}
}


func (this *Solution) Pick() int {
	index := randInt(0, this.Size)
	fmt.Printf("index: %d\n", index)
	tmp, ok := this.Mapping[index]
	if ok {
		return tmp
	}
	return index
}


func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max - min + 1)
}

func main() {
	//["Solution","pick","pick","pick","pick","pick","pick","pick"]
	//[[7,[2,3,5]],[],[],[],[],[],[],[]]
	//n:=5
	//blacklist := []int{0,1,4}
	//n:=7
	//blacklist := []int{2,3,5}
	n:=4
	//blacklist := []int{2,1}
	blacklist := []int{0,3}
	//[1000000000,[640145908]
	//n := 1000000000
	//blacklist := []int{640145908}
	obj := Constructor(n, blacklist)
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
}
