package utils

import (
	"reflect"
	"testing"
)

type TestCase[T any, U any] struct {
	Case     T
	Expected U
}

func assertComparable[T comparable](t testing.TB, result, expect T, errorMessage string) {
	t.Helper()
	if result != expect {
		t.Errorf(errorMessage, expect, result)
	}
}

func AssertDeepEqual[T any](t testing.TB, result, expect T) {
	t.Helper()
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("expect %+v, got %+v", expect, result)
	}
}

func AssertInt(t testing.TB, result, expect int) {
	t.Helper()
	assertComparable(t, result, expect, "expected %d, got %d")
}

func AssertString(t testing.TB, result, expect string) {
	t.Helper()
	assertComparable(t, result, expect, "expected %s, got %s")
}

func AssertBool(t testing.TB, result, expect string) {
	t.Helper()
	assertComparable(t, result, expect, "expected %v, got %v")
}
