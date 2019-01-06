// Some useful io utils, because io/ioutil is not enough

package pwn

import (
	"context"
	"io"
)

// Default max length for the ReadTill function.
const MaxLenDefault = 256

// ReadByte reads one byte from r and returns it. if reading one byte fails
// it will return a ErrShortRead error.
func ReadByte(r io.Reader) (byte, error) {
	var buf [1]byte

	// read into buf
	nr, err := r.Read(buf[:])
	if err != nil {
		return 0, err
	}

	// make sure we've read at least 1 byte
	if nr < 1 {
		return 0, ErrShortRead{"ReadByte: failed to read byte (nr < 1)"}
	}

	return buf[0], nil
}

// ReadTill reads till 'delim' and returns bytes read and possible error.
// if maxLen is <= 0 it will use MaxLenDefault
func ReadTill(r io.Reader, maxLen int, delim byte) (ret []byte, err error) {
	if maxLen <= 0 {
		maxLen = MaxLenDefault
	}

	for {
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

// ReadTillContext reads from r until maxLen delim or ctx.Done()
func ReadTillContext(r io.Reader, maxLen int, delim byte, ctx context.Context) (ret []byte, err error) {
	return nil, nil
}
