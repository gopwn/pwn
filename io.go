// Some useful io utils, because io/ioutil is not enough

package pwn

import (
	"io"
)

// Default max length for the ReadTill function.
const MaxLenDefault = 256

// ReadByte reads one byte from r and returns it. if reading one byte fails
// it will return a ErrShortRead error.
func ReadByte(r io.Reader) (byte, error) {
	var buf [1]byte

	// read into buf
	_, err := io.ReadFull(r, buf[:])
	if err != nil {
		return 0, err
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

// copyChan uses normal io.Copy except if it errors it goes through a channel
func copyChan(out io.Writer, in io.Reader, errChan chan error) {
	_, err := io.Copy(out, in)
	if err != nil {
		errChan <- err
	}
}
