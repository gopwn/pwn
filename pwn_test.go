package pwn

import (
	"bytes"
	"testing"
)

// Test the ToBytes function in misc.go
func TestToBytes(t *testing.T) {
	t.Parallel()
	var tt = []struct {
		// name of the subtest to execute
		name string

		// expected inputs and outputs
		input  interface{}
		output []byte

		// do you expect the conversion to fail?
		fail bool
	}{
		{
			name:   "test string",
			input:  "hello",
			output: []byte("hello"),
		},
		{
			// may be convertable some day
			name:  "test int",
			input: 42,

			fail: true,
		},
		{
			name:   "test byte",
			input:  byte(5),
			output: []byte{5},
		},
		{
			name:  "test struct",
			input: struct{}{},
			fail:  true,
		},
		{
			name:   "test rune",
			input:  'h',
			output: []byte{'h'},
		},
		{
			name:   "test []byte",
			input:  []byte{131, 41, 48},
			output: []byte{131, 41, 48},
		},
	}

	// execute the tests
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				e := recover()
				// if the error is not nil and it is not expected then fail
				if e != nil && !tc.fail {
					t.Fatal(e)
				}
			}()

			b := Bytes(tc.input)
			if !bytes.Equal(b, tc.output) {
				t.Fatalf("%q != %q", b, tc.output)
			}
		})
	}
}
