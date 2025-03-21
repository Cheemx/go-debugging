package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex
	// fmt.Println(count)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}(i) // debug
	}

	wg.Wait()
	fmt.Println("Final Count:", count)
}

// Debug is adding loop variable in the goroutines running in any loop.
// Why ?
// Loop Variables in Go are re-used across iterations.
// This means that all goroutines capture the same memory location of the loop variable, which can lead to unexpected behaviour.
// Fix
// Adding index or range values explicitly to the goroutines to ensure each goroutine gets its own copy.
