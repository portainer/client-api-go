package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/users"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListUsers lists all users
func (c *PortainerClient) ListUsers() ([]*models.PortainereeUser, error) {
	params := users.NewUserListParams()
	resp, err := c.cli.Users.UserList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return resp.Payload, nil
}

// UpdateUser updates the role of a user.
func (c *PortainerClient) UpdateUser(id int, role int64) error {
	params := users.NewUserUpdateParams().WithID(int64(id)).WithBody(&models.UsersUserUpdatePayload{
		Role: &role,
	})
	_, err := c.cli.Users.UserUpdate(params, nil)
	return err
}
