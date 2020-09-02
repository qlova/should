//Package errors is a drop-in replacement for the stdlib version except it supports constant errors and basic tracing.
package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

//Error implements error
type Error string

//New creates a new Error.
type New = Error

func (err Error) Error() string {
	return string(err)
}

//Trace prefixes a basic trace to the given error.
//Uses function name & current line-number.
//ie DoSomething:3
func Trace(err error) (out error) {
	if err == nil {
		return nil
	}
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()

		var splits = strings.Split(frame.Function, ".")

		var name = splits[len(splits)-1]

		out = fmt.Errorf("%v:%v %w", name, line, err)
	}
	return
}

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
//
// An error type might provide an Is method so it can be treated as equivalent
// to an existing error. For example, if MyError defines
//
//	func (m MyError) Is(target error) bool { return target == os.ErrExist }
//
// then Is(MyError{}, os.ErrExist) returns true. See syscall.Errno.Is for
// an example in the standard library.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
