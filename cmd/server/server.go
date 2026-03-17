package server

import (
	"fmt"
	"io"
	"net"

	"github.com/Torii/internals/tunnel"
)

// Starts Torii server
func StartClientServer(adrr string, manager *tunnel.Manager) error {
	ln, err := net.Listen("tcp", adrr)
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Println("Client Server listening on: ", adrr)

	for {
		clientConn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed accepting connection: ", err)
			continue
		}

		go handleClientConn(clientConn, manager)
	}
}

func StartPublicServer(adrr string, manager *tunnel.Manager) error {
	ln, err := net.Listen("tcp", adrr)
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Println("Public server listening on: ", adrr)

	for {
		userConn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed accepting connection: ", err)
			continue
		}

		go handlePublicConn(userConn, manager)
	}
}

func handleClientConn(clientConn net.Conn, manager *tunnel.Manager) {
	id := tunnel.GenerateID()

	ct := &tunnel.Tunnel{
		ID:   id,
		Conn: clientConn,
	}

	manager.Register(ct)
	fmt.Println("Tunnel created successfully: ", id)

	defer func() {
		manager.Remove(id)
		clientConn.Close()
		fmt.Println("Tunnel created successfully: ", id)
	}()

	select {}
}

func handlePublicConn(userConn net.Conn, manager *tunnel.Manager) {
	t := manager.GetFirst()

	if t == nil {
		fmt.Println("No tunnels available")
		userConn.Close()
		return
	}

	tunnelConn := t.Conn

	go io.Copy(tunnelConn, userConn)
	go io.Copy(userConn, tunnelConn)
}
