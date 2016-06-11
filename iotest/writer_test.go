package iotest_test

import (
	"bytes"
	"github.com/redforks/testing/reset"
	"io"
	"time"

	. "github.com/redforks/testing/iotest"

	bdd "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func validErrorWrite(w io.Writer, dataLen, writes int) {
	n, err := w.Write(make([]byte, dataLen))
	Ω(err).Should(MatchError(ErrWriter))
	Ω(n).Should(Equal(writes), "ErrorWriter should write %d bytes", writes)
}

func validErrorWriteSuccess(w io.Writer, dataLen int) {
	Ω(w.Write(make([]byte, dataLen))).Should(Equal(dataLen), "ErrorWriter should write succeed")
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
		validErrorWrite(w, 10, 0)

		w = ErrorWriter(5)
		validErrorWriteSuccess(w, 3)
		validErrorWrite(w, 10, 2)
		validErrorWrite(w, 10, 0)

		w = ErrorWriter(5)
		validErrorWriteSuccess(w, 5)
		validErrorWrite(w, 10, 0)
	})

	bdd.It("SlowWriter", func() {
		buf := &bytes.Buffer{}
		w := NewSlowWriter(buf, 15*time.Millisecond)
		tStart := time.Now()
		Ω(w.Write([]byte("abc"))).Should(Equal(3))
		tEnd := time.Now()
		d := tEnd.Sub(tStart)
		Ω(buf.Bytes()).Should(Equal([]byte("abc")))
		Ω(d > 15*time.Millisecond).Should(BeTrue())
		Ω(d < 25*time.Millisecond).Should(BeTrue())

		tStart = time.Now()
		Ω(w.Write([]byte("cde"))).Should(Equal(3))
		tEnd = time.Now()
		Ω(buf.Bytes()).Should(Equal([]byte("abccde")))
		d = tEnd.Sub(tStart)
		Ω(d > 15*time.Millisecond).Should(BeTrue())
		Ω(d < 25*time.Millisecond).Should(BeTrue())
	})

})
