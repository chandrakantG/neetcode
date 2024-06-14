package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	maxRoutine := 100

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(maxRoutine)

	// Iterate and launch each goroutine.
	for i := 0; i < maxRoutine; i++ {
		// Create an anonymous function for each goroutine.
		go func() {
			// Ensure the waitgroup is decremented when this function returns.
			defer wg.Done()

			// Generate a random number up to 1000.
			n := rand.Intn(1000)
			// Return early if the number is divisible by 2. n%2 == 0
			if n%2 != 0 {
				return
			}
			// Send the odd values through the channel.
			values <- n
		}()
	}

	// Create a goroutine that waits for the other goroutines to finish then
	// closes the channel.
	go func() {
		wg.Wait()
		close(values)
	}()

	// Receive from the channel until it is closed.
	// Store values in a slice of ints.
	var nums []int
	for v := range values {
		nums = append(nums, v)
	}

	// Print the values in our slice.
	fmt.Println("len:", len(nums))
	fmt.Println("nums:", nums)
}
