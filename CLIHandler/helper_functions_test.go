package CLIHandler

import (
	"errors"
	"reflect"
	"testing"
)

var boink = errors.New("boink")

func AssertEquals(t *testing.T, actual, expected any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

func AssertError(t *testing.T, actual, expected error) {
	if !errors.Is(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

type WriterError struct {
	err error
}

func (this *WriterError) Write(p []byte) (n int, err error) {
	return 0, this.err
}