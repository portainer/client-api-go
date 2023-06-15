// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainereeUserResourceAccess

// register flags to command
func registerModelPortainereeUserResourceAccessFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeUserResourceAccessAccessLevel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeUserResourceAccessUserID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeUserResourceAccessAccessLevel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessLevelDescription := ``

	var accessLevelFlagName string
	if cmdPrefix == "" {
		accessLevelFlagName = "AccessLevel"
	} else {
		accessLevelFlagName = fmt.Sprintf("%v.AccessLevel", cmdPrefix)
	}

	var accessLevelFlagDefault int64

	_ = cmd.PersistentFlags().Int64(accessLevelFlagName, accessLevelFlagDefault, accessLevelDescription)

	return nil
}

func registerPortainereeUserResourceAccessUserID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userIdDescription := `User identifier who created this template`

	var userIdFlagName string
	if cmdPrefix == "" {
		userIdFlagName = "UserId"
	} else {
		userIdFlagName = fmt.Sprintf("%v.UserId", cmdPrefix)
	}

	var userIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(userIdFlagName, userIdFlagDefault, userIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereeUserResourceAccessFlags(depth int, m *models.PortainereeUserResourceAccess, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accessLevelAdded := retrievePortainereeUserResourceAccessAccessLevelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessLevelAdded

	err, userIdAdded := retrievePortainereeUserResourceAccessUserIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userIdAdded

	return nil, retAdded
}

func retrievePortainereeUserResourceAccessAccessLevelFlags(depth int, m *models.PortainereeUserResourceAccess, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessLevelFlagName := fmt.Sprintf("%v.AccessLevel", cmdPrefix)
	if cmd.Flags().Changed(accessLevelFlagName) {

		var accessLevelFlagName string
		if cmdPrefix == "" {
			accessLevelFlagName = "AccessLevel"
		} else {
			accessLevelFlagName = fmt.Sprintf("%v.AccessLevel", cmdPrefix)
		}

		accessLevelFlagValue, err := cmd.Flags().GetInt64(accessLevelFlagName)
		if err != nil {
			return err, false
		}
		m.AccessLevel = accessLevelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeUserResourceAccessUserIDFlags(depth int, m *models.PortainereeUserResourceAccess, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userIdFlagName := fmt.Sprintf("%v.UserId", cmdPrefix)
	if cmd.Flags().Changed(userIdFlagName) {

		var userIdFlagName string
		if cmdPrefix == "" {
			userIdFlagName = "UserId"
		} else {
			userIdFlagName = fmt.Sprintf("%v.UserId", cmdPrefix)
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
