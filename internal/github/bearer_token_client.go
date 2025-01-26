package github

import (
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

func NewBearerTokenClient(token string) *http.Client {
	return &http.Client{
		Timeout:   defaultTimeout,
		Transport: newBearerTokenTransport(token, http.DefaultTransport),
	}
}

// newBearerTokenTransport returns a custom Transport that adds the Authorization header.
func newBearerTokenTransport(token string, baseTransport http.RoundTripper) http.RoundTripper {
	return &customTransport{
		token:         token,
		baseTransport: baseTransport,
	}
}

// customTransport is a custom HTTP RoundTripper that adds the Bearer token.
type customTransport struct {
	token         string
	baseTransport http.RoundTripper
}

// RoundTrip adds the Authorization header to each request.
func (c *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid modifying the original
	reqClone := req.Clone(req.Context())
	reqClone.Header.Set("Authorization", "Bearer "+c.token)

	// Use the base transport to execute the request
	return c.baseTransport.RoundTrip(reqClone)
}
