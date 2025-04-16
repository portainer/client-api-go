package client

import (
	"fmt"
	"io"
	"net/http"
)

// ProxyRequestOptions bundles the parameters for an API (Docker, Kubernetes, etc.) proxy request.
type ProxyRequestOptions struct {
	// Method is the HTTP method to use (GET, POST, PUT, DELETE, etc.)
	Method string
	// APIPath is the API endpoint path to proxy to (e.g., "/containers/json" or "/api/v1/namespaces/default/pods"). Must include the leading slash.
	APIPath string
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
	baseURL := fmt.Sprintf("https://%s/api/endpoints/%d/docker%s", c.proxyCli.host, environmentId, opts.APIPath)
	return c.proxyRequest(baseURL, opts)
}

// ProxyKubernetesRequest proxies a request to the Kubernetes API via the Portainer API
// using the provided options.
//
// Parameters:
//   - environmentId: The ID of the target Kubernetes environment in Portainer
//   - opts: Options defining the proxied request (method, path, query params, headers, body)
//
// Returns:
//   - *http.Response: The response from the Kubernetes API
//   - error: Any error that occurred during the request
func (c *PortainerClient) ProxyKubernetesRequest(environmentId int, opts ProxyRequestOptions) (*http.Response, error) {
	baseURL := fmt.Sprintf("https://%s/api/endpoints/%d/kubernetes%s", c.proxyCli.host, environmentId, opts.APIPath)
	return c.proxyRequest(baseURL, opts)
}

func (c *PortainerClient) proxyRequest(baseURL string, opts ProxyRequestOptions) (*http.Response, error) {
	req, err := http.NewRequest(opts.Method, baseURL, opts.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create proxy request: %w", err)
	}

	// Add query parameters if provided
	if opts.QueryParams != nil {
		q := req.URL.Query()
		for k, v := range opts.QueryParams {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("x-api-key", c.proxyCli.token)

	// Add custom headers if provided
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	resp, err := c.proxyCli.cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send proxy request: %w", err)
	}

	return resp, nil
}
