package main

import (
	"fmt"
	"time"
)

func addNumberToChan(ch chan int) {
	for {
		ch <- 1
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go addNumberToChan(ch1)
	go addNumberToChan(ch2)

	for {
		select {
		case e1 := <- ch1:
			fmt.Println("get element from ch1: ", e1)
		case e2 := <-ch2:
			fmt.Println("get element from ch2: ", e2)
		//default:
		//	fmt.Println("no element")
		//	time.Sleep(time.Millisecond * 1000)
		}
	}
}
