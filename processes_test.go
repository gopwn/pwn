package pwn

import (
	"bytes"
	"testing"
	"time"
)

func TestProcesses(t *testing.T) {
	// i'm gonna try a new testing idea i just had
	type testcase struct {
		// name of the test
		name string
		// test to be ran by t.Run
		runFunc func(t *testing.T)
	}

	testcases := []testcase{
		{
			name: "test echo",
			runFunc: func(t *testing.T) {
				expected := []byte("Hello, world!")
				p, err := Spawn("echo", "Hello, world!")
				if err != nil {
					t.Fatal(err)
				}

				output, err := p.ReadLine(time.Second)
				if !bytes.Equal(output, expected) {
					t.Fatalf("%q != %q", output, expected)
				}
			},
		},
		{
			// test comunicating with a shell
			name: "test sh",
			runFunc: func(t *testing.T) {
				// i'm not sure what the output will be, i'll just let it fail
				// then i'll change this
				expected := []byte("Hello, world")
				p, err := Spawn("sh")
				if err != nil {
					t.Fatal(err)
				}

				err = p.WriteLine("echo Hello, world")
				if err != nil {
					t.Fatal(err)
				}

				out, err := p.ReadLine(time.Second)
				if !bytes.Equal(out, expected) {
					t.Fatalf("%q != %q", out, expected)
				}
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, test.runFunc)
	}
}
