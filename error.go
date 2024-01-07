package oblio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
)

type ErrorResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"statusMessage,omitempty"`
}

func UnmarshalErrorResponse(resp *http.Response) *ErrorResponse {
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	out := &ErrorResponse{}

	if err := json.Unmarshal(body, out); err != nil {
		out.Message = string(body)
	}

	if out.Status == 0 {
		out.Status = resp.StatusCode
	}

	if out.Message == "" {
		out.Message = string(body)
	}

	return out
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("status code: %d, message: %s", e.Status, e.Message)
}
