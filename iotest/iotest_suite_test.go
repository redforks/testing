package iotest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIotest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iotest Suite")
}
