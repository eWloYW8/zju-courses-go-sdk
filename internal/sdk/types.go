package sdk

import (
	"fmt"
	"net/http"
)

type Response struct {
	*http.Response
}

type APIError struct {
	StatusCode int
	Message    string
	URL        string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s (URL: %s)", e.StatusCode, e.Message, e.URL)
}
