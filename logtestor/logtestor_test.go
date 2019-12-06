package logtestor_test

import (
	"testing"

	"github.com/redforks/testing/logtestor"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	l := logtestor.New()
	assert.Empty(t, l.Log)

	l.Append("foo")
	l.Append("bar")
	assert.Equal(t, []string{"foo", "bar"}, l.Log)
}

func TestAssert(t *testing.T) {
	l := logtestor.New()
	l.Assert(t)

	l.Append("foo")
	l.Append("bar")
	l.Assert(t, "foo", "bar")

	l.AssertEmpty(t)
}
