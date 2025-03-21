package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Debug!
	}()

	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println(n)
		}
	}()

	wg.Wait()
}

// The first goroutine finishes it's execution but since it didn't close the channel,
// The second goroutine kept on listening as the range of a channel isn't ended until the channel itself is closed.
// So that's why the program runs into deadlock as the main goroutine keeps listening for second goroutine to end and
// The second goroutine is waiting for the channel to close.
