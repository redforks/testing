package iotest

import (
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

// Assert file content equal to expected. Also failed on IO error.
func AssertFileContent(t assert.TestingT, expected, path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, expected, string(buf))
}
