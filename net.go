package pwn

import (
	"net"
)

// Dial creates a new network connection using net.Dial
// then creates a pwn.Conn using it and returns it
func Dial(network, addr string) (Conn, error) {
	rawConn, err := net.Dial(network, addr)
	if err != nil {
		return conn{}, err
	}

	return conn{
		c: rawConn,
	}, nil
}
