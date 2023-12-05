package utils

import (
	"errors"
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

func AssertRune(t testing.TB, result, expect rune) {
	t.Helper()
	assertComparable(t, string(result), string(expect), "expected rune('%s'), got rune('%s')")
}

func AssertString(t testing.TB, result, expect string) {
	t.Helper()
	assertComparable(t, result, expect, "expected '%s', got '%s'")
}

func AssertBool(t testing.TB, result, expect bool) {
	t.Helper()
	assertComparable(t, result, expect, "expected %t, got %t")
}

func AssertExpectError(t testing.TB, result, expect error) {
	t.Helper()
	if !errors.Is(result, expect) {
		t.Errorf("expected error '%v', got '%v'", expect, result)
	}
}

func AssertNotError(t testing.TB, result error) {
	t.Helper()
	if result != nil {
		t.Errorf("got error when was not expecting one: %s", result)
	}
}
