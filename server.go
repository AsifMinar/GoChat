package main

import (
	"bufio"
	"log"
	"net"
	"sync"
)

// Shared list of connected clients and a mutex for thread safety
var clients []net.Conn
var mutex sync.Mutex

// Handle a single client: read messages and broadcast them
func handleClient(conn net.Conn) {
	// Add client to the list
	mutex.Lock()
	clients = append(clients, conn)
	mutex.Unlock()

	// Clean up when client disconnects
	defer func() {
		mutex.Lock()
		for i, c := range clients {
			if c == conn {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		mutex.Unlock()
		conn.Close()
		log.Println("Client disconnected")
	}()

	// Read messages from the client
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break // Exit if client disconnects or errors occur
		}
		log.Println("Received: ", message)
		broadcast(message, conn)
	}
}

// Broadcast a message to all clients except the sender
func broadcast(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message))
			if err != nil {
				log.Println("Error sending to client: ", err)
			}
		}
	}
}

func main() {
	// Listen for connections on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
	defer listener.Close()
	log.Println("Server started on :8080")

	// Accept connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}
		// Handle each client in a separate goroutine
		go handleClient(conn)
	}
}