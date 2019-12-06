package reset_test

import (
	"testing"

	. "github.com/redforks/testing/reset"
	"github.com/stretchr/testify/assert"
)

func TestTestMode(t *testing.T) {
	assert.True(t, TestMode())
}
