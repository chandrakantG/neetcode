// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutDown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	maxCoroutine := runtime.GOMAXPROCS(0)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(maxCoroutine)

	// Create a fixed size pool of goroutines to generate random numbers.
	for gr := 0; gr < maxCoroutine; gr++ {
		go func(id int) {
			// Start an infinite loop.
			for {
				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {
				// In one case send the random number.
				case values <- n:
					fmt.Printf("worker %d send value %d \n", id, n)

				// In another case receive from the shutdown channel.
				case <-shutDown:
					fmt.Println("process shutDown by ", id)
					wg.Done()
					return
				}
			}
		}(gr)
	}
	// Create a slice to hold the random numbers.
	var nums []int
	for v := range values {
		// continue the loop if the value was even.
		if v%2 == 0 {
			fmt.Println("discarding ", v)
			continue
		}
		// Store the odd number.
		fmt.Println("keeping ", v)
		nums = append(nums, v)

		// break the loop once we have 100 results.
		if len(nums) == 100 {
			// Send the shutdown signal by closing the channel.
			fmt.Println("Receiver sending shutdown signal")
			close(shutDown)
			break
		}
	}

	// Wait for the Goroutines to finish.
	wg.Wait()

	close(values)

	// Print the values in our slice.
	fmt.Println("len:", len(nums))
	fmt.Println("nums:", nums)
}
