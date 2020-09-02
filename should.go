//Package should provides useful utilities for testing and verification.
package should

import (
	"errors"
	"fmt"

	"qlova.org/should/test"
)

//Panic returns an error if the provided function does not panic.
func Panic(f func()) (err test.Error) {
	panicked := true
	defer func() {
		recover()
		if panicked {
			err = nil
		}
	}()
	err = test.Error{errors.New("should have panicked")}
	f()
	panicked = false
	return
}

//NotPanic returns an error if the provided function panics.
func NotPanic(f func()) (err test.Error) {
	panicked := true
	defer func() {
		p := recover()
		if panicked {
			err = test.Error{fmt.Errorf("should not have panicked: %v", p)}
		}
	}()
	f()
	panicked = false
	return
}

//Error returns an error if its last argument is nil.
func Error(args ...interface{}) test.Error {
	if args[len(args)-1] == nil {
		return test.Error{errors.New("should have been an error")}
	}
	return nil
}

//NotError returns an error if its last argument is not nil.
func NotError(args ...interface{}) test.Error {
	if args[len(args)-1] != nil {
		return test.Error{errors.New("should not have been an error")}
	}
	return nil
}

//Be returns an error if the first varadic argument is not equal to the given value.
func Be(value interface{}) func(args ...interface{}) test.Error {
	return func(args ...interface{}) test.Error {
		if value != args[0] {
			return test.Error{fmt.Errorf("should be %v but found %v", value, args[0])}
		}
		return nil
	}
}
