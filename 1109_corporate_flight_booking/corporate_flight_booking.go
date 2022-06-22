package main

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

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	diffArr := Constructor(nums)
	for _, item := range bookings {
		// 这里因为从1开始编号，所以需要对索引减1
		i, j, k := item[0]-1, item[1]-1, item[2]
		diffArr.Increase(i, j, k)
	}
	return diffArr.Result()
}

