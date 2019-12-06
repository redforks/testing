package reset_test

import (
	"testing"

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
		log.Assert(t, "onReset", "onRecover")
	}))

	t.Run("nil", newTest(func(t *testing.T) {
		Register(nil, nil)

		Enable()
		Disable()
		log.AssertEmpty(t)
	}))

	t.Run("Two", newTest(func(t *testing.T) {
		Register(func() {
			log.Append("onReset 1")
		}, func() {
			log.Append("onRecover 1")
		})
		log.Assert(t, "onRecover 1")

		Register(func() {
			log.Append("onReset 2")
		}, func() {
			log.Append("onRecover 2")
		})
		log.Assert(t, "onRecover 2")

		Enable()
		Disable()
		log.Assert(t, "onReset 2", "onReset 1", "onRecover 1", "onRecover 2")
	}))
}
