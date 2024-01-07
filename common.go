package oblio

import (
	"errors"
	"net/http"
)

type Status struct {
	Status        int    `json:"status,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

type Authorized struct {
	AccessToken string `json:"-" url:"-"`
}

func (a Authorized) GetAccessToken() string {
	return a.AccessToken
}

func IsUnauthorizedError(err error) bool {
	if err == nil {
		return false
	}

	var errResp *ErrorResponse

	return errors.As(err, &errResp) && errResp.Status == http.StatusUnauthorized
}
