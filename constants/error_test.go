package constants

import (
	"testing"
)

func TestError(t *testing.T) {
	errors := []error{ErrRequiredUUID, ErrInvalidUUID}

	for _, err := range errors {
		if err == nil {
			t.Errorf("Error should not be nil")
		} else if _, ok := err.(error); !ok {
			t.Errorf("%v is not an error", err)
		} else if s, ok := err.(interface{ Error() string }); ok {
			if len(s.Error()) == 0 {
				t.Errorf("%v is empty", err)
			}
		}
	}
}
