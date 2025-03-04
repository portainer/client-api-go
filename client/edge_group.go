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

// UpdateEdgeGroup updates an existing edge group
func (c *PortainerClient) UpdateEdgeGroup(id int64, name string, environmentIds []int64) error {
	params := edge_groups.NewEgeGroupUpdateParams().WithID(id).WithBody(&models.EdgegroupsEdgeGroupUpdatePayload{
		Name:      name,
		Endpoints: environmentIds,
	})

	_, err := c.cli.EdgeGroups.EgeGroupUpdate(params, nil)
	if err != nil {
		return fmt.Errorf("failed to update edge group: %w", err)
	}
	return nil
}
