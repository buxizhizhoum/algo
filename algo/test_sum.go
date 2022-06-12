package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	numDict := make(map[int]int)
	for i:= 0; i < len(nums); i++ {
		item := nums[i]

		restIndex, ok := numDict[target - item]
		if ok {
			res := []int{i, restIndex}
			return res
		}
		numDict[item] = i

	}
	res := []int{-1, -1}
	return res
}

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
//Notice that the solution set must not contain duplicate triplets.
//
//
//
//Example 1:
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Example 2:
//
//Input: nums = []
//Output: []
//Example 3:
//
//Input: nums = [0]
//Output: []


//func threeSum(nums []int) [][]int {
//	res := [][]int{}
//	if len(nums) < 3 {
//		return res
//	}
//
//	existed := make(map[string]int)
//
//	numsDict := make(map[int]int)
//	for i:=0; i < len(nums); i++ {
//		numsDict[nums[i]] = i
//	}
//
//	for i:=0; i < len(nums); i++ {
//		for j:=i+1; j < len(nums); j++ {
//			target := 0 - (nums[i] + nums[j])
//
//			restIndex, ok := numsDict[target]
//			if ok && restIndex != i && restIndex != j {
//				tmp := []int{nums[i], nums[j], target}
//				sort.Ints(tmp)
//
//				key := '"
//				for _, item := range tmp {
//					if len(key) > 0 {
//						key += ","
//					}
//					c := strconv.Itoa(item)
//					key += c
//				}
//
//				_, keyExist := existed[key]
//				if !keyExist {
//					res = append(res, tmp)
//					existed[key] = 1
//				}
//			}
//		}
//	}
//	return res
//}

func threeSum(nums []int) [][]int {
	//res := [][]int{}
	var res [][]int
	if len(nums) < 3{
		return res
	}
	sort.Ints(nums)
	//left, right := 0, len(nums) - 1
	for i:=0; i< len(nums) - 1; i++ {
		item := nums[i]
		rest := 0 - item
		left, right := 0, len(nums) - 1

		if i > 0 && nums[i-1] == nums[i]{
			if i < len(nums) - 2{
				continue
			}
		}
		for left < right{
			if left >= i || right <= i{
				break
			}
			//if nums[left] + nums[right] < rest {
			//	left++
			//} else if nums[left] + nums[right] > rest {
			//	right--
			//} else {
			//	tmp := []int{nums[left], nums[i], nums[right]}
			//	res = append(res, tmp)
			//	left++
			//	right--
			//}
			switch {
			case nums[left] + nums[right] < rest:
				left++
			case nums[left] + nums[right] > rest:
				right--
			default:
				tmp := []int{nums[left], nums[i], nums[right]}
				res = append(res, tmp)
				//for left< right {
				//	if nums[left] == nums[left+1] {
				//		left++
				//	}else if nums[right] == nums[right-1]{
				//		right--
				//	}else{
				//		break
				//	}
				//}

				left++
				right--
			}

		}


	}
	return res
}


// how to define a tree

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}

func _inorderTraversal(root *TreeNode, res *[]int){
	if root == nil {
		return
	}

	_inorderTraversal(root.Left, res)
	//fmt.Println(root.Val)
	*res = append(*res, root.Val)
	_inorderTraversal(root.Right, res)
}



func inorderTraversal(root *TreeNode) []int {
	var res []int
	_inorderTraversal(root, &res)
	return res
}


