package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func printf(ch chan int) {
	for i:=0;i<10;i++ {
		num := <- ch
		fmt.Println(num)
		ch <- (num + 1) % 2
	}
	wg.Done()
}

func main() {
	ch := make(chan int, 1)
	ch <- 0
	wg.Add(2)
	go printf(ch)
	go printf(ch)
	wg.Wait()
}
