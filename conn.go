package pwn

import (
	"errors"
	"net"
	"sync"
)

// ErrMaxLen indecates that the max length was reached for ReadTill
var ErrMaxLen = errors.New("max length reached")

// Conn is a generic stream-oriented network connection.
//
// Multiple goroutines may invoke methods on a Conn simultaneously.
type Conn struct {
	// the embeedded net.Conn
	net.Conn

	// the max length for ReadLine and ReadTill.
	maxLen int

	// mu is used for protecting struct variables from concurrent reads / writes
	mu sync.Mutex
}

func (c *Conn) MaxLen(length int) {
	// prevent panics
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.maxLen = length
}

// ReadLine reads until '\n' and returns bytes read and possible error.
func (c *Conn) ReadLine() ([]byte, error) {
	return ReadTill(c, c.maxLen, '\n')
}

// ReadTill reads till 'delim' and returns bytes read and possible error.
func (c *Conn) ReadTill(delim byte) ([]byte, error) {
	return ReadTill(c, c.maxLen, delim)
}

// TLDR just supply a string or []byte
// WriteLine writes a line to the Connection.
// t can be anything convertable to []byte (see ToBytes function)
// ToBytes will panic if it fails to convert to bytes
func (c *Conn) WriteLine(t interface{}) error {
	return WriteLine(c, t)
}
