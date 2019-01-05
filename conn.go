package pwn

import (
	"net"
	"time"
)

type Conn struct {
	// c is the underlying connection
	c net.Conn
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c Conn) Read(buf []byte) (int, error) { return c.c.Read(buf) }

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c Conn) Write(data []byte) (int, error) { return c.c.Write(data) }

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c Conn) Close(data []byte) error { return c.c.Close() }

// LocalAddr returns the local network address.
func (c Conn) LocalAddr() net.Addr { return c.c.LocalAddr() }

// RemoteAddr returns the remote network address.
func (c Conn) RemoteAddr() net.Addr { return c.c.RemoteAddr() }

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (c Conn) SetDeadline(t time.Time) error { return c.c.SetDeadline(t) }

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c Conn) SetReadDeadline(t time.Time) error { return c.c.SetReadDeadline(t) }

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c Conn) SetWriteDeadline(t time.Time) error { return c.c.SetWriteDeadline(t) }
