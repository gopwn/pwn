package pwn

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// mp connects a pwn.Listener with a pwn.Dialer
// c2 is the client c1 is the server
func mp() (c1, c2 net.Conn, stop func(), err error) {
	addr := "127.0.0.1:" + randPort()
	connChan := make(chan net.Conn)
	errChan := make(chan error)
	go func() {
		l, err := Listen("tcp", addr)
		if err != nil {
			errChan <- err
			return
		}
		conn, err := l.Accept()
		if err != nil {
			errChan <- err
			return
		}
		connChan <- conn
	}()

	time.Sleep(20 * time.Millisecond)
	c2, err = Dial("tcp", addr)
	if err != nil {
		return
	}

	// check possible error from the server goroutine
	select {
	case err = <-errChan:
		if err != nil {
			return
		}
	case c1 = <-connChan:
		break
	}

	stop = func() {
		c1.Close()
		c2.Close()
	}
	return
}

// get a random port from min 1024 to max 65535
func randPort() string {
	var port int
	for port <= 1024 {
		port = rand.Intn(65535)
	}
	return strconv.Itoa(port)
}
