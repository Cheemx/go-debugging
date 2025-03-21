package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * 2)
		}(num)
		wg.Wait() // 1st debug
	}
	// wg.Wait()
}

// Adding wg.Wait() into the loop doesn't actually debugs the issue here it just
// serializes the execution.The actual issue was passing the correct num each
// time to the goroutine as goroutine is executed without stopping the execution
// parent program so when it catches the value of num it'll be probably be last
// element from array.

// Actual debug pass num as parameter to goroutine and work on it inside the goroutine!
