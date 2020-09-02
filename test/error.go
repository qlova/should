//Package test provides testable errors.
package test

//Error for this package.
type Error []error

func (err Error) Error() string {
	return err[0].Error()
}

//Test tests if there is an error and if there is, it fatally fails the current test.
func (err Error) Test(t State) {
	t.Helper()

	if err != nil {
		t.Fatalf("%v", err[0])
	}
}
