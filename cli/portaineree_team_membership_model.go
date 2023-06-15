// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainereeTeamMembership

// register flags to command
func registerModelPortainereeTeamMembershipFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeTeamMembershipID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTeamMembershipRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTeamMembershipTeamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTeamMembershipUserID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeTeamMembershipID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Membership Identifier`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "Id"
	} else {
		idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerPortainereeTeamMembershipRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleDescription := `Team role (1 for team leader and 2 for team member)`

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "Role"
	} else {
		roleFlagName = fmt.Sprintf("%v.Role", cmdPrefix)
	}

	var roleFlagDefault int64

	_ = cmd.PersistentFlags().Int64(roleFlagName, roleFlagDefault, roleDescription)

	return nil
}

func registerPortainereeTeamMembershipTeamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	teamIdDescription := `Team identifier`

	var teamIdFlagName string
	if cmdPrefix == "" {
		teamIdFlagName = "TeamID"
	} else {
		teamIdFlagName = fmt.Sprintf("%v.TeamID", cmdPrefix)
	}

	var teamIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(teamIdFlagName, teamIdFlagDefault, teamIdDescription)

	return nil
}

func registerPortainereeTeamMembershipUserID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userIdDescription := `User identifier`

	var userIdFlagName string
	if cmdPrefix == "" {
		userIdFlagName = "UserID"
	} else {
		userIdFlagName = fmt.Sprintf("%v.UserID", cmdPrefix)
	}

	var userIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(userIdFlagName, userIdFlagDefault, userIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereeTeamMembershipFlags(depth int, m *models.PortainereeTeamMembership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrievePortainereeTeamMembershipIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, roleAdded := retrievePortainereeTeamMembershipRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	err, teamIdAdded := retrievePortainereeTeamMembershipTeamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || teamIdAdded

	err, userIdAdded := retrievePortainereeTeamMembershipUserIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userIdAdded

	return nil, retAdded
}

func retrievePortainereeTeamMembershipIDFlags(depth int, m *models.PortainereeTeamMembership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.Id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "Id"
		} else {
			idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTeamMembershipRoleFlags(depth int, m *models.PortainereeTeamMembership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleFlagName := fmt.Sprintf("%v.Role", cmdPrefix)
	if cmd.Flags().Changed(roleFlagName) {

		var roleFlagName string
		if cmdPrefix == "" {
			roleFlagName = "Role"
		} else {
			roleFlagName = fmt.Sprintf("%v.Role", cmdPrefix)
		}

		roleFlagValue, err := cmd.Flags().GetInt64(roleFlagName)
		if err != nil {
			return err, false
		}
		m.Role = roleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTeamMembershipTeamIDFlags(depth int, m *models.PortainereeTeamMembership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	teamIdFlagName := fmt.Sprintf("%v.TeamID", cmdPrefix)
	if cmd.Flags().Changed(teamIdFlagName) {

		var teamIdFlagName string
		if cmdPrefix == "" {
			teamIdFlagName = "TeamID"
		} else {
			teamIdFlagName = fmt.Sprintf("%v.TeamID", cmdPrefix)
		}

		teamIdFlagValue, err := cmd.Flags().GetInt64(teamIdFlagName)
		if err != nil {
			return err, false
		}
		m.TeamID = teamIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTeamMembershipUserIDFlags(depth int, m *models.PortainereeTeamMembership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userIdFlagName := fmt.Sprintf("%v.UserID", cmdPrefix)
	if cmd.Flags().Changed(userIdFlagName) {

		var userIdFlagName string
		if cmdPrefix == "" {
			userIdFlagName = "UserID"
		} else {
			userIdFlagName = fmt.Sprintf("%v.UserID", cmdPrefix)
		}

		userIdFlagValue, err := cmd.Flags().GetInt64(userIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserID = userIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
