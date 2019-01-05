package pwn

import (
	"errors"
	"net"
	"sync"
	"time"
)

// ErrMaxLen indecates that the max length was reached for ReadTill
var ErrMaxLen = errors.New("max length reached")

// ErrShortRead indecates a short read error.
type ErrShortRead struct {
	// the string to be returned by Error()
	err string
}

func (e ErrShortRead) Error() string {
	return e.err
}

// Conn is a generic stream-oriented network connection.
//
// Multiple goroutines may invoke methods on a Conn simultaneously.
type Conn interface {
	// Read reads data from the connection.
	// Read can be made to time out and return an Error with Timeout() == true
	// after a fixed time limit; see SetDeadline and SetReadDeadline.
	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return an Error with Timeout() == true
	// after a fixed time limit; see SetDeadline and SetWriteDeadline.
	Write(b []byte) (n int, err error)

	// Close closes the connection.
	// Any blocked Read or Write operations will be unblocked and return errors.
	Close() error

	// LocalAddr returns the local network address.
	LocalAddr() net.Addr

	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr

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
	SetDeadline(t time.Time) error

	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	// A zero value for t means Read will not time out.
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	SetWriteDeadline(t time.Time) error

	// ReadLine reads until '\n' and returns bytes read and possible error.
	ReadLine() ([]byte, error)

	// ReadTill reads till delim and returns bytes read and possible error.
	ReadTill(delim byte) ([]byte, error)

	// MaxLen sets the maximum length for ReadTill and ReadLine.
	MaxLen(length int)
}

// conn is the underlying struct returned by pwn.Dial etc
type conn struct {
	// c is the underlying connection
	c net.Conn

	// mu is used for protecting struct variables from concurrent reads / writes
	mu sync.Mutex

	// the max length for ReadLine and ReadTill.
	maxLen int
}

func (c *conn) MaxLen(length int) {
	// prevent panics
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.maxLen = length
}

// ReadLine reads until '\n' and returns bytes read and possible error.
func (c conn) ReadLine() ([]byte, error) {
	return c.ReadTill('\n')
}

// ReadTill reads till 'delim' and returns bytes read and possible error.
func (c conn) ReadTill(delim byte) ([]byte, error) {
	// the final return value will be stored in here.
	var retval []byte

	for {
		// read one byte
		b, err := ReadByte(c.c)
		if err != nil {
			return retval, err
		}

		// if the byte is equal to delim stop reading
		if b == delim {
			break
		}

		// append the byte to retval
		retval = append(retval, b)
		if len(retval) >= c.maxLen {
			return retval, ErrMaxLen
		}
	}

	return retval, nil
}

// Below are the methods for the net.Conn interface.

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c conn) Read(buf []byte) (int, error) { return c.c.Read(buf) }

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c conn) Write(data []byte) (int, error) { return c.c.Write(data) }

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c conn) Close() error { return c.c.Close() }

// LocalAddr returns the local network address.
func (c conn) LocalAddr() net.Addr { return c.c.LocalAddr() }

// RemoteAddr returns the remote network address.
func (c conn) RemoteAddr() net.Addr { return c.c.RemoteAddr() }

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
func (c conn) SetDeadline(t time.Time) error { return c.c.SetDeadline(t) }

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c conn) SetReadDeadline(t time.Time) error { return c.c.SetReadDeadline(t) }

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c conn) SetWriteDeadline(t time.Time) error { return c.c.SetWriteDeadline(t) }
