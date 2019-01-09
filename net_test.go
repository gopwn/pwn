package pwn

import (
	"bytes"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Test the client reading from the connection
// using my custom ReadTill method.
func TestReadTill(t *testing.T) {
	// readtill testcases
	var testcases = []struct {
		send     []byte
		expected []byte
		delim    byte

		// the max length if not supplied default is used
		maxLen int
	}{
		{
			send:     []byte("Hello\nThere!"),
			expected: []byte("Hello"),
			delim:    '\n',
		},
		{
			send:     []byte("Hey there"),
			expected: []byte("Hey"),
			delim:    ' ',
		},
		{
			send:     []byte("AAAAAAAABBBBBBBBB"),
			expected: []byte("AAAAAAAA"),
			delim:    'B',
		},
		{
			send:     []byte("foobar"),
			expected: []byte("foo"),
			// delim will never be reached so maxLen better work
			delim: '\n',

			maxLen: 3,
		},
	}

	var serverConn net.Conn
	for _, tc := range testcases {
		port := randPort()

		// get the client connection
		go func() {
			l, err := net.Listen("tcp", "127.0.0.1:"+port)
			if err != nil {
				t.Fatal(err)
			}

			defer l.Close()
			serverConn, err = l.Accept()
			_, err = serverConn.Write(tc.send)
			if err != nil {
				t.Fatal(err)
			}
		}()

		// add a delay to give the server time to start up
		time.Sleep(100 * time.Millisecond)

		// connect to the server, in a function so defer works
		func() {
			c, err := Dial("tcp", "127.0.0.1:"+port)
			if err != nil {
				t.Fatal(err)
			}
			defer c.Close()
			if tc.maxLen != 0 {
				c.MaxLen(tc.maxLen)
			}

			// call ReadTill
			output, err := c.ReadTill(tc.delim)

			// error checking (very ugly please send help)
			if tc.maxLen >= len(tc.send) && err != ErrMaxLen {
				t.Fatalf("expected ErrMaxLen got: %v", err)
			}
			if err != nil && err != ErrMaxLen {
				t.Fatal(err)
			}

			// check that output is equal to the expected output
			if !bytes.Equal(output, tc.expected) {
				// if it fails print both, since i'm using text i can do %q
				// but if i was using bytes %X should be used to print the hex.
				t.Fatalf("%q != %q", output, tc.expected)
			}
		}()
	}
}

// get a random port from min 1024 to max 65535
func randPort() string {
	var port int
	for port <= 1024 {
		port = rand.Intn(65535)
	}
	return strconv.Itoa(port)
}
