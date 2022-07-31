package main
// todo: time limit exceeded

import "fmt"

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	if k > n {
		return false
	}
	sum := 0
	for i :=0;i<n;i++ {
		sum += nums[i]
	}
	if sum % k != 0 {
		return false
	}
	target := sum / k
	bucket := make([]int, k)
	return backtrack(nums, k, bucket, 0, target)
}

func backtrack(nums []int, k int, bucket []int, index int, target int) bool {
	if index == len(nums) {
		for _, v := range bucket {
			if v != target {
				return false
			}
		}
		return true
	}

	for j:= 0;j< k;j++ {
		// 剪枝
		if bucket[j] + nums[index] > target {
			continue
		}
		bucket[j] += nums[index]
		flag := backtrack(nums, k, bucket, index + 1, target)
		// 问是否能分成k个子集，不是找到所有的k个子集，所以可以提前返回
		if flag {
			return true
		}
		bucket[j] -= nums[index]
	}
	return false
}


func main() {
	// [4,3,2,3,5,2,1]
	//4
	//canPartitionKSubsets([]int{4,3,2,3,5,2,1}, 4)
	//canPartitionKSubsets([]int{1,1,1,1,2,2,2,2}, 4)
	//res := canPartitionKSubsets([]int{2,2,2,2,3,4,5}, 4)
	res := canPartitionKSubsets([]int{3,3,10,2,6,5,10,6,8,3,2,1,6,10,7,2}, 6)
	fmt.Println(res)
}