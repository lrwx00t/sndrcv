package main

import (
	"flag"
	"fmt"
	"os"

	"sndrcv/lib"
	_ "sndrcv/lib"
)

func main() {
	// Define flags
	mode := flag.String("mode", "", "Mode of operation: send or receive")
	flag.Parse()

	// Validate `-mode` flag
	if *mode == "" {
		fmt.Println("Error: mode flag is required (send or receive)")
		fmt.Println("Usage: ./sndrcv -mode=send <filename> <destination_ip:port>")
		fmt.Println("       ./sndrcv -mode=receive")
		os.Exit(1)
	}

	switch *mode {
	case "send":
		if len(os.Args) != 4 {
			fmt.Println("Usage: ./sndrcv send <filename> <destination_ip:port>")
			return
		}
		filename := os.Args[2]
		destination := os.Args[3]
		lib.Send(filename, destination)
	case "receive":
		lib.Rcv(os.Args[2])
	default:
		fmt.Printf("Error: invalid mode '%s'. Use 'send' or 'receive'\n", *mode)
		os.Exit(1)
	}
}
