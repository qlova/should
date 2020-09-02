package should_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"qlova.org/should"
)

func Test_Panic(t *testing.T) {
	should.Panic(func() {
		panic(nil)
	}).Test(t)

	should.Error(should.Panic(func() {})).Test(t)
}

func Test_NotPanic(t *testing.T) {
	should.NotPanic(
		func() {},
	).Test(t)

	should.Error(
		should.NotPanic(nil),
	).Test(t)
}

func Test_Error(t *testing.T) {

	var e interface{} = nil

	fmt.Println(reflect.TypeOf(e) == nil)

	should.Error(
		errors.New(""),
	).Test(t)

	should.Error(
		should.Error(error(nil)),
	).Test(t)
}

func Test_NotError(t *testing.T) {
	should.NotError(
		nil,
	).Test(t)

	should.Error(
		should.NotError(errors.New("")),
	).Test(t)
}

func Test_Be(t *testing.T) {
	should.Be(3)(3).Test(t)

	should.Error(
		should.Be(3)(1),
	).Test(t)
}
