// Sample program to show how to use an unbuffered channel to
// simulate a relay race between four goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {
	// Create an unbuffered channel.
	track := make(chan int)

	// Add a count of one for the last runner.
	wg.Add(1)

	// First runner to his mark.
	go runner(track)

	// Start the race.
	track <- 1

	// Wait for the race to finish.
	wg.Wait()
}

// Runner simulates a person running in the relay race.
func runner(track chan int) {

	// The number of exchanges of the baton.
	const maxExchange = 4
	var exchange int

	// Wait to receive the baton.
	baton := <-track

	// Start running around the track.
	fmt.Printf("player %d started race \n", baton)

	// New runner to the line.
	if baton < maxExchange {
		exchange = baton + 1
		fmt.Printf("player %d ready on line \n", exchange)
		go runner(track)
	}

	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if baton == maxExchange {
		fmt.Printf("player %d finished, race over \n", baton)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("player %d exchange baton with player %d  \n", baton, exchange)
	track <- exchange
}
