package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func getLocalIP() (string, error) {
	// Create a UDP connection to a remote address (doesn't need to be reachable)
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Extract the local address from the connection
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func main() {
	// Check command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: receiver <port>")
		return
	}

	// Parse command line argument
	port := os.Args[1]
	local_ip, _ := getLocalIP()
	fmt.Printf("Connect to: %s:%s\n", local_ip, port)

	// Create a TCP listener
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Waiting for incoming connections...")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue // Continue listening for new connections
		}

		go handleConnection(conn) // Handle the connection in a goroutine
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Receive the filename followed by a newline character
	reader := bufio.NewReader(conn)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving filename:", err)
		return
	}

	// Trim the newline character from the filename
	filename = strings.TrimSpace(filename)

	// Create a file for writing with the received filename
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Read the remaining data (file contents) from the connection
	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return
	}

	fmt.Printf("File '%s' received successfully!\n", filename)
}
