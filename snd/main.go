package main

import (
    "fmt"
    "net"
    "os"
    "io"
    "path/filepath"
)

func main() {
    // Check command line arguments
    if len(os.Args) != 3 {
        fmt.Println("Usage: sender <filename> <destination_ip:port>")
        return
    }

    // Parse command line arguments
    filename := os.Args[1]
    destination := os.Args[2]

    // Open the file for reading
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Connect to the destination machine
    conn, err := net.Dial("tcp", destination)
    if err != nil {
        fmt.Println("Error connecting to destination:", err)
        return
    }
    defer conn.Close()

    // Extract the filename from the path
    _, filenameToSend := filepath.Split(filename)

    // Send the filename followed by a newline character as a delimiter
    _, err = conn.Write([]byte(filenameToSend + "\n"))
    if err != nil {
        fmt.Println("Error sending filename:", err)
        return
    }

    // Transfer file contents
    _, err = io.Copy(conn, file)
    if err != nil {
        fmt.Println("Error sending file:", err)
        return
    }

    fmt.Println("File sent successfully!")
}
