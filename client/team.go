package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/team_memberships"
	"github.com/portainer/client-api-go/v2/pkg/client/teams"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListTeams lists all teams
func (c *PortainerClient) ListTeams() ([]*models.PortainerTeam, error) {
	params := teams.NewTeamListParams()
	resp, err := c.cli.Teams.TeamList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list teams: %w", err)
	}

	return resp.Payload, nil
}

// ListTeamMemberships lists all team memberships
func (c *PortainerClient) ListTeamMemberships() ([]*models.PortainerTeamMembership, error) {
	params := team_memberships.NewTeamMembershipListParams()
	resp, err := c.cli.TeamMemberships.TeamMembershipList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list team memberships: %w", err)
	}

	return resp.Payload, nil
}

// UpdateTeam updates the name of a team.
//
// Parameters:
//   - id: The ID of the team to update
//   - name: The new name for the team
func (c *PortainerClient) UpdateTeam(id int, name string) error {
	params := teams.NewTeamUpdateParams().WithID(int64(id)).WithBody(&models.TeamsTeamUpdatePayload{
		Name: name,
	})
	_, _, err := c.cli.Teams.TeamUpdate(params, nil)
	return err
}

// DeleteTeamMembership deletes a team membership.
//
// Parameters:
//   - id: The ID of the team membership to delete
func (c *PortainerClient) DeleteTeamMembership(id int) error {
	params := team_memberships.NewTeamMembershipDeleteParams().WithID(int64(id))
	_, err := c.cli.TeamMemberships.TeamMembershipDelete(params, nil)
	return err
}

// CreateTeamMembership creates a team membership.
//
// Parameters:
//   - teamId: The ID of the team to create the membership for
//   - userId: The ID of the user to add to the team
func (c *PortainerClient) CreateTeamMembership(teamId int, userId int) error {
	teamID := int64(teamId)
	userID := int64(userId)
	params := team_memberships.NewTeamMembershipCreateParams().WithBody(&models.TeammembershipsTeamMembershipCreatePayload{
		TeamID: &teamID,
		UserID: &userID,
	})

	_, _, err := c.cli.TeamMemberships.TeamMembershipCreate(params, nil)
	return err
}
