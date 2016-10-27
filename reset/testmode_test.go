package reset_test

import (
	. "github.com/redforks/testing/reset"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testmode", func() {

	It("TestMode", func() {
		Ω(TestMode()).Should(BeTrue())
	})

})
