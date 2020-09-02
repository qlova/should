package errors_test

import (
	"testing"

	"qlova.org/should"
	"qlova.org/should/check/errors"
)

func Test_Errors(t *testing.T) {
	should.Error(
		errors.Trace(errors.New("Hello")),
	).Test(t)

	should.NotError(
		errors.Trace(nil),
	).Test(t)

	should.Panic(func() {
		errors.As(nil, nil)
	}).Test(t)

	should.NotPanic(func() {
		errors.Is(nil, nil)
	}).Test(t)

	should.NotPanic(func() {
		errors.Unwrap(nil)
	}).Test(t)
}
