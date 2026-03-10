package server

import (
	"fmt"
	"io"
	"net"

	"github.com/Torii/internals/tunnel"
)

func StartServer(adrr string, manager *tunnel.Manager) error {
	ln, err := net.Listen("tcp", adrr)
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Println("Server listening on: ", adrr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed accepting connection: ", err)
			continue
		}

		go handleConnection(conn, manager)

	}
}

func handleConnection(conn net.Conn, manager *tunnel.Manager) {
	id := tunnel.GenerateID()

	t := &tunnel.Tunnel{
		ID:   id,
		Conn: conn,
	}

	manager.Register(t)
	fmt.Println("Tunnel created successfully: ", id)

	defer func() {
		manager.Remove(id)
		conn.Close()
		fmt.Println("Tunnel closed successfully: ", id)
	}()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Connection error: ", err)
			}
			break
		}
		conn.Write(buf[:n])
	}
}
