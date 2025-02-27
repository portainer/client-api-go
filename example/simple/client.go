package main

import (
	"fmt"
	"log"

	client "github.com/portainer/client-api-go/v2/client"
)

func main() {
	// Initialize client with API key
	portainer := client.NewPortainerClient(
		"portainer.dev.local",
		"ptr_XXXYYYZZZ",
		client.WithSkipTLSVerify(true),
	)

	// List all tags
	tags, err := portainer.ListTags()
	if err != nil {
		log.Fatalf("Failed to list tags: %v", err)
	}

	fmt.Println("Tags:")
	for _, tag := range tags {
		fmt.Printf("- %s (ID: %d)\n", tag.Name, tag.ID)
	}

	// List all endpoints
	endpoints, err := portainer.ListEndpoints()
	if err != nil {
		log.Fatalf("Failed to list endpoints: %v", err)
	}

	fmt.Println("\nEndpoints:")
	for _, endpoint := range endpoints {
		fmt.Printf("- %s (ID: %d, URL: %s)\n", endpoint.Name, endpoint.ID, endpoint.URL)
	}
}
