package iotest

import (
	"errors"
	"io"
	"time"
)

var (
	// Error object returned by ErrorWriter
	ErrWriter = errors.New("ErrorWrite: no more writes")
)

type errorWriter struct {
	n int
}

// ByteWriterWriter combines ByteWriter and Writer interface that ErrorWriter actually implemented.
type byteWriterWriter interface {
	io.ByteWriter
	io.Writer
}

// Returns a writer that would error after writes n bytes
func ErrorWriter(n int) byteWriterWriter {
	return &errorWriter{n}
}

func (w *errorWriter) Write(data []byte) (int, error) {
	n := w.n - len(data)
	if n >= 0 {
		w.n = n
		return len(data), nil
	}
	n, w.n = w.n, 0
	return n, ErrWriter
}

func (w *errorWriter) WriteByte(b byte) error {
	_, err := w.Write([]byte{b})
	return err
}

type slowWriter struct {
	w io.Writer
	d time.Duration
}

// Create a slow writer. Slow writer wrap a exist writer, each write operation
// will delay a while.
func NewSlowWriter(w io.Writer, delay time.Duration) io.Writer {
	return &slowWriter{w, delay}
}

func (w *slowWriter) Write(data []byte) (int, error) {
	time.Sleep(w.d)
	return w.w.Write(data)
}
