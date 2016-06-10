package reset_test

import (
	"testing"

	. "github.com/redforks/testing/reset"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reset")
}

var _ = Describe("Reset", func() {
	var log []string

	resetA := func() {
		log = append(log, `a`)
	}

	resetB := func() {
		log = append(log, `b`)
	}

	BeforeEach(func() {
		log = []string{}
	})

	AfterEach(func() {
		if Enabled() {
			Disable()
		}
	})

	It("Not Enabled", func() {
		Ω(Enabled()).Should(BeFalse())
		Add(resetA)
		Add(resetB)

		Ω(log).Should(BeEmpty())
	})

	It("Set disabled disabled", func() {
		Ω(func() {
			Disable()
		}).Should(Panic())
	})

	Context("Enabled", func() {

		BeforeEach(func() {
			Enable()
			Add(resetA)
		})

		It("Enabled", func() {
			Ω(Enabled()).Should(BeTrue())
			Ω(log).Should(BeEmpty())
			Disable()
			Ω(Enabled()).Should(BeFalse())
			Ω(log).Should(Equal([]string{"a"}))
		})

		It("Execute by reversed order", func() {
			Add(resetB)
			Disable()
			Ω(log).Should(Equal([]string{"b", "a"}))
		})

		It("Add dup action", func() {
			Add(resetA)
			Add(resetA)
			Disable()
			Ω(log).Should(Equal([]string{"a"}))
		})

		It("Not allow Add() while executing", func() {
			Add(func() {
				Add(resetA)
			})
			Ω(func() {
				Disable()
			}).Should(Panic())
		})

	})

})
