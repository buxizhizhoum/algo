package main

import "fmt"

func totalFruit(fruits []int) int {
	left, right := 0, 0
	bucket := make(map[int]int, 0)
	res := 0
	for ;right < len(fruits); {
		// 右侧操作
		bucket[fruits[right]]++
		// 更新窗口内结果
		if len(bucket) <= 2 {
			//fmt.Println(right, left)
			res = max(res, right - left + 1)
		}
		// 遇到超过的情况，抢救一下
		for ;len(bucket) > 2; {
			//fmt.Println("    ", right, left)
			//fmt.Println(bucket)
			leftItem := fruits[left]
			val, _ := bucket[leftItem]
			// 如果多于1个，减少数量，如果只有1个，直接移出
			if val > 1{
				bucket[leftItem]--
			} else if val == 1 {
				delete(bucket, leftItem)
			}
			left++
			// 更新结果
			if len(bucket) <= 2 {
				//fmt.Println(right, left)
				res = max(res, right - left + 1)
			}
		}
		right++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	//testFruits := []int{1,2,3,2,2}
	//[3,3,3,1,2,1,1,2,3,3,4] 测试结果:6 期望结果:5 stdout:
	testFruits := []int{3,3,3,1,2,1,1,2,3,3,4}
	res := totalFruit(testFruits)
	fmt.Println(res)
}
