package pwn

import (
	"net"
)

// Dial creates a new network connection using net.Dial
// then creates a pwn.Conn using it and returns it
// MaxLineLength is by default set to 256, you can change it in the returned
// Conn using Conn.MaxLen(i int).
func Dial(network, addr string) (Conn, error) {
	rawConn, err := net.Dial(network, addr)
	if err != nil {
		return &conn{}, err
	}

	return &conn{
		c: rawConn,
		// the default line length to be used with conn.ReadLine
		maxLen: 256,
	}, nil
}
