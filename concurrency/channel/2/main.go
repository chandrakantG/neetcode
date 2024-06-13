// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	// stop := make(chan struct{})
	// channelCancellation(stop)
	// cancellation()
	// drop()
	// boundedWorkPooling()
	// fanOutSem()
	// pooling()
	// waitForTask()
	// fanOut()
	waitForResult()
}

// channelCancellation shows how you can take an existing channel being
// used for cancellation and convert that into using a context where
// a context is needed.
func channelCancellation(stop <-chan struct{}) {

	// Create a cancel context for handling the stop signal.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// If a signal is received on the stop channel, cancel the
	// context. This will propagate the cancel into the p.Run
	// function below.
	go func() {
		select {
		case <-stop:
			fmt.Println("call cancel func")
			cancel()
		case <-ctx.Done():
			fmt.Println("ctx done")
		}
	}()

	// Imagine a function that is performing an I/O operation that is
	// cancellable.
	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.ardanlabs.com/blog/index.xml", nil)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		fmt.Println("resp:", resp)
		return nil
	}(ctx)
}

// cancellation: In this pattern, the parent goroutine creates a child
// goroutine to perform some work. The parent goroutine is only willing to
// wait 150 milliseconds for that work to be completed. After 150 milliseconds
// the parent goroutine walks away.
func cancellation() {
	duration := 150 * time.Millisecond
	ch := make(chan string, 1)
	ctx, cancle := context.WithTimeout(context.Background(), duration)
	defer cancle()

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

// drop: In this pattern, the parent goroutine signals 20 pieces of work to
// a single child goroutine that can't handle all the work. If the parent
// performs a send and the child is not ready, that work is discarded and dropped.
func drop() {
	ch := make(chan string, 5)
	const work = 20

	go func() {
		for wrk := range ch {
			fmt.Println("child: received signal ", wrk)
		}
	}()

	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent: sent signal ", w)
		default:
			fmt.Println("parent: drop signal ", w)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// boundedWorkPooling: In this pattern, a pool of child goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signalling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the child
// goroutines terminate.
func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", 10: "paper"}
	fmt.Printf("%+v,%v \n", work, len(work))
	ch := make(chan string)
	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, w := range work {
		ch <- w
	}
	close(ch)
	wg.Wait()
}

// A semaphore is a variable used to control access to a shared resource within the operating system,
// and a mutex is simply a lock acquired before entering a critical section and releasing it.
// A semaphore is better for multiple instances of a resource, but a mutex is better for a single shared resource

// fanOutSem: In this pattern, a semaphore is added to the fan out pattern
// to restrict the number of child goroutines that can be schedule to run.
func fanOutSem() {
	ch := make(chan string)
	children := 20
	g := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS: ", g)
	sem := make(chan bool, g)

	for i := 0; i < children; i++ {
		go func(child int) {
			sem <- true
			ch <- "data"
			fmt.Printf("child: %v sent signal \n ", child)
			<-sem
		}(i)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Printf("parent: received signal %v for %v \n", d, children)
	}
	// time.Sleep(1 * time.Second)
	fmt.Println("-------------------------------------------------")
}

// pooling: In this pattern, the parent goroutine signals 100 pieces of work
// to a pool of child goroutines waiting for work to perform.
func pooling() {
	ch := make(chan string)
	const work = 100
	g := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS: ", g)
	for c := 0; c < g; c++ {
		go func(child int) {
			for v := range ch {
				fmt.Printf("child: %d : received signal %s \n", child, v)
			}
			fmt.Printf("child: %d : received shutdown signal \n", child)
		}(c)
	}

	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent: sent signal ", w)
	}
	close(ch)
	fmt.Println("parent: sent shutdown signal")
	fmt.Println("-------------------------------------------------")
}

// waitForTask: In this pattern, the parent goroutine sends a signal to a
// child goroutine waiting to be told what to do.
func waitForTask() {
	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("child: received signal ->", d)
	}()

	ch <- "data"
	fmt.Println("parent: sent signal")
	fmt.Println("-------------------------------------------------")
}

// fanOut: In this pattern, the parent goroutine creates 2000 child goroutines
// and waits for them to signal their results.
func fanOut() {
	ch := make(chan string)
	children := 2000

	for i := 0; i < children; i++ {
		go func(child int) {
			ch <- "data"
			fmt.Printf("child: %v sent signal \n ", child)
		}(i)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Printf("parent: received signal %v for %v \n", d, children)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("-------------------------------------------------")
}

// waitForResult: In this pattern, the parent goroutine waits for the child
// goroutine to finish some work to signal the result.
func waitForResult() {
	ch := make(chan string)

	go func() {
		ch <- "data"
		fmt.Println("child: sent signal")
	}()

	d := <-ch

	fmt.Println("parent: received signal-> ", d)
	fmt.Println("-------------------------------------------------")
}
