package main

type NumArray struct {
	Nums []int
	Sums []int
}


func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums))
	// 计算和
	if len(nums) > 0 {
		sumArray[0] = nums[0]
	}
	for i:=1; i< len(nums); i++ {
		sumArray[i] = sumArray[i-1] + nums[i]
	}
	return NumArray{
		Nums: nums,
		Sums: sumArray,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.Sums[right] - this.Sums[left] + this.Nums[left]
}

