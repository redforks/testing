package testdb_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMongotest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mongotest Suite")
}
