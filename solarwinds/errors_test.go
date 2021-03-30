package solarwinds

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientErrors(t *testing.T) {
	user := RandString(10) + "@nordcloud.com"
	errs := []error{
		NewErrorAttemptDeleteActiveUser(user),
		NewNetworkError(errors.New("underlying network error")),
	}
	expectedErrMsg := []string{
		fmt.Sprintf("status: %d, err: deleting active user %v is not supported", AttemptDeleteActiveUser, user),
		fmt.Sprintf("status: %d, err: underlying network error", NetworkError),
	}
	for i, err := range errs {
		if err != nil {
			errMsg := err.Error()
			assert.Equal(t, errMsg, expectedErrMsg[i])
		} else {
			t.Error("should not reach here")
		}
	}
}
