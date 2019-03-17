package testing

import "testing"

// SetupTeardown add setup/teardown support to nested testing.T.Run() sub
// tests.
//
// Wrap sub test function, run setup before sub test, defer teardown after
// sub test.
//
// Both setup and teardown can be nil.
func SetupTeardown(setup, teardown func(), f func(t *testing.T)) func(t *testing.T) {
	return func(t *testing.T) {
		if setup != nil {
			setup()
		}

		if teardown != nil {
			defer teardown()
		}

		f(t)
	}
}
