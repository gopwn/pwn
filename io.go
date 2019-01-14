// Some useful io utils, because io/ioutil is not enough

package pwn

import (
	"context"
	"errors"
	"io"
)

// ErrMaxLen indecates that the max length was reached for ReadTill.
var ErrMaxLen = errors.New("max length reached")

// ErrNilReader indecates that a nil reader was supplied.
var ErrNilReader = errors.New("nil reader")

// MaxLenDefault is the max length default for the ReadTill function.
const MaxLenDefault = 256

// ReadByte reads one byte from r and returns it,
// if it fails it will return io.ErrUnexpectedEOF.
func ReadByte(r io.Reader) (byte, error) {
	var buf [1]byte

	// read into buf
	_, err := io.ReadFull(r, buf[:])
	return buf[0], err
}

// ReadTill reads till 'delim' (non inclusive) and returns bytes read and possible error.
// if maxLen is <= 0 it will use MaxLenDefault.
func ReadTill(r io.Reader, maxLen int, delim byte) ([]byte, error) {
	return ReadTillContext(r, maxLen, delim, context.Background())
}

// this function's params are very long, i don't want to create a struct
// just for it though, should not be too long when calling it

// ReadTill reads till 'delim' (non inclusive) or ctx.Done()
// and returns bytes read and possible error.
// if maxLen is <= 0 it will use MaxLenDefault.
func ReadTillContext(r io.Reader, maxLen int, delim byte,
	ctx context.Context) (ret []byte, err error) {

	if maxLen <= 0 {
		maxLen = MaxLenDefault
	}
	if r == nil {
		return ret, ErrNilReader
	}

	for {
		select {
		case <-ctx.Done():
			return ret, err
		default:
		}

		// read one byte
		b, err := ReadByte(r)
		if err != nil {
			return ret, err
		}

		// if the byte is equal to delim stop reading
		if b == delim {
			break
		}

		// append the byte to ret
		ret = append(ret, b)
		if len(ret) >= maxLen {
			return ret, ErrMaxLen
		}
	}

	return ret, nil
}

// WriteLine writes a line to the writer
// it will panic if it ToBytes fails to convert t to []byte
func WriteLine(w io.Writer, t interface{}) error {
	// convert t to bytes
	b := Bytes(t)

	// add the newline, we are "WriteLine" after all!
	b = append(b, '\n')
	_, err := w.Write(b)
	return err
}
