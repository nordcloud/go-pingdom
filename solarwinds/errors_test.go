package solarwinds

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientErrors(t *testing.T) {
	user := RandString(10) + "@nordcloud.com"
	err := NewErrorAttemptDeleteActiveUser(user)
	if err != nil {
		errMsg := err.Error()
		assert.Equal(
			t, errMsg,
			fmt.Sprintf("status: %d, err: deleting active user %v is not supported", AttemptDeleteActiveUser, user),
		)
	} else {
		t.Error("should not reach here")
	}
}
