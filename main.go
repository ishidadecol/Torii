package main

import (
	"fmt"
	"github.com/Torii/cmd/server"
	"github.com/Torii/internals/tunnel"
	"os"
)

func main() {
	manager := tunnel.NewManager()
	fmt.Println("Tunnel Manager started")

	// Start the server
	err := server.StartServer(":9000", manager)
	if err != nil {
		fmt.Println("Server error:", err)
		os.Exit(1)
	}
}

