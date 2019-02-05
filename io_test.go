package pwn

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"
)

// TestReadTill testcases the ReadTill function in io.go
func TestReadTill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		// expected input and expecteds
		input    []byte
		expected []byte

		delim byte
		// do we expect ErrMaxLen?
		overMaxLen bool
		maxLen     int
	}{
		{
			input:    []byte("AAAAAAAAAABBBBBBBBBB"),
			expected: []byte("AAAAAAAAAA"),

			delim: 'B',
		},
		{
			input:    []byte("Hello\n World"),
			expected: []byte("Hello"),
			delim:    '\n',
		},
		{
			// What happens with a nil delim?
			input:    []byte("Hello\n World"),
			expected: []byte("Hello\n World"),
		},
		{
			// test max len
			input:      []byte("Hello, World!"),
			expected:   []byte("Hello,"),
			maxLen:     6,
			overMaxLen: true,
		},
	}

	for _, tc := range tests {
		r := bytes.NewBuffer(tc.input)
		output, err := ReadTill(r, tc.maxLen, tc.delim)
		if err != nil && err != io.EOF {
			if !tc.overMaxLen && err != ErrMaxLen {
				t.Fatal(err)
			}
		}

		if !bytes.Equal(output, tc.expected) {
			t.Fatalf("wanted %q got %q", tc.expected, output)
		}
	}

	// test that readtill returns correctly on a nil reader
	t.Run("test nil", func(t *testing.T) {
		_, err := ReadTill(nil, 0, '\n')
		if err != ErrNilReader {
			t.Fatalf("expected ErrNilReader, got: %v", err)
		}
	})
}

// Test context cancelation working correctly.
func TestReadTillContext(t *testing.T) {
	r := sleepyReader{
		Reader: strings.NewReader("hello context!\n"),
		delay:  100 * time.Millisecond,
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	_, err := ReadTillContext(r, 0, '\n', ctx)
	if err != context.Canceled {
		t.Fatalf("wanted %q; got %v", context.Canceled, err)
	}
}

// sleepyReader delays for a set duration before every read call.
type sleepyReader struct {
	// How long to delay on every Read call.
	delay time.Duration

	// Where to read from
	io.Reader
}

func (s sleepyReader) Read(buf []byte) (int, error) {
	time.Sleep(s.delay)
	return s.Reader.Read(buf)
}

// badReader returns its own string as an error when read is called.
type badReader string

func (b badReader) Read([]byte) (int, error) {
	return 0, errors.New(string(b))
}
