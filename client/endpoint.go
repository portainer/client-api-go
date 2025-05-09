package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/client/utils"
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

// UpdateEndpoint updates an existing endpoint in Portainer.
// An endpoint represents a Docker environment that Portainer can manage.
//
// Parameters:
//   - id: The ID of the endpoint to update
//   - tagIds: Optional. If provided (including empty slice), updates the endpoint's tags.
//     Use nil to keep existing tags. Slice of tag IDs to associate with the endpoint.
//   - userAccesses: Optional. If provided (including empty map), updates user access policies.
//     Use nil to keep existing policies. Map of user ID to role name.
//   - teamAccesses: Optional. If provided (including empty map), updates team access policies.
//     Use nil to keep existing policies. Map of team ID to role name.
//
// Valid roles for access policies are:
//   - environment_administrator
//   - helpdesk_user
//   - standard_user
//   - readonly_user
//   - operator_user
//
// Invalid roles are ignored.
// Returns an error if the update fails.
func (c *PortainerClient) UpdateEndpoint(id int64, tagIds *[]int64, userAccesses *map[int64]string, teamAccesses *map[int64]string) error {
	params := endpoints.NewEndpointUpdateParams().WithID(id).WithBody(&models.EndpointsEndpointUpdatePayload{})

	if tagIds != nil {
		params.Body.TagIDs = *tagIds
	}

	if userAccesses != nil {
		params.Body.UserAccessPolicies = utils.BuildAccessPolicies[models.PortainerUserAccessPolicies](*userAccesses)
	}

	if teamAccesses != nil {
		params.Body.TeamAccessPolicies = utils.BuildAccessPolicies[models.PortainerTeamAccessPolicies](*teamAccesses)
	}

	_, err := c.cli.Endpoints.EndpointUpdate(params, nil)
	return err
}

// CreateLocalDockerEndpoint creates a new local Docker endpoint
//
// Parameters:
//   - name: The name of the endpoint
//
// Returns the ID of the created endpoint and an error if the creation fails
func (c *PortainerClient) CreateLocalDockerEndpoint(name string) (int64, error) {
	containerEngine := "docker"
	params := endpoints.NewEndpointCreateParams()
	params.Name = name
	params.EndpointCreationType = 1
	params.ContainerEngine = &containerEngine

	resp, err := c.cli.Endpoints.EndpointCreate(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create endpoint: %w", err)
	}

	return resp.Payload.ID, nil
}

// CreateEdgeDockerEndpoint creates a new edge docker endpoint
//
// Parameters:
//   - name: The name of the endpoint
//
// Returns the ID of the created endpoint and an error if the creation fails
func (c *PortainerClient) CreateEdgeDockerEndpoint(name string) (int64, error) {
	containerEngine := "docker"
	params := endpoints.NewEndpointCreateParams()
	params.Name = name
	params.EndpointCreationType = 4
	params.ContainerEngine = &containerEngine

	resp, err := c.cli.Endpoints.EndpointCreate(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create endpoint: %w", err)
	}

	return resp.Payload.ID, nil
}
