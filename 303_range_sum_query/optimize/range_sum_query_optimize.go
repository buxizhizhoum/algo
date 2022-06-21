package optimize

type NumArray struct {
	Sums []int
}


func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums) + 1)
	// 计算和，使用技巧，多分配一个元素，第一个元素默认给0
	sumArray[0] = 0
	for i:=1; i < len(nums) + 1; i++ {
		sumArray[i] = sumArray[i-1] + nums[i-1]
	}
	return NumArray{
		Sums: sumArray,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.Sums[right + 1] - this.Sums[left]
}
