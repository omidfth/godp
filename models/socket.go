package models

import "net"

type Socket struct {
	SocketID string
	Address  net.Addr
}
