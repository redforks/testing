// tassert package contains supplement assert funcs to testify assert
package tassert

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
)

// Compare two json string, assert if not equal, return true if they are equal.
// Json string can not compare literally, because space and order of json object properties.
func JsonEqual(t assert.TestingT, expected, actual string) bool {
	var expJson, actJson interface{}

	if err := json.Unmarshal(([]byte)(expected), &expJson); err != nil {
		t.Errorf("Invalid json string: %#v\n %s", expected, err)
		return false
	}

	if err := json.Unmarshal(([]byte)(actual), &actJson); err != nil {
		t.Errorf("Invalid json string: %#v\n %s", actual, err)
		return false
	}

	if !assert.ObjectsAreEqual(expJson, actJson) {
		sExp, _ := json.Marshal(expJson)
		sAct, _ := json.Marshal(actJson)
		t.Errorf("Json string not equal, expected:\n%#v, actual:\n%#v", (string)(sExp), (string)(sAct))
		return false
	}
	return true
}

// Compare two slices, return true if they are equal, but its order may different.
// Return true if they are equivalent.
func Equivalent(t assert.TestingT, expected, actual interface{}) bool {
	if !equivalent(expected, actual) {
		t.Errorf("%#v\nnot equivalent to:\n%#v", expected, actual)
		return false
	}
	return true
}

// Like testify assert.Panics(), but also compare panic object.
// Return true if fn raised panicked with expected value.
// If the function panicked with error object, then compare expected with error.Error().
func Panics(t assert.TestingT, fn assert.PanicTestFunc, expected interface{}) (ok bool) {
	defer func() {
		e := recover()
		ok = e != nil
		if !ok {
			t.Errorf("Expected panic with %#v, but not panicked", expected)
			return
		}

		if err, ok := e.(error); ok {
			e = err.Error()
		}

		if !assert.ObjectsAreEqual(expected, e) {
			t.Errorf("Expected panic with %#v, but got %#v", expected, e)
			ok = false
		}
	}()

	fn()
	return true
}

func equivalent(expected, actual interface{}) bool {
	vExp, vAct := reflect.ValueOf(expected), reflect.ValueOf(actual)
	if vExp.Type().Kind() != reflect.Slice || vAct.Type().Kind() != reflect.Slice {
		return false
	}

	if vExp.Type().Elem() != vAct.Type().Elem() {
		return false
	}

	if vExp.Len() != vAct.Len() {
		return false
	}

	return allContains(vExp, vAct) && allContains(vAct, vExp)
}

// return true if every item in slice y exist in slice x
func allContains(x, y reflect.Value) bool {
	for i := 0; i < x.Len(); i++ {
		if !contains(x.Index(i).Interface(), y) {
			return false
		}
	}
	return true
}

// return true if item exist in slice.
func contains(item interface{}, slice reflect.Value) bool {
	for i := 0; i < slice.Len(); i++ {
		if assert.ObjectsAreEqual(item, slice.Index(i).Interface()) {
			return true
		}
	}
	return false
}
