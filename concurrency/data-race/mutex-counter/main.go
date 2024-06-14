// go build -race or go run -race main.go

// Sample program to show how to use a mutex to define critical
// sections of code that need synchronous access.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

var mu sync.Mutex

func main() {
	// Number of goroutines to use.
	maxroutine := runtime.GOMAXPROCS(0)

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(maxroutine)

	// Create goroutines.
	for gr := 0; gr < maxroutine; gr++ {
		go func() {
			for i := 0; i < maxroutine; i++ {
				// Only allow one goroutine through this critical section at a time.
				mu.Lock()
				// Capture the value of Counter.
				value := counter
				// Increment our local value of Counter.
				value++
				// Store the value back into Counter.
				counter = value
				// counter++
				// Release the lock and allow any waiting goroutine through.
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("counter:", counter)
}
