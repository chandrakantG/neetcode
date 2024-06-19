// Sample program demonstrating that type assertions are a runtime and
// not compile time construct.
package main

import (
	"fmt"
	"math/rand"
)

// car represents something you drive.
type car struct{}

// String implements the fmt.Stringer interface.
func (c car) String() string {
	return "vroom..."
}

// cloud represents somewhere you store information.
type cloud struct{}

// String implements the fmt.Stringer interface.
func (c cloud) String() string {
	return "Big Data..."
}

func main() {
	// Create a slice of the Stringer interface values.
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}
	for i := 0; i < 10; i++ {
		// Choose a random number from 0 to 1.
		rn := rand.Intn(2)

		// Perform a type assertion that we have a concrete type
		// of cloud in the interface value we randomly chose.
		if v, ok := mvs[rn].(cloud); ok {
			fmt.Println("got lucky:", v)
			continue
		}
		fmt.Println("got unlucky")
	}
}
