package client

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/portainer/client-api-go/v2/pkg/client"
)

// PortainerClient provides a simplified interface to the Portainer API client
// that uses API token authentication for all operations
type PortainerClient struct {
	cli      *client.PortainerClientAPI
	proxyCfg ProxyConfiguration
}

// ProxyConfiguration contains the configuration used to proxy requests via the Portainer API
type ProxyConfiguration struct {
	// Host is the host of the Portainer server
	// It includes the hostname and port
	Host string
	// Token is the API token used to authenticate to the Portainer server
	Token string
	// SkipTLSVerify is used to disable TLS verification for the proxy connection
	SkipTLSVerify bool
}

// ClientOption defines a functional option for configuring the Portainer client
type ClientOption func(*clientOptions)

// clientOptions holds all configuration for the Portainer client
type clientOptions struct {
	host          string
	basePath      string
	scheme        string
	apiKey        string
	skipTLSVerify bool
}

// WithBasePath sets the base path for the Portainer client
func WithBasePath(basePath string) ClientOption {
	return func(o *clientOptions) {
		o.basePath = basePath
	}
}

// WithScheme sets the scheme (http/https) for the Portainer client
func WithScheme(scheme string) ClientOption {
	return func(o *clientOptions) {
		o.scheme = scheme
	}
}

// WithSkipTLSVerify enables or disables TLS verification
func WithSkipTLSVerify(skip bool) ClientOption {
	return func(o *clientOptions) {
		o.skipTLSVerify = skip
	}
}

// NewPortainerClient creates a new Portainer client with required parameters and optional configuration
// Host and apiKey are required, while other settings can be customized through options
func NewPortainerClient(host, apiKey string, opts ...ClientOption) *PortainerClient {
	options := &clientOptions{
		host:          host,
		apiKey:        apiKey,
		basePath:      "/api",
		scheme:        "https",
		skipTLSVerify: false,
	}

	for _, opt := range opts {
		opt(options)
	}

	transport := httptransport.New(options.host, options.basePath, []string{options.scheme})

	// Configure TLS if needed
	if options.skipTLSVerify {
		transport.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	// Configure API key authentication
	apiKeyAuth := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("x-api-key", options.apiKey)
	})
	transport.DefaultAuthentication = apiKeyAuth

	return &PortainerClient{
		cli: client.New(transport, nil),
		proxyCfg: ProxyConfiguration{
			Host:          options.host,
			Token:         options.apiKey,
			SkipTLSVerify: options.skipTLSVerify,
		},
	}
}
