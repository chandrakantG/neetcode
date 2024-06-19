// Package task provides a pool of goroutines to perform tasks.
package task

import "sync"

// Worker must be implemented by types that want to use
// the run pool.
type Worker interface {
	Work()
}

// Task provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Task struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New creates a new work pool.
func New(maxRoutine int) *Task {
	t := Task{
		// Using an unbuffered channel because we want the
		// guarantee of knowing the work being submitted is
		// actually being worked on after the call to Run returns.
		work: make(chan Worker),
	}

	t.wg.Add(maxRoutine)
	for range maxRoutine {
		go func() {
			for w := range t.work {
				w.Work()
			}
			t.wg.Done()
		}()
	}
	return &t
}

// Shutdown waits for all the goroutines to shutdown.
func (t *Task) Shutdown() {
	close(t.work)
	t.wg.Wait()
}

// Do submits work to the pool.
func (t *Task) Do(w Worker) {
	t.work <- w
}
