// Test Helpers for package `io' and `io/ioutil'.
package ioth

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/assert"
)

func ReadAll(t assert.TestingT, r io.Reader) []byte {
	res, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	return res
}

func WriteFile(t assert.TestingT, filename string, data []byte) {
	assert.NoError(t, ioutil.WriteFile(filename, data, os.ModePerm))
}
