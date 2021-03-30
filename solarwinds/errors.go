package solarwinds

import "fmt"

const (
	NetworkError uint32 = iota
	AttemptDeleteActiveUser
)

type ClientError struct {
	StatusCode uint32 `json:"statusCode"`
	Err        error  `json:"err"`
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("status: %d, err: %v", c.StatusCode, c.Err)
}

func NewNetworkError(cause error) error {
	return &ClientError{
		StatusCode: NetworkError,
		Err:        cause,
	}
}

func NewErrorAttemptDeleteActiveUser(user string) error {
	return &ClientError{
		StatusCode: AttemptDeleteActiveUser,
		Err:        fmt.Errorf("deleting active user %v is not supported", user),
	}
}
