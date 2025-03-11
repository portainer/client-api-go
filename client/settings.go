package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/settings"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

func (c *PortainerClient) GetSettings() (*models.PortainereeSettings, error) {
	params := settings.NewSettingsInspectParams()
	resp, err := c.cli.Settings.SettingsInspect(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get settings: %w", err)
	}

	return resp.Payload, nil
}
