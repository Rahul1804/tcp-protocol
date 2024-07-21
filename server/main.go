package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"tcp-protocol/protocol"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		message = strings.TrimSpace(message)
		fmt.Printf("Received: %s\n", message)

		response := protocol.ProcessMessage(message)
		_, err = conn.Write([]byte(response + "\n"))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func main() {
	fmt.Println("Starting server...")

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(conn)
	}
}
