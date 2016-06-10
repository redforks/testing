package reset_test

import (
	bdd "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/redforks/testing/reset"
)

var _ = bdd.Describe("reset - recover", func() {
	var (
		log       = ""
		appendLog = func(msg string) {
			log += msg + "\n"
		}
		assertLog = func(expected string) {
			Ω(log).Should(Equal(expected))
			log = ""
		}
	)

	bdd.BeforeEach(func() {
		log = ""
	})

	bdd.AfterEach(func() {
		ClearInternal()

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
		assertLog("onRecover\n")

		Enable()
		assertLog("")

		Disable()
		assertLog("onReset\nonRecover\n")

		Enable()
		Disable()
		// Register once, effect for ever.
		assertLog("onReset\nonRecover\n")
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
		assertLog("onRecover 1\n")

		Register(func() {
			appendLog("onReset 2")
		}, func() {
			appendLog("onRecover 2")
		})
		assertLog("onRecover 2\n")

		Enable()
		Disable()
		assertLog("onReset 2\nonReset 1\nonRecover 1\nonRecover 2\n")
	})

})
