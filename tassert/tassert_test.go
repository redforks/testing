package tassert

import (
	"errors"
	"fmt"

	bdd "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

// testingTLogger implement assert.TestingT interface, log my assert functions
// error outputs.
type testingTLogger struct {
	log string
}

func (l *testingTLogger) Errorf(format string, args ...interface{}) {
	l.log += fmt.Sprintf(format, args...)
}

func (l *testingTLogger) Assert(t assert.TestingT, expected string) {
	assert.Regexp(t, expected, l.GetAndClear())
}

func (l *testingTLogger) GetAndClear() string {
	var r string
	r, l.log = l.log, ""
	return r
}

var _ = bdd.Describe("tassert", func() {
	var logger *testingTLogger

	bdd.BeforeEach(func() {
		logger = &testingTLogger{}
	})

	bdd.AfterEach(func() {
		logger.Assert(t(), "")
	})

	bdd.Context("JsonEqual", func() {

		assertSucceed := func(expected, actual string) {
			assert.True(t(), JsonEqual(logger, expected, actual))
			logger.Assert(t(), "")
		}

		assertFailed := func(expected, actual string) {
			assert.False(t(), JsonEqual(logger, expected, actual))
			logger.Assert(t(), "Json string not equal, expected:.*")
		}

		bdd.It("Bad json string", func() {
			assert.False(t(), JsonEqual(logger, "bad json", "bad json"))
			logger.Assert(t(), `^Invalid json string: "bad json".*`)
		})

		bdd.It("null", func() {
			assertSucceed("null", "null")
			assertFailed("null", "1")
		})

		bdd.It("number", func() {
			assertSucceed("1", "1")
			assertFailed("1", "2")
			assertFailed("1", "null")
		})

		bdd.It("bool", func() {
			assertSucceed("true", "true")
			assertFailed("true", "false")
			assertFailed("true", "1")
		})

		bdd.It("string", func() {
			assertSucceed(`""`, `""`)
			assertFailed(`"foo"`, `"bar"`)
			assertFailed(`"foo"`, `null`)
		})

		bdd.It("List", func() {
			assertSucceed("[]", "[]")
			assertSucceed("[1, 2]", "[1, 2]")
			assertFailed("[]", "1")
			assertFailed("[1]", "[2]")
			assertFailed("[1,2]", "[2,1]")
		})

		bdd.It("Object", func() {
			assertSucceed("{}", "{}")
			assertSucceed(`{"a": 1}`, `{"a": 1}`)
			assertSucceed(`{"a": 1, "b": 2}`, `{"b": 2, "a": 1}`)
			assertSucceed(`{"a": 1, "b": {"c": 2, "d": {}}}`, `{"b": {"d": {}, "c": 2}, "a": 1}`)
			assertFailed("{}", "1")
			assertFailed(`{"a":1}`, `{"A": 2}`)
		})

	})

	bdd.Context("Equivalent", func() {

		bdd.It("Success", func() {
			for i, v := range []struct {
				X, Y interface{}
			}{
				{[]int{}, []int{}},
				{[]int{}, ([]int)(nil)},
				{[]int{1, 2, 3}, []int{3, 2, 1}},
				{[]int{1, 1, 3}, []int{1, 3, 1}},
			} {
				assert.True(t(), Equivalent(logger, v.X, v.Y), "%d", i)
			}
		})

		bdd.It("Failed", func() {
			for i, v := range []struct {
				X, Y interface{}
			}{
				{[]int{}, []int32{}},
				{[]int{1}, []int32{}},
				{[]int{1, 2}, []int32{2, 3}},
				{[]int{}, []int32{2}},
				{[]int{1, 1, 3}, []int{1, 3, 2}},
				{1, []int{1}},
				{[]int{1}, 1},
			} {
				assert.False(t(), Equivalent(logger, v.X, v.Y), "%d", i)
				logger.Assert(t(), ".*not equivalent to.*")
			}
		})

		// TODO: argument is not slice

	})

	bdd.Context("Panics", func() {
		var hit int

		bdd.BeforeEach(func() {
			hit = 0
		})

		bdd.AfterEach(func() {
			assert.Equal(t(), 1, hit)
		})

		bdd.It("success", func() {
			assert.True(t(), Panics(logger, func() {
				hit++
				panic("abc")
			}, "abc"))
		})

		bdd.It("failed because not paniced", func() {
			assert.False(t(), Panics(logger, func() {
				hit++
			}, "abc"))
			logger.Assert(t(), `Expected panic with "abc", but not panicked`)
		})

		bdd.It("Mismatched error object 1", func() {
			assert.False(t(), Panics(logger, func() {
				hit++
				panic("cde")
			}, "abc"))
			logger.Assert(t(), `Expected panic with "abc", but got "cde"`)
		})

		bdd.It("Mismatched error object 2", func() {
			assert.False(t(), Panics(logger, func() {
				hit++
				panic(2)
			}, "abc"))
			logger.Assert(t(), `Expected panic with "abc", but got 2`)
		})

		bdd.It("Match error object", func() {
			assert.True(t(), Panics(logger, func() {
				hit++
				panic(errors.New("foo"))
			}, "foo"))
		})

	})

})
