package test

import (
	"io/ioutil"
	"testing"
	"time"
)

//State is an interface of *testing.T
type State interface {
	Cleanup(f func())
	Deadline() (deadline time.Time, ok bool)
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Parallel()
	Run(name string, f func(t *testing.T)) bool
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
	TempDir() string
}

//Discard returns a state that discards all operations.
func Discard() State {
	return discard{}
}

type discard struct{}

func (d discard) Cleanup(f func()) {}

func (d discard) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (d discard) Error(args ...interface{}) {}

func (d discard) Errorf(format string, args ...interface{}) {}

func (d discard) Fail() {}

func (d discard) FailNow() {}

func (d discard) Failed() bool {
	return false
}

func (d discard) Fatal(args ...interface{}) {}

func (d discard) Fatalf(format string, args ...interface{}) {}

func (d discard) Helper() {}

func (d discard) Log(args ...interface{}) {}

func (d discard) Logf(format string, args ...interface{}) {}

func (d discard) Name() string {
	return ""
}

func (d discard) Parallel() {}

func (d discard) Run(name string, f func(t *testing.T)) bool {
	return true
}

func (d discard) Skip(args ...interface{}) {}

func (d discard) SkipNow() {}

func (d discard) Skipf(format string, args ...interface{}) {}

func (d discard) Skipped() bool {
	return false
}

func (d discard) TempDir() string {
	s, _ := ioutil.TempDir("", "")
	return s
}
