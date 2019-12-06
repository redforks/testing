package reset_test

import (
	"testing"

	bdd "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	myTesting "github.com/redforks/testing"
	"github.com/redforks/testing/logtestor"
	. "github.com/redforks/testing/reset"
)

func TestResetRecover(t *testing.T) {
	var log *logtestor.LogTestor

	beforeEach := func() {
		log = logtestor.New()
	}

	afterEach := func() {
		ClearInternal()

		if Enabled() {
			Disable()
		}
	}

	newTest := func(f func(t *testing.T)) func(t *testing.T) {
		return myTesting.SetupTeardown(beforeEach, afterEach, f)
	}

	t.Run("One", newTest(func(t *testing.T) {
		Register(func() {
			log.Append("onReset")
		}, func() {
			log.Append("onRecover")
		})
		log.Assert(t, "onRecover")

		Enable()
		log.AssertEmpty(t)

		Disable()
		log.Assert(t, "onReset", "onRecover")

		Enable()
		Disable()
		// Register once, effect for ever.
		log.Assert(t, "OnReset", "onRecover")
	}))
}

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
