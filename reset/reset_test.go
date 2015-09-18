package reset

import (
	"sync/atomic"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var t = GinkgoT

func TestRest(t *testing.T) {
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

	It("Not Enabled", func() {
		assert.False(t(), Enabled())
		Add(resetA)
		Add(resetB)

		assert.Empty(t(), log)
	})

	It("Set disabled disabled", func() {
		assert.Panics(t(), func() {
			Disable()
		})
	})

	Context("Enabled", func() {

		BeforeEach(func() {
			Enable()
			Add(resetA)
		})

		It("Enabled", func() {
			assert.True(t(), Enabled())
			assert.Empty(t(), log)
			Disable()
			assert.False(t(), Enabled())
			assert.Equal(t(), []string{"a"}, log)
		})

		It("Execute by reversed order", func() {
			Add(resetB)
			Disable()
			assert.Equal(t(), []string{"b", "a"}, log)
		})

		It("Add dup action", func() {
			Add(resetA)
			Add(resetA)
			Disable()
			assert.Equal(t(), []string{"a"}, log)
		})

		It("Not allow Add() while executing", func() {
			Add(func() {
				Add(resetA)
			})
			assert.Panics(t(), func() {
				Disable()
			})
			atomic.StoreInt32((*int32)(&state), int32(st_disabled))
		})

	})

})
