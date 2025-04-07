package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/edge_groups"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListEdgeGroups lists all edge groups
func (c *PortainerClient) ListEdgeGroups() ([]*models.EdgegroupsDecoratedEdgeGroup, error) {
	params := edge_groups.NewEdgeGroupListParams()
	resp, err := c.cli.EdgeGroups.EdgeGroupList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list edge groups: %w", err)
	}

	return resp.Payload, nil
}

// GetEdgeGroupByName returns the first edge group that matches the specified name
func (c *PortainerClient) GetEdgeGroupByName(name string) (*models.EdgegroupsDecoratedEdgeGroup, error) {
	edgeGroups, err := c.ListEdgeGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to list edge groups: %w", err)
	}

	for _, edgeGroup := range edgeGroups {
		if edgeGroup.Name == name {
			return edgeGroup, nil
		}
	}

	return nil, fmt.Errorf("edge group not found")
}

// GetEdgeGroup returns the first edge group that matches the specified ID
// Note that this operation do not leverage the edge_groups.EdgeGroupInspect operation
// because it returns a different object.
func (c *PortainerClient) GetEdgeGroup(id int64) (*models.EdgegroupsDecoratedEdgeGroup, error) {
	edgeGroups, err := c.ListEdgeGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to list edge groups: %w", err)
	}

	for _, edgeGroup := range edgeGroups {
		if edgeGroup.ID == id {
			return edgeGroup, nil
		}
	}

	return nil, fmt.Errorf("edge group not found")
}

// CreateEdgeGroup creates a new edge group
func (c *PortainerClient) CreateEdgeGroup(name string, environmentIds []int64) (int64, error) {
	params := edge_groups.NewEdgeGroupCreateParams().WithBody(&models.EdgegroupsEdgeGroupCreatePayload{
		Name:      name,
		Endpoints: environmentIds,
		Dynamic:   false,
	})

	resp, err := c.cli.EdgeGroups.EdgeGroupCreate(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create edge group: %w", err)
	}
	return resp.Payload.ID, nil
}

// UpdateEdgeGroup updates an existing edge group.
//
// Parameters:
//   - id: The ID of the edge group to update
//   - name: The name of the edge group, required
//   - environmentIds: Optional. If provided, updates the environments associated with the group. Use nil to keep existing environments
//   - tagIds: Optional. If provided, updates the tags associated with the group. Use nil to keep existing tags
//
// Returns an error if the update fails.
func (c *PortainerClient) UpdateEdgeGroup(id int64, name string, environmentIds *[]int64, tagIds *[]int64) error {
	params := edge_groups.NewEgeGroupUpdateParams().WithID(id).WithBody(&models.EdgegroupsEdgeGroupUpdatePayload{
		Name: name,
	})

	if environmentIds != nil {
		params.Body.Endpoints = *environmentIds
	}

	if tagIds != nil {
		params.Body.TagIDs = *tagIds
		params.Body.Dynamic = true
	}

	_, err := c.cli.EdgeGroups.EgeGroupUpdate(params, nil)
	if err != nil {
		return fmt.Errorf("failed to update edge group: %w", err)
	}
	return nil
}
