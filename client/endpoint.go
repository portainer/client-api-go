package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/endpoints"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListEndpoints lists all endpoints
func (c *PortainerClient) ListEndpoints() ([]*models.PortainereeEndpoint, error) {
	params := endpoints.NewEndpointListParams()
	resp, err := c.cli.Endpoints.EndpointList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoints: %w", err)
	}

	return resp.Payload, nil
}

// GetEndpoint gets a specific endpoint by ID
func (c *PortainerClient) GetEndpoint(id int64) (*models.PortainereeEndpoint, error) {
	params := endpoints.NewEndpointInspectParams().WithID(id)
	resp, err := c.cli.Endpoints.EndpointInspect(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	return resp.Payload, nil
}

// UpdateEndpoint updates the tags for an endpoint.
func (c *PortainerClient) UpdateEndpoint(id int64, tagIds []int64, userAccesses map[int64]string, teamAccesses map[int64]string) error {
	params := endpoints.NewEndpointUpdateParams().WithID(id).WithBody(&models.EndpointsEndpointUpdatePayload{
		TagIDs:             tagIds,
		UserAccessPolicies: buildAccessPolicies[models.PortainerUserAccessPolicies](userAccesses),
		TeamAccessPolicies: buildAccessPolicies[models.PortainerTeamAccessPolicies](teamAccesses),
	})
	_, err := c.cli.Endpoints.EndpointUpdate(params, nil)
	return err
}
