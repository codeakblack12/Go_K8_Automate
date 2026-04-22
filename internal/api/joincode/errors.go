package joincode

import "fmt"

type HTTPError struct {
	StatusCode int
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("join-code service returned status %d: %s", e.StatusCode, e.Body)
}
