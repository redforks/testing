// Contains test helper for package `net', `net/http'
package netth

import (
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func NewRequest(t assert.TestingT, method, urlStr string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, urlStr, body)
	assert.NoError(t, err)
	return r
}

func HttpClientDo(t assert.TestingT, client *http.Client, request *http.Request) *http.Response {
	r, err := client.Do(request)
	assert.NoError(t, err)
	return r
}
