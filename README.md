# Portainer Client GO SDK

Swagger-generated client SDK in Golang for Portainer.

## Installation

Use the following command to install the latest version of the SDK:

```sh
go get -u github.com/portainer/client-api-go/v2
```

To install a specific version of the SDK (to target a specific Portainer server version):

```sh
go get -u github.com/portainer/client-api-go/v2@VERSION
```

For example, to install version 2.27.1:

```sh
go get -u github.com/portainer/client-api-go/v2@v2.27.1
```

Available versions can be found at: https://github.com/portainer/client-api-go/tags

## Using the Client SDK

There are two ways to use the SDK:

### 1. Using the Simple Client (Recommended)

The simple client provides an easy-to-use interface to interact with the Portainer API. All the Portainer API operations are not supported yet.

```go
import client "github.com/portainer/client-api-go/v2/client"

// Initialize client with API key
cli := client.NewPortainerClient(
	"portainer.dev.local",   // Portainer host
	"ptr_XXXYYYZZZ",         // Portainer API key
	client.WithSkipTLSVerify(true),  // Optional: disables TLS certificate verification (default: false)
	client.WithScheme("https"),      // Optional: defaults to "https"
	client.WithBasePath("/api"),     // Optional: defaults to "/api"
)

// List all environments
endpoints, err := cli.ListEndpoints()
if err != nil {
	log.Fatalf("Failed to list Portainer environments: %v", err)
}

// Process endpoints as needed...
```

See the `example/simple/client.go` file for a complete example using the simple client.

### 2. Using the Swagger-Generated Client (Advanced)

The simple client is still a work in progress; if you need to access API operations that aren't yet supported by the simple client, you can use the underlying Swagger-generated client directly.

```go
import (
	"github.com/go-openapi/runtime"
	client "github.com/portainer/client-api-go/v2/pkg/client"
	"github.com/portainer/client-api-go/v2/pkg/client/endpoints"
)

// Create transport
transport := httptransport.New(
	"portainer.dev.local",
	"/api",
	[]string{"https"},
)

// Create client instance
portainerClient := client.New(transport, strfmt.Default)

// Set up API key authentication
apiKeyAuth := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
	return r.SetHeaderParam("x-api-key", "ptr_XXXYYYZZZ")
})
transport.DefaultAuthentication = apiKeyAuth

// List all environments
endpointsParams := endpoints.NewEndpointListParams()
endpointsResp, err := portainerClient.Endpoints.EndpointList(endpointsParams, nil)
if err != nil {
	log.Fatalf("Failed to list Portainer environments: %v", err)
}

// Process endpoints as needed...
```

See the `example/swagger/client.go` file for a complete example using the Swagger-generated client directly.

## Development

To use the latest version of the Portainer API, you must regenerate the underlying API client using Swagger.

1. Install the Swagger CLI:

```sh
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

2. Generate the client (adjust the VERSION parameter as needed):

```sh
make generate-client VERSION=2.27.1
```
