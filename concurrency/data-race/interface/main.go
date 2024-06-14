// Sample program to show a more complicated race condition using
// an interface value. This produces a read to an interface value after
// a partial write.
package main

import (
	"fmt"
	"os"
	"sync"
)

// Speaker allows for speaking behavior.
type Speaker interface {
	Speak() bool
}

// Aarav is a person who can speak.
type Aarav struct {
	name string
}

// Speak allows Aarav to say hello. It returns false if the method is
// called through the interface value after a partial write.
func (a *Aarav) Speak() bool {
	if a.name != "Aarav" {
		fmt.Printf("Aarav says, my name is %s \n", a.name)
		return false
	}
	fmt.Println("I am Aarav")
	return true
}

// Shriyansh is a person who can speak.
type Shriyansh struct {
	name string
}

// Speak allows Shriyansh to say hello. It returns false if the method is
// called through the interface value after a partial write.
func (s *Shriyansh) Speak() bool {
	if s.name != "Shriyansh" {
		fmt.Printf("Shriyansh says, my name is %s \n", s.name)
		return false
	}
	fmt.Println("I am Shriyansh")
	return true
}

func main() {
	// Create values of type Aarav and Shriyansh.
	aarav := Aarav{name: "Aarav"}
	shriyansh := Shriyansh{name: "Shriyansh"}

	// create an interface varaible to assign different interface implemented types
	var person Speaker

	// Have a goroutine constantly assign the pointer of
	// the Aarav value to the interface and then Speak.
	go func() {
		for {
			person = &aarav
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	// Have a goroutine constantly assign the pointer of
	// the Shriyansh value to the interface and then Speak.
	go func() {
		for {
			person = &shriyansh
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	// Just hold main from returning. The data race will
	// cause the program to exit.
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
