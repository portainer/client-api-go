package client

import (
	"fmt"
	"strconv"

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
		params.Body.UserAccessPolicies = buildAccessPolicies[models.PortainerUserAccessPolicies](*userAccesses)
	}
	if teamAccesses != nil {
		params.Body.TeamAccessPolicies = buildAccessPolicies[models.PortainerTeamAccessPolicies](*teamAccesses)
	}

	_, err := c.cli.EndpointGroups.EndpointGroupUpdate(params, nil)
	if err != nil {
		return fmt.Errorf("failed to update endpoint group: %w", err)
	}
	return nil
}

func buildAccessPolicies[T models.PortainerTeamAccessPolicies | models.PortainerUserAccessPolicies](accesses map[int64]string) T {
	policies := make(T)
	for id, role := range accesses {
		roleID := getRoleFromName(role)
		if roleID != 0 {
			policies[strconv.Itoa(int(id))] = models.PortainerAccessPolicy{
				RoleID: roleID,
			}
		}
	}
	return policies
}

func getRoleFromName(roleName string) int64 {
	switch roleName {
	case "environment_administrator":
		return 1
	case "helpdesk_user":
		return 2
	case "standard_user":
		return 3
	case "readonly_user":
		return 4
	case "operator_user":
		return 5
	default:
		return 0
	}
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
