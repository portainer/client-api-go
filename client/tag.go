package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/tags"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListTags lists all tags
func (c *PortainerClient) ListTags() ([]*models.PortainerTag, error) {
	params := tags.NewTagListParams()
	resp, err := c.cli.Tags.TagList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}

	return resp.Payload, nil
}

// CreateTag creates a new tag
func (c *PortainerClient) CreateTag(name string) (int64, error) {
	params := tags.NewTagCreateParams().WithBody(&models.TagsTagCreatePayload{
		Name: &name,
	})
	resp, err := c.cli.Tags.TagCreate(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create tag: %w", err)
	}

	return resp.Payload.ID, nil
}
