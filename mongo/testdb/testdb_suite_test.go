package testdb

import (
	. "github.com/onsi/ginkgo"

	"testing"
)

var t = GinkgoT

func TestMongotest(t *testing.T) {
	RunSpecs(t, "Mongotest Suite")
}
