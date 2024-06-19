// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/chandrakantG/neetcode/concurrency/patteens/task/task"
)

// names provides a set of names to display.
var names = []string{
	"ajay",
	"vijay",
	"sujay",
	"jay",
	"jay-vijay",
}

// namePrinter provides special support for printing names.
type NamePrinter struct {
	name string
}

// Work implements the Worker interface.
func (n NamePrinter) Work() {
	fmt.Println("Name:", n.name)
	time.Sleep(3 * time.Second)
}

func main() {
	maxRoutine := 10

	// Create a task pool.
	t := task.New(maxRoutine)
	var wg sync.WaitGroup
	wg.Add(maxRoutine * len(names))

	for range maxRoutine {
		// Iterate over the slice of names.
		for _, name := range names {
			go func(n string) {
				// Create a namePrinter
				np := NamePrinter{name: n}
				// Submit the task to be worked on. When Do
				// returns, we know it is being handled.
				t.Do(np)
				wg.Done()
			}(name)
		}
	}
	// wait for goroutine to finish their task
	wg.Wait()
	// Shutdown the task pool and wait for all existing work
	// to be completed.
	t.Shutdown()
}
