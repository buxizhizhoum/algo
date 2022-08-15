package main

import "fmt"

//例如 2, 6, 3, 5, 4, 1 这个排列， 我们想要找到下一个刚好比他大的排列，于是可以从后往前看
//我们先看后两位 4, 1 能否组成更大的排列，答案是不可以，同理 5, 4, 1也不可以
//直到3, 5, 4, 1这个排列，因为 3 < 5， 我们可以通过重新排列这一段数字，来得到下一个排列
//因为我们需要使得新的排列尽量小，所以我们从后往前找第一个比3更大的数字，发现是4
//然后，我们调换3和4的位置，得到4, 5, 3, 1这个数列 因为我们需要使得新生成的数列尽量小，
//于是我们可以对5, 3, 1进行排序，可以发现在这个算法中，我们得到的末尾数字一定是倒序排列的，
//于是我们只需要把它反转即可 最终，我们得到了4, 1, 3, 5这个数列 完整的数列则是2, 6, 4, 1, 3, 5

//其实就是从数组倒着查找，找到nums[i] 比nums[i+1]小的时候，
//就将nums[i]跟nums[i+1]到nums[nums.length - 1]当中找到一个最小的比nums[i]大的元素交换。
//交换后，再把nums[i+1]到nums[nums.length-1]排序，就ok了
import "sort"
func nextPermutation(nums []int)  {
	n := len(nums)
	if n == 1 {
		return
	}
	for i:= n-2;i>=0;i--{
		if nums[i] < nums[i+1] {
			// 寻找nums[i+1:]里面比nums[i]小的数
			for j:= n-1;j>=i+1;j-- {
				// 从后往前，找到第一个比nums[i]大的数
				if nums[j] > nums[i] {
					// 交换
					nums[i], nums[j] = nums[j], nums[i]
					// 后面部分排序
					tmp := nums[i+1:]
					sort.Slice(tmp, func(i, j int) bool {
						return tmp[i] < tmp[j]
					})
					// 这里排序影响到了吗？
					return
				}
			}
		}
	}
	// 如果到这里还没返回，说明已经是最大的了，上面是降序的，这里直接排序，就是最小的
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

func main() {
	//testNums := []int{1,2,3}
	testNums := []int{3,2,1}
	nextPermutation(testNums)
	fmt.Println(testNums)
}
