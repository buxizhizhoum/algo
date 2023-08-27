package main

import (
	"fmt"
)

//var wg sync.WaitGroup

func Print(ch chan int) {
	for num := range ch {
		//num := <-ch
		fmt.Printf("%d\n", num)
		//fmt.Printf("for loop, num: %d\n", num)
		switch num {
		case 0:
			ch <- 1
		case 1:
			ch <- 0
		default:
			ch <- 2
		}
	}
}

func PrintV2(ch chan int) {
	for num := range ch {
		fmt.Printf("%d\n", num)
		ch <- (num + 1) % 2
	}
}

func PrintV3(chans []chan int, i int) {
	//defer wg.Done()
	for num := range chans[i] {
		fmt.Printf("%d\n", num)
		index := (i + 1) % 10
		chans[index] <- index
	}

}

func main() {
	//ch := make(chan int, 0)

	//go func() {
	//	ch <- 0
	//}()
	//go Print(ch)
	//go Print(ch)
	// 交替打印0, 1
	//go PrintV2(ch)
	//go PrintV2(ch)
	//ch <- 0

	// 多个协程交替打印
	count := 10
	//wg.Add(count)
	chans := make([]chan int, count)
	for i := range chans {
		chans[i] = make(chan int)
	}

	for i := 0; i < count; i++ {
		go PrintV3(chans, i)
	}

	chans[0] <- 0
	//wg.Wait()
	//time.Sleep(10 * time.Millisecond)
}
