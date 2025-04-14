package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

// ProxyRequestOptions bundles the parameters for a Docker proxy request.
type ProxyRequestOptions struct {
	// Method is the HTTP method to use (GET, POST, PUT, DELETE, etc.)
	Method string
	// DockerAPIPath is the Docker API endpoint path to proxy to (e.g., "/containers/json"). Must include the leading slash.
	DockerAPIPath string
	// QueryParams is a map of query parameters to include in the request URL (can be nil)
	QueryParams map[string]string
	// Headers is a map of headers to include in the request (can be nil)
	Headers map[string]string
	// Body is the request body to send (can be nil for requests that don't have a body)
	Body io.Reader
}

// ProxyDockerRequest proxies a request to the Docker API via the Portainer API
// using the provided options.
//
// Parameters:
//   - environmentId: The ID of the target Docker environment in Portainer
//   - opts: Options defining the proxied request (method, path, query params, headers, body)
//
// Returns:
//   - *http.Response: The response from the Docker API
//   - error: Any error that occurred during the request
func (c *PortainerClient) ProxyDockerRequest(environmentId int, opts ProxyRequestOptions) (*http.Response, error) {
	baseURL := fmt.Sprintf("https://%s/api/endpoints/%d/docker%s", c.proxyCfg.Host, environmentId, opts.DockerAPIPath)

	req, err := http.NewRequest(opts.Method, baseURL, opts.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker proxy request: %w", err)
	}

	// Add query parameters if provided
	if opts.QueryParams != nil {
		q := req.URL.Query()
		for k, v := range opts.QueryParams {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("x-api-key", c.proxyCfg.Token)

	// Add custom headers if provided
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.proxyCfg.SkipTLSVerify,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send Docker proxy request: %w", err)
	}

	return resp, nil
}
