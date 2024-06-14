// go build -race or go run -race main.go

// Sample program to show how to use the atomic package to
// provide safe access to numeric types.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all goroutines.
var counter int64

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
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("counter:", counter)
}
