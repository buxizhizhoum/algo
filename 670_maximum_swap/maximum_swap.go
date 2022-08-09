package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximumSwap(num int) int {
	// 比他大的里面最大的进行交换 转换为 知道数字索引，从大到小来，如果大于则交换
	numStr := strconv.Itoa(num)
	nums := []byte(numStr)
	if len(nums) == 1 {
		return num
	}
	// 数字最后一次在num中出现的索引
	lastIndex := make(map[byte]int, 0)
	for i, ch := range nums {
		lastIndex[ch] = i
	}
	swap(nums, lastIndex)
	// convert to int
	numStrSlice := make([]string, 0)
	for _, ch := range nums {
		numStrSlice = append(numStrSlice, string(ch))
	}
	numStr = strings.Join(numStrSlice, "")
	res, _ := strconv.Atoi(numStr)
	return res
}

func swap(nums []byte, lastIndex map[byte]int) {
	for i:= 0;i<len(nums);i++ {
		for j:= byte('9'); j > nums[i];j-- {
			if index, ok := lastIndex[j]; ok {
				// 最后一次出现，一定要是在后面出现，表示i的右半段，这样交换数字才能变大
				if index > i {
					nums[i], nums[index] = nums[index], nums[i]
					return
				}
			}
		}
	}
}

func main() {
	//res := maximumSwap(1234)
	//res := maximumSwap(2736)
	res := maximumSwap(9973)
	//res := maximumSwap(10)
	fmt.Println(res)
}
