package main

import (
    "fmt"
    "net"
    "os"
    "io"
    "bufio"
    "strings"
)

func main() {
    // Check command line arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: receiver <port>")
        return
    }

    // Parse command line argument
    port := os.Args[1]

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
