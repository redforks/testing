package matcher

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

type panics struct {
	err           interface{}
	actPanicValue interface{}
}

// Panics return a gomega matcher, the actual value is a function
// without any argument and returns nothing. Panics() expect function
// panics with expected value.
//
// Compare the panic value likes gomega.Equal(), also accept gomega Matcher
// interface.
func Panics(expected interface{}) types.GomegaMatcher {
	return &panics{err: expected}
}

func (p *panics) Match(actual interface{}) (success bool, err error) {
	v := reflect.ValueOf(actual)

	if v.Kind() != reflect.Func {
		return false, fmt.Errorf("Panics expects a function.  Got:\n%s", format.Object(actual, 1))
	}

	t := v.Type()
	if !(t.NumIn() == 0 && t.NumOut() == 0) {
		return false, fmt.Errorf("Panics expects a function with no arguments and no return value.  Got:\n%s", format.Object(actual, 1))
	}

	defer func() {
		if e := recover(); e != nil {
			p.actPanicValue = e
			if m, ok := p.err.(types.GomegaMatcher); ok {
				success, err = m.Match(e)
			} else {
				success = reflect.DeepEqual(e, p.err)
			}
		}
	}()

	v.Call([]reflect.Value{})

	return
}

func (p *panics) FailureMessage(actual interface{}) (message string) {
	if p.actPanicValue == nil {
		return format.Message(actual, "to panic with:", p.err)
	}

	return fmt.Sprintf(
		"Expected\n%s\nto panic with:\n%s\nbut got:\n%s",
		format.Object(actual, 1),
		format.Object(p.err, 1),
		format.Object(p.actPanicValue, 1),
	)
}

func (p *panics) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to panic, or not to panic with:", p.err)
}

// PanicsWithSubstring returns a gomega matcher, like Panics() but matches substring of err.Error().
func PanicsWithSubstring(substr string) types.GomegaMatcher {
	return &panicsWithSubstring{
		substr: substr,
		panics: panics{err: gomega.WithTransform(func(v interface{}) string {
			return fmt.Sprint(v)
		}, gomega.ContainSubstring(substr)),
		},
	}
}

type panicsWithSubstring struct {
	panics

	substr string
}

func (p *panicsWithSubstring) FailureMessage(actual interface{}) string {
	if p.panics.actPanicValue == nil {
		return format.Message(actual, "to panic contains:", p.substr)
	}

	return fmt.Sprintf(
		"Expected\n%s\nto panic contains:\n%s\nbut got:\n%s",
		format.Object(actual, 1),
		format.Object(p.substr, 1),
		format.Object(p.actPanicValue, 1),
	)
}

func (p *panicsWithSubstring) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not paniced, or not panic contains:", p.substr)
}
