package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/client/utils"
	"github.com/portainer/client-api-go/v2/pkg/client/endpoint_groups"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListEndpointGroups lists all endpoint groups
func (c *PortainerClient) ListEndpointGroups() ([]*models.PortainerEndpointGroup, error) {
	params := endpoint_groups.NewEndpointGroupListParams()
	resp, err := c.cli.EndpointGroups.EndpointGroupList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoint groups: %w", err)
	}

	return resp.Payload, nil
}

// CreateEndpointGroup creates a new endpoint group
func (c *PortainerClient) CreateEndpointGroup(name string, associatedEndpoints []int64) (int64, error) {
	params := endpoint_groups.NewPostEndpointGroupsParams()
	params.Body = &models.EndpointgroupsEndpointGroupCreatePayload{
		Name:                &name,
		AssociatedEndpoints: associatedEndpoints,
	}
	resp, err := c.cli.EndpointGroups.PostEndpointGroups(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create endpoint group: %w", err)
	}

	return resp.Payload.ID, nil
}

// GetEndpointGroup returns an endpoint group by ID
func (c *PortainerClient) GetEndpointGroup(id int64) (*models.PortainerEndpointGroup, error) {
	params := endpoint_groups.NewGetEndpointGroupsIDParams().WithID(id)
	resp, err := c.cli.EndpointGroups.GetEndpointGroupsID(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint group: %w", err)
	}

	return resp.Payload, nil
}

// GetEndpointGroupByName returns the first endpoint group that matches the specified name
func (c *PortainerClient) GetEndpointGroupByName(name string) (*models.PortainerEndpointGroup, error) {
	endpointGroups, err := c.ListEndpointGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoint groups: %w", err)
	}

	for _, endpointGroup := range endpointGroups {
		if endpointGroup.Name == name {
			return endpointGroup, nil
		}
	}

	return nil, fmt.Errorf("endpoint group not found")
}

// UpdateEndpointGroup updates an existing endpoint group.
//
// Parameters:
//   - id: The ID of the endpoint group to update
//   - name: Optional. If provided, updates the group name. Use nil to keep existing name
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
func (c *PortainerClient) UpdateEndpointGroup(id int64, name *string, userAccesses *map[int64]string, teamAccesses *map[int64]string) error {
	params := endpoint_groups.NewEndpointGroupUpdateParams().WithID(id).WithBody(&models.EndpointgroupsEndpointGroupUpdatePayload{})

	if name != nil {
		params.Body.Name = *name
	}
	if userAccesses != nil {
		params.Body.UserAccessPolicies = utils.BuildAccessPolicies[models.PortainerUserAccessPolicies](*userAccesses)
	}
	if teamAccesses != nil {
		params.Body.TeamAccessPolicies = utils.BuildAccessPolicies[models.PortainerTeamAccessPolicies](*teamAccesses)
	}

	_, err := c.cli.EndpointGroups.EndpointGroupUpdate(params, nil)
	if err != nil {
		return fmt.Errorf("failed to update endpoint group: %w", err)
	}
	return nil
}

// AddEnvironmentToEndpointGroup adds an environment to an endpoint group
func (c *PortainerClient) AddEnvironmentToEndpointGroup(groupId int64, environmentId int64) error {
	params := endpoint_groups.NewEndpointGroupAddEndpointParams().WithID(groupId).WithEndpointID(environmentId)
	_, err := c.cli.EndpointGroups.EndpointGroupAddEndpoint(params, nil)
	if err != nil {
		return fmt.Errorf("failed to add environment to endpoint group: %w", err)
	}
	return nil
}

// RemoveEnvironmentFromEndpointGroup removes an environment from an endpoint group
func (c *PortainerClient) RemoveEnvironmentFromEndpointGroup(groupId int64, environmentId int64) error {
	params := endpoint_groups.NewEndpointGroupDeleteEndpointParams().WithID(groupId).WithEndpointID(environmentId)
	_, err := c.cli.EndpointGroups.EndpointGroupDeleteEndpoint(params, nil)
	if err != nil {
		return fmt.Errorf("failed to remove environment from endpoint group: %w", err)
	}
	return nil
}
