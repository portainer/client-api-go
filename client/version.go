package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/system"
)

// GetVersion returns the version of the Portainer API
func (c *PortainerClient) GetVersion() (string, error) {
	params := system.NewSystemStatusParams()
	resp, err := c.cli.System.SystemStatus(params)
	if err != nil {
		return "", fmt.Errorf("failed to get version: %w", err)
	}

	return resp.Payload.Version, nil
}
