package reset

import (
	bdd "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = bdd.Describe("reset - recover", func() {
	var (
		log       = ""
		appendLog = func(msg string) {
			log += msg + "\n"
		}
		assertLog = func(expected string) {
			assert.Equal(t(), expected, log)
		}
	)

	bdd.BeforeEach(func() {
		log = ""
	})

	bdd.AfterEach(func() {
		resetFns, recoverFns = nil, nil
		if Enabled() {
			Disable()
		}
	})

	bdd.It("One", func() {
		Register(func() {
			appendLog("onReset")
		}, func() {
			appendLog("onRecover")
		})

		Enable()
		assertLog("")

		Disable()
		assertLog("onReset\nonRecover\n")

		Enable()
		Disable()
		// Register once, effect for ever.
		assertLog("onReset\nonRecover\nonReset\nonRecover\n")
	})

	bdd.It("nil", func() {
		Register(nil, nil)

		Enable()
		Disable()
		assertLog("")
	})

	bdd.It("Two", func() {
		Register(func() {
			appendLog("onReset 1")
		}, func() {
			appendLog("onRecover 1")
		})
		Register(func() {
			appendLog("onReset 2")
		}, func() {
			appendLog("onRecover 2")
		})

		Enable()
		Disable()
		assertLog("onReset 2\nonReset 1\nonRecover 1\nonRecover 2\n")
	})

})