func inorderTraversalWithStack(root *TreeNode) []int {
	//stack := []*TreeNode{}
	var stack []*TreeNode
	//res := []int{}
	var res []int

	cur := root
	for cur != nil || len(stack) > 0 {
		for cur !=  nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack) - 1]
		stack = stack[: len(stack) - 1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func _minDiffInBST(root *TreeNode, val int) int {
	var minL int
	var minR int
	if root.Left == nil {
		 minL = val
	} else {
		valL := min(val, root.Val - root.Left.Val)
		minL = _minDiffInBST(root.Left, valL)
	}

	if root.Right == nil {
		minR = val
	} else {
		valR := min(val, root.Right.Val - root.Val)
		minR = _minDiffInBST(root.Right, valR)
	}

	return min(minL, minR)
}

func minDiff(l []int) int {
	if len(l) == 0{
		return 0
	}
	if len(l) == 0 {
		return l[0]
	}
	res := 1000000000000
	for i:=0;i < len(l) - 1; i++ {
		res = min(l[i+1] - l[i], res)
	}
	return res
}


func minDiffInBST(root *TreeNode) int {
	l := inorderTraversal(root)
	fmt.Println(l)
	return minDiff(l)
}



func testTwoSum() {
	//testNums := []int{2,7,11,15}
	testNums := []int{3,2,4}
	//balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	target := 6
	fmt.Println(twoSum(testNums, target))
}


func testThreeSum() {
	testNums := []int{-1,0,1,2,-1,-4}
	//testNums := []int{-2,0,1,1,2}
	//testNums := []int{0,0,0,0}
	fmt.Println(threeSum(testNums))
}


func testInorderTraversal() {
	rightRightNode := TreeNode{
		Val: 3,
		Left: nil,
		Right: nil,
	}
	rightNode := TreeNode{
		Val: 2,
		Left: &rightRightNode,
		Right: nil,
	}

	root := TreeNode{
		Val: 1,
		Left: nil,
		Right: &rightNode,
	}
	//res := inorderTraversal(&root)
	//res := inorderTraversalWithStack(&root)
	res := minDiffInBST(&root)
	fmt.Println(res)
}


func testMinDistanceOfBST() {
	root := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: nil,
			Right: nil,
		},
	}
	res := minDiffInBST(&root)
	fmt.Println(res)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{}
	i := 0
	for l1.Next != nil || l2.Next != nil{
		node := ListNode{}
		carry := 0

		if l1.Next != nil && l2.Next != nil {
			val := l1.Val + l2.Val + carry
			curVal := val % 10
			carry = val / 10
			node.Val = curVal
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Next != nil && l2.Next == nil {
			val := l1.Val + carry
			curVal := val % 10
			carry = val / 10
			node.Val = curVal
			l1 = l1.Next
		} else {
			val := l2.Val + carry
			curVal := val % 10
			carry = val / 10
			node.Val = curVal
			l2 = l2.Next
		}
		if i == 0 {
			res = node
		} else {
			node.Next = &node
		}
		res.Next = &node
		i += 1
	}
	return &res
}

func testAddTwoNumbers() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(l3)
}


func bs(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) / 2
		if  target < nums[mid]{
			right = mid - 1
		} else if target > nums[mid]{
			left = mid + 1
		} else {
			return mid
		}
	}
	return  -1
}


func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	//var mid int
	for left < right {
		mid := left + (right - left) / 2

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid]{
			right = mid - 1
		} else {
			return mid
		}

	}

	if nums[left] < target {
		return left + 1
	} else {
		return left
	}
}


func testSearchInsertIndex() {
	//nums := []int{1,2,3,4,5,6,7,8,9}
	//target := 5
	//nums := []int{1,3}
	//target := 2
	nums := []int{1,3,5,6}
	target := 4
	index := searchInsert(nums, target)
	fmt.Println("index: ", index)

}

func testBs() {
	//nums := []int{1,2,3,4,5,6,7,8,9}
	//target := 5
	nums := []int{1,3}
	target := 2

	index := bs(nums, target)
	fmt.Println("index is:", index)
}


func max(i int, j int ) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	cache := make(map[byte]int)
	maxLength := 0
	i, j := 0, 0
	for j < len(s) {
		c := s[j]
		_, ok := cache[c]
		if !ok {
			cache[s[j]] = j
			maxLength = max(maxLength, j - i + 1)
			j += 1
		} else {
			index := cache[c]
			i = max(i, index + 1)
			delete(cache, c)
			//for i < j {
			//	tmp := s[i]
			//	i += 1
			//	delete(cache, tmp)
			//	if tmp == c {
			//		break
			//	}
			//}

		}

	}
	return maxLength

}

func testLengthOfLongestSubstring() {
	length := lengthOfLongestSubstring("abcabcdef")
	//length := lengthOfLongestSubstring("pwwkew")
	fmt.Println("length is: ", length)
}


func findChar(s string, ch byte) int {
	for i:= 0; i < len(s); i ++ {
		if s[i] == ch {
			return i
		}
	}
	return -1
}


func longestSubstring(s string, k int) int {
	counter := make(map[byte]int)
	for i:=0; i< len(s); i++ {
		counter[s[i]] = counter[s[i]] + 1
	}

	//for _, c := range s {
	//	counter[byte(c)] = counter[c] + 1
	//}
	splitIndex := -1
	for key, v := range counter {
		if v < k {
			//find index of key, split it
			splitIndex = findChar(s, key)
			if splitIndex == -1 {
				continue
			} else {
				break
			}
		}
	}

	if splitIndex == -1 {
		return len(s)
	} else {
		maxLeft := longestSubstring(s[:splitIndex], k)
		maxRight := longestSubstring(s[splitIndex + 1:], k)
		return max(maxLeft, maxRight)
	}

}


func testLongestSubstring() {
	l := longestSubstring("aaabb", 3)
	fmt.Println("lenght of longestSubstring is", l)
}


