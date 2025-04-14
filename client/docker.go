package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

// ProxyDockerRequest proxies a request to the Docker API via the Portainer API.
// The function constructs a request to the Portainer server's Docker API proxy endpoint,
// authenticates using the configured API key, and forwards the request to the target Docker environment.
//
// Parameters:
//   - environmentId: The ID of the target Docker environment in Portainer
//   - method: The HTTP method to use (GET, POST, PUT, DELETE, etc.)
//   - dockerAPIPath: The Docker API endpoint path to proxy to (e.g., "/containers/json"). Must include the leading slash.
//   - queryParams: A map of query parameters to include in the request URL (can be nil)
//   - body: The request body to send (can be nil for requests that don't have a body)
//
// Returns:
//   - *http.Response: The response from the Docker API
//   - error: Any error that occurred during the request
func (c *PortainerClient) ProxyDockerRequest(environmentId int, method string, dockerAPIPath string, queryParams map[string]string, body io.Reader) (*http.Response, error) {
	baseURL := fmt.Sprintf("https://%s/api/endpoints/%d/docker%s", c.proxyCfg.Host, environmentId, dockerAPIPath)

	req, err := http.NewRequest(method, baseURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker proxy request: %w", err)
	}

	// Add query parameters if provided
	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range queryParams {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("x-api-key", c.proxyCfg.Token)

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
