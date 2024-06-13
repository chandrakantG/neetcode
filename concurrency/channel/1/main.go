// Sample program to show the order of channel communication for unbuffered,
// buffered and closing channels based on the specification.
package main

import "fmt"

func main() {
	unBuffered()
	buffered()
	closed()
}

// With unbuffered channels, the receive happens before the corresponding send.
// The write to msg happens before the receive on c, which happens before the
// corresponding send on c completes, which happens before the print.
func unBuffered() {
	c := make(chan int)
	var msg string

	go func() {
		msg = "hello, world"
		<-c
	}()

	c <- 0

	// We are guaranteed to print "hello, world".
	fmt.Println(msg)
}

// With buffered channels, the send happens before the corresponding receive.
// The write to msg happens before the send on c, which happens before the
// corresponding receive on c completes, which happens before the print.
func buffered() {
	c := make(chan int, 10)
	var msg string

	go func() {
		msg = "hello, world"
		c <- 0
	}()

	<-c

	// We are guaranteed to print "hello, world".
	fmt.Println(msg)
}

// With both types of channels, a close happens before the corresponding receive.
// The write to msg happens before the close on c, which happens before the
// corresponding receive on c completes, which happens before the print.
func closed() {
	c := make(chan int, 10)
	var msg string

	go func() {
		msg = "hello, world"
		close(c)
	}()

	<-c

	// We are guaranteed to print "hello, world".
	fmt.Println(msg)
}
