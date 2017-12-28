package circular

import (
	"errors"
)

// Buffer represents a circular Buffer
type Buffer struct {
	size        int    //size of buffer
	n           int    //capacity
	data        []byte // data of the buffer
	write, read int    // position of last write or read operation
}

// NewBuffer creates a new buffer of the given size
func NewBuffer(size int) *Buffer {

	return &Buffer{
		size: size,
		data: make([]byte, size),
	}
}

// ReadByte returns the Bytes of the current buffer element
func (b *Buffer) ReadByte() (byte, error) {
	if b.n == 0 {
		return 0, errors.New("no bytes read")
	}
	d := b.data[b.read]
	b.read = (b.read + 1) % b.size
	b.n--
	return d, nil
}

// WriteByte writes the provided byte to the next free buffer postion
func (b *Buffer) WriteByte(c byte) error {
	if b.write == b.read && b.n > 0 {
		return errors.New("buffer is full")
	}
	b.Overwrite(c)
	return nil
}

// Overwrite writes the provided byte to the current buffer position
func (b *Buffer) Overwrite(c byte) {
	if b.read == b.write && b.n > 0 {
		b.read = (b.read + 1) % b.size
		b.n--
	}
	b.data[b.write] = c
	b.write = (b.write + 1) % b.size
	b.n++
}

// Reset clears all elements of the buffer
func (b *Buffer) Reset() {
	*b = Buffer{
		size: b.size,
		data: make([]byte, b.size),
	}
}
