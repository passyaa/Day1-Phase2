package main

import (
	"fmt"
	"time"
)

func sendToChannel(evenCh chan<- int, oddCh chan<- int, errorCh chan<- error, done chan bool) {
	for i := 1; i <= 30; i++ { // Intentional error for numbers greater than 20
		if i > 20 {
			errorCh <- fmt.Errorf("number %d is greater than 20", i)
		} else if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	// Close channels when done
	close(evenCh)
	close(oddCh)
	close(errorCh)
	done <- true
}

func readFromChannel(evenCh <-chan int, oddCh <-chan int, errorCh <-chan error, done chan bool) {
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
		case err, ok := <-errorCh:
			if ok {
				fmt.Printf("Error: %s\n", err.Error())
			}
		case <-done:
			return
		}
	}
}

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	errorChannel := make(chan error)
	done := make(chan bool)

	go sendToChannel(evenChannel, oddChannel, errorChannel, done)
	go readFromChannel(evenChannel, oddChannel, errorChannel, done)

	// Allow time for goroutines to complete before exiting main
	time.Sleep(1 * time.Second)
}
