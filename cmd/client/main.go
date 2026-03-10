package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//Creates TCP socket
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	//Sends message
	fmt.Fprintf(conn, "Hello")

	//Read response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(buffer[:n]))

	conn.Close()
}
