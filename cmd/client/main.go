package main

import (
	"io"
	"log"
	"net"
)

func main() {

	serverConn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	localConn, err := net.Dial("tcp", "localhost:4321")
	if err != nil {
		log.Fatal(err)
	}

	go io.Copy(localConn, serverConn)
	io.Copy(serverConn, localConn)
}