//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	res := climbStairs(n-1) + climbStairs(n-2)
//	fmt.Println("n, res", n, res)
//	return res
//
//}

func climbStairs(n int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	cache[1] = 1
	for i:=2; i <=n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}


func testClimbStairs() {
	res := climbStairs(10)
	fmt.Println("climb stairs: ", res)
}


func countBits(n int) []int {
	cache := make([]int, n+1)
	cache[0] = 0
	for i:=1; i<=n; i++ {
		cache[i] = cache[i & (i - 1)] + 1
	}
	return cache
}

func countBits1(n int) []int {
	cache := make([]int, n+1)
	for i:=0; i <=n; i++ {
		if i % 2 == 0 {
			cache[i] = cache[i/2]
		} else {
			cache[i] = cache[i / 2] + 1
		}
	}
	return cache
}



func testCountBits() {
	res := countBits(15)
	//res := countBits1(nil)
	fmt.Println("count bit: ", res)
}


func minCount(counter map[string]int) int {
	res := 0
	for _, v := range counter {
		res = min(v, res)
	}
	return res
}


//func leastInterval(tasks []string, n int) int {
//	res := 0
//	counter := make(map[string]int)
//	for _, ch := range tasks {
//		_, ok := counter[ch]
//		if !ok {
//			counter[ch] = 1
//		} else {
//			counter[ch] = counter[ch] + 1
//		}
//	}
//
//	length := len(counter)
//	for length > 1 {
//		minNum := minCount(counter)
//
//		if minNum > 0 {
//			if length > 1 {
//				// all count - minNum
//				for k, v := range counter {
//					if v == minNum {
//						// pop if count is 0
//						delete(counter, k)
//					} else {
//						counter[k] = v - minNum
//					}
//				}
//				res += minNum * length
//			}
//		}
//		length = len(counter)
//
//	}
//
//
//	res += (length - 1) * n + length
//
//
//
//	return res
//}

func maxCount(counter map[byte]int) int {
	res := 0
	for _, v := range counter {
		res = max(v, res)
	}
	return res
}


func leastInterval(tasks []byte, n int) int {
	counter := make(map[byte]int)
	res := 0
	for _, ch := range tasks {
		counter[ch]++
	}

	taskMaxNum := maxCount(counter)
	numOfMaxTask := 0
	for _, v := range counter {
		if v == taskMaxNum {
			numOfMaxTask += 1
		}
	}

	res = (taskMaxNum - 1) * (n + 1) + numOfMaxTask
	res = max(len(tasks), res)
	return res

}


func testLeastInterval() {
    tasks := []byte {'A','A','A','A','A','A','B','C','D','E','F','G'}
    res := leastInterval(tasks, 2)
    fmt.Println("least interval: ", res)
}


func checkSubstring(s string, start int, end int) int {
	res := 0
	for start >=0 && end < len(s) && s[start] == s[end] {
		start--
		end++
		res++
	}
	return res

}


func countSubstrings(s string) int {
	res := 0
	for i:= 0; i < len(s); i++ {
		res += checkSubstring(s, i, i)
		res += checkSubstring(s, i, i+1)
	}
	return res



	//res := 0
	//length := len(s)
	////var dp [][]int
	//dp := make([][]int, length)
	//for i:= 0; i < length; i++ {
	//	row := make([]int, length)
	//	dp[i] = row
	//	//dp = append(dp, row)
	//}
	//
	//for i := length - 1; i >= 0; i-- {
	//	for j:=i; j<length; j++ {
	//		if s[i] == s[j] {
	//			if j-i <= 1 {
	//				dp[i][j] = 1
	//				res += 1
	//			} else if dp[i+1][j-1] == 1 {
	//				dp[i][j] = 1
	//				res += 1
	//			}
	//		}
	//	}
	//}
	//return res

}

func testCountSubstring() {
	//s := "aaa"
	s := "abc"
	res := countSubstrings(s)
	fmt.Println("substrings ", res)
}


func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i:= len(temperatures) - 1; i >=0; i-- {
		for stack != nil && temperatures[stack[0]] > temperatures[i] {
			stack = stack[1:]
		}
		if stack == nil {
			stack = append(stack, i)
			res[i] = 0
		} else {
			res[i] = stack[0] - i
			stack = append(stack, i)
		}
	}
	return res

}












func main () {
	//testTwoSum()
	//testThreeSum()
	//testInorderTraversal()
	//testMinDistanceOfBST()
	//testAddTwoNumbers()
	//testBs()
	//testSearchInsertIndex()
	//testLengthOfLongestSubstring()
	//testLongestSubstring()
	//testClimbStairs()
	//testCountBits()
	//testLeastInterval()
	testCountSubstring()


}




