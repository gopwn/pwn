package pwn

import (
	"bytes"
	"io"
	"testing"
)

// TestReadTill tests the ReadTill function in io.go
func TestReadTill(t *testing.T) {
	var testcases = []struct {
		// expected input and expecteds
		input    []byte
		expected []byte

		delim byte
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
	}

	for _, tc := range testcases {
		r := bytes.NewBuffer(tc.input)
		output, err := ReadTill(r, 0, tc.delim)
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}

		if !bytes.Equal(output, tc.expected) {
			t.Fatalf("%q != %q", output, tc.expected)
		}
	}
}
