package utils_test

import (
	"01_car_steps/utils"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to see a panic")
		}
	}()
	filePanicFcn()
}

func TestPanic2(t *testing.T) {
	assertPanic(t, filePanicFcn)
}

func TestPanic3(t *testing.T) {
	assertPanic(t, zeroPanicFcn)
}

// -------------
//    Helpers
// -------------

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to see a panic")
		}
	}()
	f()
}

func filePanicFcn() {
	_, err := os.Open("doesnot-exist.txt")
	if err != nil {
		utils.HandlePanic(errors.Wrap(err, "unable to read file"))
	}
}

func divFcn(d int) error {
	if d == 0 {
		return errors.New("divide by 0 attempted")
	}
	return nil
}

func zeroPanicFcn() {
	err := divFcn(0)
	if err != nil {
		utils.HandlePanic(err)
	}
}
