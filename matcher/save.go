package matcher

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

type save struct {
	p interface{} // pointer to receive the result value
}

func Save(ptr interface{}) types.GomegaMatcher {
	return save{ptr}
}

func (s save) Match(actual interface{}) (bool, error) {
	argType := reflect.TypeOf(s.p)
	if argType.Kind() != reflect.Ptr {
		return false, fmt.Errorf("Cannot save a value From:\n%s\nTo:\n%s\nYou need to pass a pointer!", format.Object(actual, 1), format.Object(s.p, 1))
	}

	actualType := reflect.TypeOf(actual)
	assignable := actualType.AssignableTo(argType.Elem())
	if !assignable {
		return false, fmt.Errorf("Cannot assign a value from:\n%s\nTo:\n%s", format.Object(actual, 1), format.Object(s.p, 1))
	}

	outValue := reflect.ValueOf(s.p)
	reflect.Indirect(outValue).Set(reflect.ValueOf(actual))
	return true, nil
}

func (s save) FailureMessage(actual interface{}) string {
	return ""
}

func (s save) NegatedFailureMessage(actual interface{}) string {
	return ""
}
