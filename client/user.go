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

// CreateUser creates a new user.
func (c *PortainerClient) CreateUser(username, password string, role int64) (int64, error) {
	params := users.NewUserCreateParams().WithBody(&models.UsersUserCreatePayload{
		Username: &username,
		Password: &password,
		Role:     &role,
	})
	resp, err := c.cli.Users.UserCreate(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return resp.Payload.ID, nil
}

// GetUser returns a user by ID
func (c *PortainerClient) GetUser(id int) (*models.PortainereeUser, error) {
	params := users.NewUserInspectParams().WithID(int64(id))
	resp, err := c.cli.Users.UserInspect(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return resp.Payload, nil
}

// UpdateUserRole updates the role of a user.
func (c *PortainerClient) UpdateUserRole(id int, role int64) error {
	params := users.NewUserUpdateParams().WithID(int64(id)).WithBody(&models.UsersUserUpdatePayload{
		Role: &role,
	})
	_, err := c.cli.Users.UserUpdate(params, nil)
	return err
}
