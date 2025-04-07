package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/settings"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// GetSettings retrieves the Portainer settings
func (c *PortainerClient) GetSettings() (*models.PortainereeSettings, error) {
	params := settings.NewSettingsInspectParams()
	resp, err := c.cli.Settings.SettingsInspect(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get settings: %w", err)
	}

	return resp.Payload, nil
}

// UpdateSettings updates the Portainer settings
func (c *PortainerClient) UpdateSettings(enableEdgeComputeFeatures bool, edgePortainerURL, tunnelServerAddress string) error {
	params := settings.NewSettingsUpdateParams().WithBody(&models.SettingsSettingsUpdatePayload{
		EdgePortainerURL:          edgePortainerURL,
		EnableEdgeComputeFeatures: enableEdgeComputeFeatures,
		Edge: &models.SettingsSettingsUpdatePayloadEdge{
			TunnelServerAddress: tunnelServerAddress,
		},
	})
	_, err := c.cli.Settings.SettingsUpdate(params, nil)
	return err
}
