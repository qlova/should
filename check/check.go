//Package check is a error-checking package for convient ways to Go errors to the calling function.
package check

import "fmt"

//Return should be deferred in functions that pass errors to their caller, pass the named return value that errors should be passed to.
func Return(into *error) {
	if err := recover().(error); err != nil {
		*into = err
	}
}

//Returnf should be deferred in functions that pass errors to their caller, pass the named return value that errors should be passed to.
//The format string should be a constant.
func Returnf(format string, into *error) {
	if err := recover().(error); err != nil {
		*into = fmt.Errorf(format, err)
	}
}

//Error should be called in functions that pass errors to their caller, pass the error to check.
func Error(err error) {
	if err != nil {
		panic(err)
	}
}

//Errorf should be called in functions that pass errors to their caller, pass the error to check and a format string for the error.
//The format string should be a constant.
func Errorf(format string, err error) {
	if err != nil {
		panic(fmt.Errorf(format, err))
	}
}
