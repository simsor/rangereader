package rangereader

import (
	"fmt"
	"io"
)

// RangeReader is an object implementing io.Reader which will only yield data between the given range (inclusive)
type RangeReader struct {
	r          io.Reader
	start, end int

	offset int
}

// New creates a new RangeReader object
func New(r io.Reader, start, end int) (*RangeReader, error) {
	if start >= end {
		return nil, fmt.Errorf("'start' must be strictly smaller than 'end' (%d >= %d)", start, end)
	}

	if start < 0 || end < 0 {
		return nil, fmt.Errorf("'start' and 'end' must be positive")
	}

	if r == nil {
		return nil, fmt.Errorf("the Reader cannot be nil")
	}

	return &RangeReader{
		r:     r,
		start: start,
		end:   end,
	}, nil
}

// Read
func (r *RangeReader) Read(p []byte) (n int, err error) {
	// Handle case where we passed the read bounds already
	if r.offset >= r.end {
		return 0, io.EOF
	}

	// Handle case where we are not yet up to the beggining offset
	buf := make([]byte, 1)
	for r.offset < r.start {
		n, err = r.r.Read(buf)
		if n != 1 {
			err = fmt.Errorf("Did not read the correct number of bytes")
			return
		}

		if err != nil {
			return
		}

		r.offset++
	}

	// If directly reading to "p" would read outside our allowed bounds, we instead read to a temporary, smaller buffer
	offsetToEnd := r.end - r.offset
	buf = p
	if len(p) > offsetToEnd {
		buf = make([]byte, offsetToEnd)
	}

	n, err = r.r.Read(buf)
	r.offset += n

	if len(p) > offsetToEnd {
		for i, b := range buf {
			p[i] = b
		}

		err = io.EOF
	}

	return
}
