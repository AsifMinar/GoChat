# GoChat

GoChat is a simple GoLang-based chat application demonstrating basic networking concepts. It consists of a server that manages multiple client connections and broadcasts messages, and a client that connects to the server to send and receive messages. Built for beginners, it uses Go’s `net` package and goroutines to handle concurrent communication. Perfect for learning TCP networking and message handling in Go.

## Features
- Supports multiple clients chatting via a central server
- Handles real-time message broadcasting
- Uses TCP for reliable communication
- Beginner-friendly with clear error handling

## Requirements
- Go 1.x

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/AsifMinar/GoChat.git
   ```
2. Ensure `server.go` and `client.go` are in the project directory.

## Usage
1. **Run the Server**:
   ```bash
   go run server.go
   ```
   The server starts on `localhost:8080`.

2. **Run Clients** (in separate terminals):
   ```bash
   go run client.go
   ```
   Type a message and press Enter to send; received messages appear in the terminal.

Example:
- Start the server: `go run server.go`
- Start two clients in separate terminals: `go run client.go`
- Type "Hello" in one client; it appears as "Received: Hello" in the other.

## Testing
- Run the server in one terminal.
- Run multiple clients in separate terminals.
- Send messages from one client and verify they appear in others.
- Close a client (Ctrl+C) to test disconnection handling.

## Limitations
- Uses `localhost:8080` for simplicity (edit code for remote servers).
- Supports basic text messages without advanced formatting.