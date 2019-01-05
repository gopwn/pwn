package pwn

import (
	"bufio"
	"net"
)

// Dial creates a new network connection using net.Dial
// then creates a pwn.Conn using it and returns it
func Dial(network, addr string) (Conn, error) {
	rawConn, err := net.Dial(network, addr)
	if err != nil {
		return Conn{}, err
	}

	reader := bufio.NewReader(rawConn)
	return Conn{
		c: rawConn,
		r: reader,
	}, nil
}
