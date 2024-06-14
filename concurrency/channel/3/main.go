// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.
package main

import (
	"fmt"
	"log"
	"math/rand"
)

// result is what is sent back from each operation.
type result struct {
	id  int
	op  string
	err error
}

func main() {
	// Set the number of routines and inserts.
	routines := 10
	insertions := 2 * routines

	// Number of responses we need to handle.
	waitInserts := insertions

	// Buffered channel to receive information about any possible insert.
	ch := make(chan result, insertions)

	// Perform all the inserts.
	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)
			ch <- insertTrans(id)
		}(i)
	}

	// Process the insert results as they complete.
	for waitInserts > 0 {
		// Wait for a response from a goroutine.
		r := <-ch

		// Display the result.
		log.Printf("N: %d ID: %d OP: %s ERR: %v", waitInserts, r.id, r.op, r.err)
		waitInserts--
	}
	log.Println("Inserts Complete")
}

// insertUser simulates a database operation.
func insertUser(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf(" insert Users value (%d)", id),
	}
	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("unable to insert %d in to Users table", id)
	}

	return r
}

// insertTrans simulates a database operation.
func insertTrans(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf(" insert Trans value (%d)", id),
	}
	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("unable to inser %d in to Trans table", id)
	}

	return r
}
