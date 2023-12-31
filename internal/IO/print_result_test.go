package IO

import (
	"bytes"
	"errors"
	"testing"
)

func TestPrintPBIResult(t *testing.T) {
	testFunction := func(t testing.TB, url, want string, err error) {
		t.Helper()
		buffer := &bytes.Buffer{}
		PrintGeneratedUrl(buffer, url, err)
		got := buffer.String()

		if got != want {
			t.Errorf("expected %s but got %s", want, got)
		}
	}

	t.Run("Validate print with success the URL", func(t *testing.T) {
		testFunction(t, "http://blablabla", successMessagePrefix+"http://blablabla\n", nil)
	})

	t.Run("Should print an error message when err is not nil", func(t *testing.T) {
		testFunction(t, "", errorMessagePrefix+"Whatever\n", errors.New("Whatever"))
	})
}
