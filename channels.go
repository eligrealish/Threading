package main

import "fmt"

func main() {
	ch := make(chan int)

	// Sender
	go func() {
		ch <- 42
	}()

	// Receiver
	value := <-ch
	fmt.Println("Received:", value)
}
