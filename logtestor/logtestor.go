package logtestor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// LogTestor is test helper that asserts gathered logs.
type LogTestor struct {
	// Internal Log buffer
	Log []string
}

// New returns a new LogTestor
func New() *LogTestor {
	return &LogTestor{}
}

// Append new string to log
func (l *LogTestor) Append(s string) {
	l.Log = append(l.Log, s)
}

// Assert log content, clear internal log after assert.
func (l *LogTestor) Assert(t *testing.T, exp ...string) {
	bak := l.Log
	l.Log = nil
	assert.Equal(t, exp, bak)
}

// Assert log buffer is empty.
func (l *LogTestor) AssertEmpty(t *testing.T) {
	l.Assert(t)
}
