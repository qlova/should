package check_test

import (
	"image"
	"os"
	"testing"

	"qlova.org/should"
	"qlova.org/should/check"
)

func OpenImage(name string) (img image.Image, err error) {
	defer check.Return(&err)

	f, err := os.Open(name)
	check.Errorf("failed to open image: %w", err)

	img, _, err = image.Decode(f)
	check.Errorf("failed to decode image: %w", err)

	return
}

func OpenImage2(name string) (img image.Image, err error) {
	defer check.Returnf("error opening image: %w", &err)

	f, err := os.Open(name)
	check.Error(err)

	img, _, err = image.Decode(f)
	check.Error(err)

	return
}

func Test_Check(t *testing.T) {
	should.Error(
		OpenImage("test.jpg"),
	).Test(t)

	should.Error(
		OpenImage2("test.jpg"),
	).Test(t)
}
