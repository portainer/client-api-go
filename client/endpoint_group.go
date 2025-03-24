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

// UpdateEndpointGroup updates an existing endpoint group
// User and Team access policies are maps of user/team ID to role name.
// Valid roles are environment_administrator, helpdesk_user, standard_user, readonly_user or operator_user.
// Invalid roles are ignored.
func (c *PortainerClient) UpdateEndpointGroup(id int64, name string, userAccesses map[int64]string, teamAccesses map[int64]string) error {
	params := endpoint_groups.NewEndpointGroupUpdateParams().WithID(id).WithBody(&models.EndpointgroupsEndpointGroupUpdatePayload{
		Name:               name,
		TeamAccessPolicies: buildAccessPolicies[models.PortainerTeamAccessPolicies](teamAccesses),
		UserAccessPolicies: buildAccessPolicies[models.PortainerUserAccessPolicies](userAccesses),
	})
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
