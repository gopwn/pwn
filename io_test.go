package pwn

import (
	"bytes"
	"errors"
	"io"
	"testing"
	"time"
)

// TestReadTill tests the ReadTill function in io.go
func TestReadTill(t *testing.T) {
	t.Parallel()
	var testcases = []struct {
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

	for _, tc := range testcases {
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

// badwriter returns its own string as an error when write is called.
type badReader string

func (b badReader) Read([]byte) (int, error) {
	return 0, errors.New(string(b))
}

func Test_copyChan(t *testing.T) {

	var testcases = []struct {
		wantOut string
		in      io.Reader
		errChan chan error
		wantErr string
	}{
		{
			in:      bytes.NewBufferString("hello"),
			errChan: make(chan error),
			wantOut: "hello",

			// we don't want an error.
			wantErr: "",
		},
		{
			in:      badReader("bad reader"),
			errChan: make(chan error),
			wantOut: "",
			wantErr: "bad reader",
		},
	}
	for _, tc := range testcases {
		out := new(bytes.Buffer)
		go copyChan(out, tc.in, tc.errChan)
		// if we expected the error then its okay
		select {
		case err := <-tc.errChan:
			if err != nil && err.Error() != tc.wantErr {
				t.Fatal(err)
			}
		case <-time.After(time.Second):
			break
		}

		if gotOut := out.String(); gotOut != tc.wantOut {
			t.Errorf("copyChan() = %v, want %v", gotOut, tc.wantOut)
		}
	}
}
