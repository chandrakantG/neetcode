// Sample program to show how to use a read/write mutex to define critical
// sections of code that needs synchronous access.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Number of reads occurring at ay given time.
var readCount int64

// rwMutex is used to define a critical section of code.
var rwMux sync.RWMutex

// data is a slice that will be shared.
var data []string

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a writer goroutine.
	go func() {
		for i := 0; i < 10; i++ {
			write(i)
		}
		wg.Done()
	}()

	// Create reader goroutines.
	for i := 0; i < 8; i++ {
		go func(id int) {
			for {
				read(id)
			}
		}(i)
	}

	// Wait for the write goroutine to finish.
	wg.Wait()

	fmt.Println("rc:", atomic.LoadInt64(&readCount))
	fmt.Println("Program Complete")
}

// writer adds a new string to the slice.
func write(id int) {
	// Only allow one goroutine to read/write to the slice at a time.
	rwMux.Lock()
	// Capture the current read count.
	// Keep this safe though we can due without this call.
	rc := atomic.LoadInt64(&readCount)

	fmt.Printf("writing operation for %d : RCount[%d] \n", id, rc)
	data = append(data, "data")

	// Release the lock.
	rwMux.Unlock()
}

// reader wakes up and iterates over the data slice.
func read(id int) {
	// Any goroutine can read when no write operation is taking place.
	rwMux.RLock()

	// Increment the read count value by 1.
	rc := atomic.AddInt64(&readCount, 1)
	fmt.Printf("reading operation by %d : length[%d]:RCount[%d] \n", id, len(data), rc)

	// Decrement the read count value by 1.
	atomic.AddInt64(&readCount, -1)
	rwMux.RUnlock()
}
