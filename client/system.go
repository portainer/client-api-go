package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/system"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// GetSystemStatus returns the system status
func (c *PortainerClient) GetSystemStatus() (*models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus, error) {
	params := system.NewSystemStatusParams()
	resp, err := c.cli.System.SystemStatus(params)
	if err != nil {
		return nil, fmt.Errorf("failed to get system status: %w", err)
	}

	return resp.Payload, nil
}
