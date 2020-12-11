package bytes

import (
	"io"
	"syscall"
)

// ---------------------------------------------------

// Reader implement io.Reader, io.ReadAt
type Reader struct {
	b   []byte
	off int
}

// NewReader new reader with byte
func NewReader(val []byte) *Reader {
	return &Reader{val, 0}
}

// Len length of bytes
func (r *Reader) Len() int {
	if r.off >= len(r.b) {
		return 0
	}
	return len(r.b) - r.off
}

// Bytes to bytes
func (r *Reader) Bytes() []byte {
	return r.b[r.off:]
}

// SeekToBegin move pointer to begin, means can read again
func (r *Reader) SeekToBegin() (err error) {
	r.off = 0
	return
}

// Seek can read from specified pos
func (r *Reader) Seek(offset int64, whence int) (ret int64, err error) {
	switch whence {
	case 0:
	case 1:
		offset += int64(r.off)
	case 2:
		offset += int64(len(r.b))
	default:
		err = syscall.EINVAL
		return
	}
	if offset < 0 {
		err = syscall.EINVAL
		return
	}
	if offset >= int64(len(r.b)) {
		r.off = len(r.b)
	} else {
		r.off = int(offset)
	}
	ret = int64(r.off)
	return
}

// Read read data from reader
func (r *Reader) Read(val []byte) (n int, err error) {
	n = copy(val, r.b[r.off:])
	if n == 0 && len(val) != 0 {
		err = io.EOF
		return
	}
	r.off += n
	return
}

// Close close reader
func (r *Reader) Close() (err error) {
	return
}

// ---------------------------------------------------

// Writer implement io.Writer, io.WriteAt
type Writer struct {
	b []byte
	n int
}

// NewWriter new writer with bytes
func NewWriter(buff []byte) *Writer {
	return &Writer{buff, 0}
}

// Write write byte to writer
func (p *Writer) Write(val []byte) (n int, err error) {
	n = copy(p.b[p.n:], val)
	if n == 0 && len(val) > 0 {
		err = io.EOF
		return
	}
	p.n += n
	return
}

// Len length of writer content
func (p *Writer) Len() int {
	return p.n
}

// Bytes to bytes
func (p *Writer) Bytes() []byte {
	return p.b[:p.n]
}

// Reset set pointer to zero
func (p *Writer) Reset() {
	p.n = 0
}

// ---------------------------------------------------

// Buffer ...
type Buffer struct {
	b []byte
}

// NewBuffer ..
func NewBuffer() *Buffer {
	return new(Buffer)
}

// ReadAt implement io.ReadAt
func (p *Buffer) ReadAt(buf []byte, off int64) (n int, err error) {
	ioff := int(off)
	if len(p.b) <= ioff {
		return 0, io.EOF
	}
	n = copy(buf, p.b[ioff:])
	if n != len(buf) {
		err = io.EOF
	}
	return
}

// WriteAt implement io.WriteAt
func (p *Buffer) WriteAt(buf []byte, off int64) (n int, err error) {
	ioff := int(off)
	iend := ioff + len(buf)
	if len(p.b) < iend {
		if len(p.b) == ioff {
			p.b = append(p.b, buf...)
			return len(buf), nil
		}
		zero := make([]byte, iend-len(p.b))
		p.b = append(p.b, zero...)
	}
	copy(p.b[ioff:], buf)
	return len(buf), nil
}

// WriteStringAt  write string
func (p *Buffer) WriteStringAt(buf string, off int64) (n int, err error) {
	ioff := int(off)
	iend := ioff + len(buf)
	if len(p.b) < iend {
		if len(p.b) == ioff {
			p.b = append(p.b, buf...)
			return len(buf), nil
		}
		zero := make([]byte, iend-len(p.b))
		p.b = append(p.b, zero...)
	}
	copy(p.b[ioff:], buf)
	return len(buf), nil
}

// Truncate truncate buffer
func (p *Buffer) Truncate(fsize int64) (err error) {
	size := int(fsize)
	if len(p.b) < size {
		zero := make([]byte, size-len(p.b))
		p.b = append(p.b, zero...)
	} else {
		p.b = p.b[:size]
	}
	return nil
}

// Buffer ...
func (p *Buffer) Buffer() []byte {
	return p.b
}

// Len length of buffer
func (p *Buffer) Len() int {
	return len(p.b)
}

// ---------------------------------------------------
