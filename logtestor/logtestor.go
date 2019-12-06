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

// Assert log content.
func (l *LogTestor) Assert(t *testing.T, exp ...string) {
	assert.Equal(t, exp, l.Log)
}

// Assert log buffer is empty.
func (l *LogTestor) AssertEmpty(t *testing.T) {
	l.Assert(t)
}
