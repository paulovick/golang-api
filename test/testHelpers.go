package testHelpers

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	debug.PrintStack()
	t.Errorf("%s | Received %v (type %v), expected %v (type %v)", message, a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}