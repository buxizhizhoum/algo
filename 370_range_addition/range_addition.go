// 区间加法，假设你有一个数组，长度n, 初始情况下，所有元素为0，现在需要进行k次更新操作，
// 每个操作会被描述成一个三元组，[i, j, k]，表示对从i到j的元素增加k，两侧都为闭区间
// 返回k次操作后的结果
// 差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减

package main

import "fmt"

type Difference struct {
	Diff []int
}

// 构建差分数组
func Constructor(nums []int) Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i:=1; i< len(nums); i++ {
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
	for i:=1; i< len(this.Diff);i++ {
		res[i] = res[i-1] + this.Diff[i]
	}
	return res
}

func getModifiedArray(length int, updates [][]int) []int {
	nums := make([]int, length)
	diffArr := Constructor(nums)
	for _, item := range updates {
		i, j, k := item[0], item[1], item[2]
		diffArr.Increase(i, j, k)
	}
	return diffArr.Result()
}

func main() {
	//testUpdates := [][]int{
	//	{1,3,2},
	//	{2,4,3},
	//	{0,2,-2},
	//}
	testUpdates :=[][]int{
		{1,2,10},
		{2,3,20},
		{2,5,25},
	}
	//[[1,2,10],[2,3,20],[2,5,25]] 5 测试结果:[0,10,55,45,25] 期望结果:[10,55,45,25,25] stdout
	tmp := getModifiedArray(5, testUpdates)
	fmt.Println(tmp)

}