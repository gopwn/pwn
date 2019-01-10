package pwn

import (
	"net"
	"sync"
)

// A Listener is a generic network listener for stream-oriented protocols.
// Multiple goroutines may invoke methods on a Listener simultaneously.
type Listener interface {
	// Accept waits for and returns the next Connection to the listener.
	Accept() (Conn, error)

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() net.Addr
}

type listener struct {
	l net.Listener
}

func (l listener) Addr() net.Addr {
	return l.l.Addr()
}

func (l listener) Close() error {
	return l.l.Close()
}

func (l listener) Accept() (Conn, error) {
	rawConn, err := l.l.Accept()
	if err != nil {
		return Conn{}, err
	}

	return Conn{
		rawConn,
		// the default line length to be used with Conn.ReadLine
		MaxLenDefault,

		sync.Mutex{},
	}, nil
}

// Dial creates a new network Connection using net.Dial
// then creates a pwn.Conn using it and returns it
// MaxLineLength is by default set to 256, you can change it in the returned
// Conn using Conn.MaxLen(i int).
func Dial(network, addr string) (Conn, error) {
	rawConn, err := net.Dial(network, addr)
	if err != nil {
		return Conn{}, err
	}

	return Conn{
		rawConn,
		// the default line length to be used with Conn.ReadLine
		MaxLenDefault,

		sync.Mutex{},
	}, nil
}

// Listen creates a net.Listener that will accept Connections
// and wrap them in a pwn.Conn
func Listen(network, addr string) (Listener, error) {
	rawListener, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	return listener{
		l: rawListener,
	}, nil
}
