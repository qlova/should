package test_test

import (
	"fmt"
	"testing"

	"qlova.org/should"
	"qlova.org/should/check/errors"
	"qlova.org/should/test"
)

type TestSuite struct {
	test.Suite
}

func (t TestSuite) SetupSuite() {}

func (t TestSuite) Test_Error() {
	should.Error(
		errors.New("hello"),
	).Test(t.T())

	should.NotError(
		nil,
	).Test(t.T())

	should.Panic(func() {
		test.New(TestSuite{})
	}).Test(t.T())

	should.Panic(func() {
		fmt.Println(test.Error{nil}.Error())
	}).Test(t.T())

	//Test-test TEST! utililties.
	var d = test.Discard()

	should.Panic(func() {

		d.Cleanup(nil)
		d.Error()
		d.Errorf("")
		d.Fail()
		d.FailNow()
		d.Fatal()
		d.Fatalf("")
		d.Helper()
		d.Log()
		d.Logf("")
		d.Parallel()
		d.Skip()
		d.SkipNow()
		d.Skipf("")

		d.Deadline()
		d.Failed()
		d.Name()
		d.Skipped()
		d.TempDir()
		d.Run("", nil)
	}).Test(d)
}

func Test_Suite(t *testing.T) {
	test.New(new(TestSuite))(t)
}
