package main

import (
	"github.com/Torii/cmd/server"
	"github.com/Torii/internals/tunnel"
)

func main() {

	//Start Client listener at port 9000
	manager := tunnel.NewManager()
	go server.StartClientServer(":9000", manager)

	//Start Public listener at port 8000
	server.StartPublicServer(":8888", manager)
}
