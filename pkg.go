// Package testing define helper / mock for unit testing
//
// Because golang likes return error object instead of exception/panic,
// always handle error return values is a good practise. But sometimes it is
// impossible to got error, such as read from memory buffer, not handler them
// maybe maybe loose error because someday code changes, but handle them needs
// a lot of duplicate codes.
//
// In package testing contains many test helper packages, suffix with `th', to
// handle these never happen errors. Test helper check the error result, if it
// is not nil, using testing.Fatal(err) to log the error object and abort current
// test case execution.
package testing

import (
	"time"
)

// TryWait the action until it returns true, call timeout if timeout.
func TryWait(d time.Duration, try func() bool, timeout func()) {
	tick := int64(d) / 100
	for i := 0; i < 100; i++ {
		if try() {
			return
		}
		time.Sleep(time.Duration(tick))
	}
	timeout()
}
