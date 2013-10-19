package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func handle(err error) bool {
	if err != nil {
		// panic(err)
		fmt.Fprintln(os.Stderr, err)
		return true
	}
	return false
}

func main() {
	host := flag.String("host", "", "select the hostname or IP address on which to serve.")
	port := flag.Int("port", 8080, "select the port number on which to host.")
	auth := flag.String("auth", "", "path to JSON file describing authentication details.")

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	listener, err := net.Listen("tcp", addr)
	if handle(err) {
		return
	}

	for {
		conn, err := listener.Accept()
		if handle(err) {
			continue
		}

		//
	}
}
