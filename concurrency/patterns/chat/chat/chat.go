// Package chat implements a basic chat room.
package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

// message is the data received and sent to users in the chatroom.
type message struct {
	data string
	conn net.Conn
}

// client represents a single connection in the room.
type client struct {
	name   string
	room   *Room
	reader *bufio.Reader
	writer *bufio.Writer
	conn   net.Conn
	wg     sync.WaitGroup
}

// Room contains a set of networked client connections.
type Room struct {
	listener net.Listener
	clients  []*client
	joining  chan net.Conn
	outgoing chan message
	shutdown chan struct{}
	wg       sync.WaitGroup
}

// read waits for message and sends it to the chatroom for processing.
func (c *client) read() {
	for {
		// Wait for a message to arrive.
		line, err := c.reader.ReadString('\n')
		if err == nil {
			c.room.outgoing <- message{
				data: line,
				conn: c.conn,
			}
			continue
		}
		fmt.Printf("client read error: %v \n", err)
		c.wg.Done()
		return
	}
}

// drop closes the client connection and read goroutine.
func (c *client) drop() {
	// Close the connection.
	c.conn.Close()
	c.wg.Wait()
}

// newClient create a new client for an incoming connection.
func newClient(conn net.Conn, rm *Room, name string) *client {
	c := client{
		name:   name,
		room:   rm,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		conn:   conn,
	}
	c.wg.Add(1)
	go c.read()
	return &c
}

// join takes a new connection and adds it to the room.
func (rm *Room) join(conn net.Conn) {
	name := fmt.Sprintf("conn-%d", len(rm.clients))
	fmt.Println("New conn joing room:", name)
	c := newClient(conn, rm, name)
	rm.clients = append(rm.clients, c)
}

// write is a goroutine to handle processing outgoing
// messages to this client.
func (c *client) write(msg message) {
	m := fmt.Sprintf("%v : %s", c.name, msg.data)
	log.Println(m)
	c.writer.WriteString(m)
	c.writer.Flush()
}

// sendGroupMessage sends a message to all clients in the room.
func (rm *Room) sendGroupMessage(msg message) {
	for _, c := range rm.clients {
		if c.conn != msg.conn {
			c.write(msg)
		}
	}
}

// start turns the chatroom on.
func (rm *Room) start() {
	rm.wg.Add(1)

	// Chatroom processing goroutne.
	go func() {
		for {
			select {
			case message := <-rm.outgoing:
				// Sent message to the group.
				rm.sendGroupMessage(message)
			case conn := <-rm.joining:
				// Join this connection to the room.
				rm.join(conn)
			case <-rm.shutdown:
				// Chatroom shutting down.
				rm.wg.Done()
				return
			}
		}
	}()

	// Chatroom connection accept goroutine.
	go func() {
		var err error
		if rm.listener, err = net.Listen("tcp", ":6000"); err != nil {
			log.Fatalln(err)
		}
		log.Println("Chat room started: 6000")
		for {
			conn, err := rm.listener.Accept()
			if err != nil {
				log.Println("accept-routine", err)
				continue
			}
			// Add this new connection to the room.
			rm.joining <- conn
		}
	}()
}

// Close shutdown the chatroom and closes all connections.
func (rm *Room) Close() error {

	// Don't accept anymore client connections.
	rm.listener.Close()

	// Signal the chatroom processing goroutine to stop.
	close(rm.shutdown)
	rm.wg.Wait()

	// Drop all existing connections.
	for _, c := range rm.clients {
		c.drop()
	}
	return nil
}

// New creates a new chatroom.
func New() *Room {
	// Create a Room value.
	rm := &Room{
		joining:  make(chan net.Conn),
		outgoing: make(chan message),
		shutdown: make(chan struct{}),
	}
	// Start the chatroom.
	rm.start()

	return rm
}
