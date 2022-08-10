package utils

import "fmt"

func PrintGrid(grid [][]int) {
	for _, tmp := range grid {
		fmt.Println(tmp)
	}
	fmt.Println()
}
