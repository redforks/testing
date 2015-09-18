package iotest

import (
	. "github.com/onsi/ginkgo"

	"testing"
)

var t = GinkgoT

func TestIotest(t *testing.T) {
	RunSpecs(t, "Iotest Suite")
}
