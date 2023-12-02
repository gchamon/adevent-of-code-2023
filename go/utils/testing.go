package utils

type TestCase[T any, U any] struct {
	Case     T
	Expected U
}
