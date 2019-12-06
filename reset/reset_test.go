package reset_test

import (
	"testing"

	myTesting "github.com/redforks/testing"
	"github.com/redforks/testing/logtestor"
	. "github.com/redforks/testing/reset"
	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	var log *logtestor.LogTestor

	beforeEach := func() {
		log = logtestor.New()
	}

	afterEach := func() {
		if Enabled() {
			Disable()
		}
	}

	resetA := func() { log.Append("a") }

	resetB := func() { log.Append("b") }

	newTest := func(f func(t *testing.T)) func(t *testing.T) {
		return myTesting.SetupTeardown(beforeEach, afterEach, f)
	}

	newEnabledTest := func(f func(t *testing.T)) func(t *testing.T) {
		return myTesting.SetupTeardown(func() {
			beforeEach()
			Enable()
			Add(resetA)
		}, afterEach, f)
	}

	t.Run("Not Enabled", newTest(func(t *testing.T) {
		assert.False(t, Enabled())

		Add(resetA)
		Add(resetB)

		log.AssertEmpty(t)
	}))

	t.Run("Set disabled disabled", newTest(func(t *testing.T) {
		assert.Panics(t, Disable)
	}))

	t.Run("Enable/Disable", newEnabledTest(func(t *testing.T) {
		assert.True(t, Enabled())
		log.AssertEmpty(t)
		Disable()
		assert.False(t, Enabled())
		log.Assert(t, "a")
	}))

	t.Run("Execute by reversed order", newEnabledTest(func(t *testing.T) {
		Add(resetB)
		Disable()
		log.Assert(t, "b", "a")
	}))

	t.Run("Dup action", newEnabledTest(func(t *testing.T) {
		Add(resetA)
		Add(resetA)
		Disable()
		log.Assert(t, "a", "a", "a")
	}))

	t.Run("Not allow Add() while executing", newEnabledTest(func(t *testing.T) {
		Add(func() {
			Add(resetA)
		})
		assert.Panics(t, Disable)
	}))
}
