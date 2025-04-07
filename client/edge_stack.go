package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/edge_stacks"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListEdgeStacks lists all edge stacks
func (c *PortainerClient) ListEdgeStacks() ([]*models.PortainereeEdgeStack, error) {
	params := edge_stacks.NewEdgeStackListParams()
	resp, err := c.cli.EdgeStacks.EdgeStackList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list edge stacks: %w", err)
	}

	return resp.Payload, nil
}

// GetEdgeStack returns an edge stack by ID
func (c *PortainerClient) GetEdgeStack(id int64) (*models.PortainereeEdgeStack, error) {
	params := edge_stacks.NewEdgeStackInspectParams().WithID(id)
	resp, err := c.cli.EdgeStacks.EdgeStackInspect(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get edge stack: %w", err)
	}

	return resp.Payload, nil
}

// GetEdgeStackByName returns the first edge stack that matches the specified name
func (c *PortainerClient) GetEdgeStackByName(name string) (*models.PortainereeEdgeStack, error) {
	edgeStacks, err := c.ListEdgeStacks()
	if err != nil {
		return nil, fmt.Errorf("failed to list edge stacks: %w", err)
	}

	for _, edgeStack := range edgeStacks {
		if edgeStack.Name == name {
			return edgeStack, nil
		}
	}

	return nil, fmt.Errorf("edge stack not found")
}

// CreateEdgeStack creates a new edge stack
func (c *PortainerClient) CreateEdgeStack(name string, file string, environmentGroupIds []int64) (int64, error) {
	params := edge_stacks.NewEdgeStackCreateStringParams().WithBody(&models.EdgestacksEdgeStackFromStringPayload{
		Name:             &name,
		StackFileContent: &file,
		EdgeGroups:       environmentGroupIds,
		DeploymentType:   0,
	})

	resp, err := c.cli.EdgeStacks.EdgeStackCreateString(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create edge stack: %w", err)
	}

	return resp.Payload.ID, nil
}

// UpdateEdgeStack updates an existing edge stack
func (c *PortainerClient) UpdateEdgeStack(id int64, file string, environmentGroupIds []int64) error {
	params := edge_stacks.NewEdgeStackUpdateParams().WithID(id).WithBody(&models.EdgestacksUpdateEdgeStackPayload{
		StackFileContent: file,
		EdgeGroups:       environmentGroupIds,
		UpdateVersion:    true,
	})

	_, err := c.cli.EdgeStacks.EdgeStackUpdate(params, nil)
	if err != nil {
		return fmt.Errorf("failed to update edge stack: %w", err)
	}

	return nil
}

// GetEdgeStackFile gets the file for an edge stack
func (c *PortainerClient) GetEdgeStackFile(id int64) (string, error) {
	params := edge_stacks.NewEdgeStackFileParams().WithID(id)
	resp, err := c.cli.EdgeStacks.EdgeStackFile(params, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get edge stack file: %w", err)
	}

	return resp.Payload.StackFileContent, nil
}
