package tassert

import (
	. "github.com/onsi/ginkgo"

	"testing"
)

var t = GinkgoT

func TestTassert(t *testing.T) {
	RunSpecs(t, "Tassert Suite")
}
