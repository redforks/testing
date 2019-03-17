package testing_test

import (
	"testing"

	myTesting "github.com/redforks/testing"
	"github.com/stretchr/testify/assert"
)

func TestSetupTeardown(t *testing.T) {
	t.Run("Setup", func(t *testing.T) {
		hit, setupHit := 0, 0
		tf := func(innerT *testing.T) {
			hit++
			assert.Equal(t, t, innerT)
		}
		setup := func() {
			setupHit++
		}
		f := myTesting.SetupTeardown(setup, nil, tf)

		f(t)

		assert.Equal(t, 1, setupHit)
		assert.Equal(t, 1, hit)
	})

	t.Run("Teardown on no error", func(t *testing.T) {
		hit, tearDownHit := 0, 0
		tf := func(innerT *testing.T) {
			hit++
		}
		tearDown := func() {
			tearDownHit++
		}
		f := myTesting.SetupTeardown(nil, tearDown, tf)

		f(t)
		assert.Equal(t, 1, hit)
		assert.Equal(t, 1, tearDownHit)
	})

	t.Run("Teardown on panic", func(t *testing.T) {
		tearDownHit := 0
		tf := func(innerT *testing.T) {
			panic("foo")
		}
		tearDown := func() {
			tearDownHit++
		}
		f := myTesting.SetupTeardown(nil, tearDown, tf)

		assert.PanicsWithValue(t, "foo", func() {
			f(t)
		})

		assert.Equal(t, 1, tearDownHit)
	})
}
