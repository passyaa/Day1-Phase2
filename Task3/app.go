package main

import (
	"fmt"
	"time"
)

func sendToChannel(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		//fmt.Println("Send:", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func readFromChannel(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)

	go sendToChannel(ch)
	go readFromChannel(ch)

	time.Sleep(1 * time.Second)
}
