package matcher

import (
	"log"

	"github.com/onsi/ginkgo"
)

func init() {
	// Set log using GinkgoWriter, see we can log message when ginkgo unit test failed
	log.SetOutput(ginkgo.GinkgoWriter)
	log.SetFlags(log.Lshortfile)
}
