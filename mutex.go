package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex
	var value int

	// Writer
	go func() {
		lock.Lock()
		value = 42
		lock.Unlock()
	}()

	// Reader
	lock.Lock()
	fmt.Println("Value:", value)
	lock.Unlock()
}
