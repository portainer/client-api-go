package main

import (
	"fmt"
	"log"
	"net/http"

	"crypto/tls"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	client "github.com/portainer/client-api-go/v2/pkg/client"
	"github.com/portainer/client-api-go/v2/pkg/client/endpoints"
)

func main() {
	// Create a transport with TLS verification disabled
	transport := httptransport.New(
		"170.64.222.231:9443",
		"/api",
		[]string{"https"},
	)

	transport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	portainerClient := client.New(transport, strfmt.Default)

	// API key authentication
	apiKeyAuth := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("x-api-key", "ptr_ti2+v9lINWRIWqyKnWsg+ZU7mQ+i9vgJNp4eqvMIMNg=")
	})
	transport.DefaultAuthentication = apiKeyAuth

	// List all endpoints (environments)
	endpointsParams := endpoints.NewEndpointListParams()
	endpointsResp, err := portainerClient.Endpoints.EndpointList(endpointsParams, nil)
	if err != nil {
		log.Fatalf("Error listing endpoints: %v", err)
	}

	fmt.Printf("Found %d endpoints\n", len(endpointsResp.Payload))
	for _, endpoint := range endpointsResp.Payload {
		fmt.Printf("- %s (ID: %d)\n", endpoint.Name, endpoint.ID)
	}
}
