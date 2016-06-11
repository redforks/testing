package matcher_test

import (
	. "github.com/redforks/testing/matcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Save", func() {

	It("Match", func() {
		var foo int
		m := Save(&foo)
		Ω(m.Match(3)).Should(BeTrue())
		Ω(foo).Should(Equal(3))
	})

	It("Non-pointer", func() {
		var foo int
		m := Save(foo)
		ok, err := m.Match(3)
		Ω(ok).Should(BeFalse())
		Ω(err.Error()).Should(ContainSubstring("You need to pass a pointer!"))
	})

})
