# YOU SHOULD CHECK ERRORS AND TEST!

![Report Card](https://goreportcard.com/badge/github.com/qlova/should)

This module provides error-checking and testing utilities that **YOU SHOULD USE**.

```
package should_test

import (
	"strconv"
	"testing"

	//You should check errors and test!
	"qlova.org/should"
	"qlova.org/should/check"
	"qlova.org/should/check/errors"
	"qlova.org/should/test"
)

type Example struct {
	test.Suite

	Value string
}

func (e *Example) SetupSuite() {
	e.Value = "Hello World"
}

func (e Example) TestValue() {
	should.Be("Hello World")(e.Value).Test(e.T())
}

func (e Example) TestPanic() {
	should.Panic(func() {
		panic("Hello World")
	}).Test(e.T())
}

func (e Example) TestErrors() {

	AddStringIntegers := func(a, b string) (sum int, err error) {
		defer check.Returnf("tostring failed: %w", &err)

		i1, err := strconv.Atoi(a)
		check.Error(errors.Trace(err))

		i2, err := strconv.Atoi(b)
		check.Error(errors.Trace(err))

		return i1 + i2, nil
	}

	should.Be(5)(AddStringIntegers("3", "2")).Test(e.T())

	should.Error(AddStringIntegers("asdsadsad", "2")).Test(e.T())
}

func TestExample(t *testing.T) {
	test.New(new(Example))(t)
}
```

## FAQ:
* _So check.Error panics and returns the error to the caller? That's implicit control flow!_  
Yes, you should check errors.

* _Why is there a replacement for the Go errors package?_  
Go is opioninated, Qlova is opinionated too. Either way, you should check errors.  

* _How does this rate to other Go testing frameworks?_  
First of all, this is not a framework, it is a lightweight module. Additionally, the packages in this module are not restricted to 'testing' contexts and you can use the 'should' package for general assertion purposes, validation, checking errors etc.
