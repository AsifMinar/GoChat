package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Failed to connect to server: ", err)
	}
	defer conn.Close()
	log.Println("Connected to server at localhost:8080")

	// Goroutine to receive messages from the server
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Server closed connection")
				os.Exit(0)
			}
			fmt.Print("Received: ", message)
		}
	}()

	// Main loop to send messages to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}