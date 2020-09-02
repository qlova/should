package test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

//Suite can be embedded inside of a type to turn it into a testsuite.
//Any methods beginning with Test are run when testing.
type Suite interface {
	T() *testing.T
	SetT(*testing.T)
}

//Test is a testing function.
type Test func(*testing.T)

//New allows a test to be created from a Suite.
func New(testsuite Suite) Test {
	var value = reflect.ValueOf(testsuite)

	if value.Kind() != reflect.Ptr {
		panic(fmt.Errorf("%v value must be passed by reference to test.New", value.Type()))
	}

	value = value.Elem()

	field, ok := value.Type().FieldByName("Suite")
	if ok && field.Type == reflect.TypeOf([0]Suite{}).Elem() && field.Anonymous {
		value.FieldByName("Suite").Set(reflect.ValueOf(new(suite)))
	}

	return func(t *testing.T) {
		testsuite.SetT(t)

		if _, ok := value.Type().MethodByName("SetupSuite"); ok {
			value.MethodByName("SetupSuite").Call(nil)
		}

		for i := 0; i < value.NumMethod(); i++ {
			method := value.Method(i)

			if strings.HasPrefix(value.Type().Method(i).Name, "Test") {
				method.Call(nil)
			}
		}
	}
}

type suite struct {
	t *testing.T
}

func (s suite) T() *testing.T {
	return s.t
}

func (s *suite) SetT(t *testing.T) {
	s.t = t
}
