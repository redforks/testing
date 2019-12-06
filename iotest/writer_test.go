package iotest_test

import (
	"bytes"
	"io"
	"testing"
	"time"

	. "github.com/redforks/testing/iotest"
	"github.com/stretchr/testify/assert"
)

func validErrorWrite(t *testing.T, w io.Writer, dataLen, writes int) {
	n, err := w.Write(make([]byte, dataLen))
	assert.Equal(t, ErrWriter, err)
	assert.Equal(t, writes, n)
}

func validErrorWriteSuccess(t *testing.T, w io.Writer, dataLen int) {
	buf := make([]byte, dataLen)
	assertWrite(t, w, buf)
}

func TestErrorWriter(t *testing.T) {
	w := ErrorWriter(0)
	validErrorWrite(t, w, 10, 0)

	w = ErrorWriter(5)
	validErrorWriteSuccess(t, w, 3)
	validErrorWrite(t, w, 10, 2)
	validErrorWrite(t, w, 10, 0)

	w = ErrorWriter(5)
	validErrorWriteSuccess(t, w, 5)
	validErrorWrite(t, w, 10, 0)
}

func assertWrite(t *testing.T, w io.Writer, buf []byte) {
	n, err := w.Write(buf)
	assert.NoError(t, err)
	assert.Equal(t, len(buf), n)
}

func TestSlowWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewSlowWriter(buf, 15*time.Millisecond)
	tStart := time.Now()
	assertWrite(t, w, []byte("abc"))
	tEnd := time.Now()
	assert.Equal(t, []byte("abc"), buf.Bytes())
	assert.WithinDuration(t, tStart.Add(15*time.Millisecond), tEnd, 10*time.Millisecond)

	tStart = time.Now()
	assertWrite(t, w, []byte("cde"))
	tEnd = time.Now()
	assert.Equal(t, []byte("abccde"), buf.Bytes())
	assert.WithinDuration(t, tStart.Add(15*time.Millisecond), tEnd, 10*time.Millisecond)
}
