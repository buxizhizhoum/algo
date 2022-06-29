package main

import (
	"fmt"
	"time"
)

import (
	"math/rand"
)

type RandomizedSet struct {
	Val2Index map[int]int
	Nums []int
}


func Constructor() RandomizedSet {
	return RandomizedSet{
		Val2Index: make(map[int]int, 0),
		Nums: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.Val2Index[val]
	if ok {
		return false
	} else {
		this.Nums = append(this.Nums, val)
		// 维护map记录
		this.Val2Index[val] = len(this.Nums) - 1
		return true
	}
}


func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.Val2Index[val]
	if !ok {
		return false
	} else {
		// 最后一个元素对应的索引，改成现在这个值对应的索引，准备后续交换
		this.Val2Index[this.Nums[len(this.Nums) - 1]] = index
		// 交换，删除最后一个元素
		this.Nums[index], this.Nums[len(this.Nums) - 1] = this.Nums[len(this.Nums) - 1], this.Nums[index]
		this.Nums = this.Nums[0:len(this.Nums) - 1]
		// 维护map记录，这一句和上面把nums里面最后一个元素的索引改成index一起，完成了交换时期的索引维护工作
		delete(this.Val2Index, val)
		return true
	}
}


func (this *RandomizedSet) GetRandom() int {
	index := randInt(0, len(this.Nums) - 1)
	return this.Nums[index]
}


func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max - min + 1)
}

//测试用例:["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"] [[],[1],[2],[2],[],[1],[2],[]] 测试结果:[null,true,false,true,1,false,true,1] 期望结果:[null,true,false,true,2,true,false,2]
func main() {
	obj := Constructor()
	//fmt.Println(obj.Insert(1))
	//fmt.Println(obj.Remove(2))
	//fmt.Println(obj.Insert(2))
	//fmt.Println(obj.GetRandom())
	//fmt.Println(obj.Remove(1))
	//fmt.Println(obj.Insert(2))
	//fmt.Println(obj.GetRandom())
	//["RandomizedSet","remove","remove","insert","getRandom","remove","insert"] [[],[0],[0],[0],[],[0],[0]] 测试结果:[null,false,false,true,0,true,false] 期望结果:[null,false,false,true,0,true,true]
	//fmt.Println(obj.Remove(0))
	//fmt.Println(obj.Remove(0))
	//fmt.Println(obj.Insert(0))
	//fmt.Println(obj.GetRandom())
	//fmt.Println(obj.Remove(0))
	//fmt.Println(obj.Insert(0))
	//["RandomizedSet","insert","insert","remove","insert","remove","getRandom"] [[],[0],[1],[0],[2],[1],[]] 测试结果:[null,true,true,true,true,true,1] 期望结果:[null,true,true,true,true,true,2]
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())
}
