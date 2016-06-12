package matcher_test

import (
	"errors"

	. "github.com/redforks/testing/matcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Panics", func() {

	It("Expected panic", func() {
		Ω(func() {}).ShouldNot(Panics("foo"))
	})

	DescribeTable("Invalid actual value", func(actual interface{}) {
		p := Panics("foo")
		ok, err := p.Match(actual)
		Ω(err).Should(HaveOccurred())
		Ω(ok).Should(BeFalse())
	},
		Entry("nil", nil),
		Entry("not function", 1),
		Entry("Has argument", func(foo int) {}),
		Entry("Has return value", func() int { return 1 }),
	)

	Context("FailureMessage", func() {

		DescribeTable("expected is value", func(actual interface{}, msg, negMsg string) {
			p := Panics("foo")
			Ω(p.Match(actual)).Should(BeFalse())
			Ω(p.FailureMessage(actual)).Should(ContainSubstring(msg))
			Ω(p.NegatedFailureMessage(actual)).Should(ContainSubstring(negMsg))
		},
			Entry("not paniced", func() {},
				"to panic with",
				"not to panic, or not to panic with"),
			Entry("paniced with mismatch value", func() { panic("bar") },
				"but got:",
				"not to panic, or not to panic with"),
		)

	})

	It("paniced with unexpected value", func() {
		Ω(func() { panic("bar") }).ShouldNot(Panics("foo"))
	})

	It("paniced as expected", func() {
		Ω(func() { panic("foo") }).Should(Panics("foo"))
	})

	It("expected is a matcher", func() {
		Ω(func() { panic("foo") }).Should(Panics(Equal("foo")))
	})

	Context("PanicsWithSubstring", func() {

		It("Expected panic", func() {
			Ω(func() {}).ShouldNot(PanicsWithSubstring("foo"))
		})

		It("Paniced with message not contains expected", func() {
			Ω(func() { panic("foo") }).ShouldNot(PanicsWithSubstring("bar"))
		})

		DescribeTable("Succeed", func(val interface{}, substr string) {
			Ω(func() { panic(val) }).Should(PanicsWithSubstring(substr))
		},
			Entry("string", "foobar", "foo"),
			Entry("error", errors.New("foobar"), "foo"),
		)

		DescribeTable("FailurMessage", func(actual func(), msg, negMsg string) {
			p := PanicsWithSubstring("foo")
			Ω(p.Match(actual)).Should(BeFalse())
			Ω(p.FailureMessage(actual)).Should(ContainSubstring(msg))
			Ω(p.NegatedFailureMessage(actual)).Should(ContainSubstring(negMsg))
		},
			Entry("not paniced", func() {}, "to panic contains:\n    <string>: foo",
				"not paniced, or not panic contains:\n    <string>: foo"),
			Entry("paniced not contained substring", func() { panic("bar") },
				"to panic contains:\n    <string>: foo\nbut got:\n    <string>: bar",
				"not paniced, or not panic contains:\n    <string>: foo"),
		)

	})
})
