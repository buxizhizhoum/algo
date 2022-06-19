package main

import "fmt"




func main() {
	testNums := []int{0,0,1,1,1,1,2,3,3}
	//testNums := []int{1,2,2}
	k := removeDuplicates(testNums)
	fmt.Println(testNums[:k])
}
