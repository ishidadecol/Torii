package tunnel

import "net"

type Tunnel struct {
	ID   string
	Conn net.Conn
}
