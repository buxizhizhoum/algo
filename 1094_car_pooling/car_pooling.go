package main

import "fmt"

type Difference struct {
	Diff []int
}

func Constructor(nums []int) Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i:=1;i<len(nums);i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return Difference{
		Diff: diff,
	}
}

func (this *Difference) Increase(i, j, k int) {
	this.Diff[i] += k
	if j + 1 < len(this.Diff) {
		this.Diff[j+1] -= k
	}
}

func (this *Difference) Result() []int {
	res := make([]int, len(this.Diff))
	res[0] = this.Diff[0]
	for i:= 1; i<len(this.Diff);i++ {
		res[i] = this.Diff[i] + res[i-1]
	}
	return res
}


func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 10)
	diffArr := Constructor(nums)
	for _, item := range trips {
		// 上下车在同一个站点，这里需要对区间末尾进行调整，比如第9站上车，第9站下车
		// 乘客在车上的区间实际是第9站以前，可以用1-8表示，后面乘客在9可以再上车
		k, i, j := item[0], item[1], item[2]-1
		diffArr.Increase(i, j, k)
	}
	res := diffArr.Result()
	for _, value := range res {
		if value > capacity {
			return false
		}
	}
	return true
}

func main() {
	testTrips := [][]int{
		//{2,1,5},{3,3,7},
		{2,1,5},{3,5,7},
	}
	res := carPooling(testTrips, 3)
	fmt.Println(res)
}
