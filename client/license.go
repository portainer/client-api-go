package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/license"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListLicenses lists all licenses
func (c *PortainerClient) ListLicenses() ([]*models.LiblicensePortainerLicense, error) {
	params := license.NewLicensesListParams()
	resp, err := c.cli.License.LicensesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list licenses: %w", err)
	}

	return resp.Payload, nil
}
