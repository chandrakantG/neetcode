// go build -race or go run -race main.go

// Sample program to show how to create race conditions in
// our programs. We don't want to do this.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

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
				// Capture the value of Counter.
				value := counter
				// Increment our local value of Counter.
				value++
				// Store the value back into Counter.
				counter = value
				// counter++
			}
			wg.Done()
		}()
	}
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("counter:", counter)
}
