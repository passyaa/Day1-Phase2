package main

import (
	"fmt"
	"time"
)

func sendToChannel(evenCh chan<- int, oddCh chan<- int, done chan bool) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	// Close channels when done
	close(evenCh)
	close(oddCh)
	done <- true
}

func readFromChannel(evenCh <-chan int, oddCh <-chan int, done chan bool) {
	for {
		select {
		case even, ok := <-evenCh:
			if ok {
				fmt.Printf("Received Even Number : %d\n", even)
			}
		case odd, ok := <-oddCh:
			if ok {
				fmt.Printf("Received Odd Number : %d\n", odd)
			}
		case <-done:
			return
		}
	}
}

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	done := make(chan bool)

	go sendToChannel(evenChannel, oddChannel, done)
	go readFromChannel(evenChannel, oddChannel, done)

	// Allow time for goroutines to complete before exiting main
	time.Sleep(1 * time.Second)
}
