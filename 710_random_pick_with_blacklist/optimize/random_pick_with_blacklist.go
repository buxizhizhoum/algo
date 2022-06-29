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

// time limit exceed
func Constructor(n int, blacklist []int) Solution {
	// 把黑名单中的数字，都交换到末尾
	//nums := make([]int, n)
	//for i:=0;i<n;i++{
	//	nums[i] = i
	//}
	mapping := make(map[int]int, 0)
	for i:=0;i<len(blacklist);i++ {
		mapping[blacklist[i]] = blacklist[i]
	}
	// n个元素，总共有len(blacklist)个黑名单，那么对索引从0到n - len(blacklist) + 1的元素进行检查
	// 如果是黑名单元素，则和末尾元素交换，否则略过
	// 但是这里交换并没有实际交换元素位置，只是在mapping中进行了虚拟的映射
	right := n - 1
	for i := 0; i <= n - len(blacklist) - 1; i++ {
		_, lok := mapping[i]
		if lok {
			// 寻找右侧可以交换的非黑名单元素
			for j := right; j > n - len(blacklist) - 1; j--{
				_, rok := mapping[j]
				if !rok {
					// 交换
					mapping[i] = j
					right=j-1
					break
				}
			}
		}
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
	//n:=7
	//n:=5
	//blacklist := []int{2,3,5}
	//blacklist := []int{0,1,4}
	//[1000000000,[640145908]
	n := 1000000000
	blacklist := []int{640145908}
	obj := Constructor(n, blacklist)
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
}
