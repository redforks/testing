package iotest

import (
	"bytes"
	"io"
	"time"

	bdd "github.com/onsi/ginkgo"
	"github.com/redforks/testing/reset"
	"github.com/stretchr/testify/assert"
)

func validErrorWrite(t assert.TestingT, w io.Writer, dataLen, writes int) {
	if n, err := w.Write(make([]byte, dataLen)); err == nil {
		t.Errorf(`ErrorWriter should return error`)
	} else if n != writes {
		assert.Equal(t, ErrWriter, err)
		t.Errorf(`ErrorWriter should report write bytes %d, but %d`, writes, n)
	}
}

func validErrorWriteSuccess(t assert.TestingT, w io.Writer, dataLen int) {
	if n, err := w.Write(make([]byte, dataLen)); err != nil {
		assert.Equal(t, ErrWriter, err)
		t.Errorf(`ErrorWriter should succeed %d bytes`, dataLen)
	} else if n != dataLen {
		t.Errorf(`ErrorWriter should report write bytes %d, but %d`, dataLen, n)
	}

}

var _ = bdd.Describe("Writers", func() {

	bdd.BeforeEach(func() {
		reset.Enable()
	})

	bdd.AfterEach(func() {
		reset.Disable()
	})

	bdd.It("ErrorWriter", func() {
		w := ErrorWriter(0)
		validErrorWrite(t(), w, 10, 0)

		w = ErrorWriter(5)
		validErrorWriteSuccess(t(), w, 3)
		validErrorWrite(t(), w, 10, 2)
		validErrorWrite(t(), w, 10, 0)

		w = ErrorWriter(5)
		validErrorWriteSuccess(t(), w, 5)
		validErrorWrite(t(), w, 10, 0)
	})

	bdd.It("SlowWriter", func() {
		buf := &bytes.Buffer{}
		w := NewSlowWriter(buf, 10*time.Millisecond)
		tStart := time.Now()
		w.Write([]byte("abc"))
		tEnd := time.Now()
		assert.Equal(t(), []byte("abc"), buf.Bytes())
		assert.InDelta(t(), 10*int64(time.Millisecond), int64(tEnd.Sub(tStart)), float64(int64(time.Millisecond)))

		w.Write([]byte("cde"))
		tEnd = time.Now()
		assert.Equal(t(), []byte("abccde"), buf.Bytes())
		assert.InDelta(t(), 20*int64(time.Millisecond), int64(tEnd.Sub(tStart)), float64(int64(time.Millisecond)))
	})

})
