package reset_test

import (
	"testing"

	myTesting "github.com/redforks/testing"
	. "github.com/redforks/testing/reset"
	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	var log []string

	beforeEach := func() {
		log = []string{}
	}

	afterEach := func() {
		if Enabled() {
			Disable()
		}
	}

	resetA := func() {
		log = append(log, `a`)
	}

	resetB := func() {
		log = append(log, `b`)
	}

	assertLog := func(exp ...string) {
		assert.Equal(t, exp, log)
	}

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

		assert.Empty(t, log)
	}))

	t.Run("Set disabled disabled", newTest(func(t *testing.T) {
		assert.Panics(t, Disable)
	}))

	t.Run("Enable/Disable", newEnabledTest(func(t *testing.T) {
		assert.True(t, Enabled())
		assert.Empty(t, log)
		Disable()
		assert.False(t, Enabled())
		assertLog("a")
	}))

	t.Run("Execute by reversed order", newEnabledTest(func(t *testing.T) {
		Add(resetB)
		Disable()
		assertLog("b", "a")
	}))

	t.Run("Dup action", newEnabledTest(func(t *testing.T) {
		Add(resetA)
		Add(resetA)
		Disable()
		assertLog("a", "a", "a")
	}))

	t.Run("Not allow Add() while executing", newEnabledTest(func(t *testing.T) {
		Add(func() {
			Add(resetA)
		})
		assert.Panics(t, Disable)
	}))
}
