// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainereeHelmUserRepository

// register flags to command
func registerModelPortainereeHelmUserRepositoryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeHelmUserRepositoryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeHelmUserRepositoryURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeHelmUserRepositoryUserID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeHelmUserRepositoryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPortainereeHelmUserRepositoryURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := `Helm repository URL`

	var urlFlagName string
	if cmdPrefix == "" {
		urlFlagName = "URL"
	} else {
		urlFlagName = fmt.Sprintf("%v.URL", cmdPrefix)
	}

	var urlFlagDefault string

	_ = cmd.PersistentFlags().String(urlFlagName, urlFlagDefault, urlDescription)

	return nil
}

func registerPortainereeHelmUserRepositoryUserID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userIdDescription := `User identifier`

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
func retrieveModelPortainereeHelmUserRepositoryFlags(depth int, m *models.PortainereeHelmUserRepository, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrievePortainereeHelmUserRepositoryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, urlAdded := retrievePortainereeHelmUserRepositoryURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	err, userIdAdded := retrievePortainereeHelmUserRepositoryUserIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userIdAdded

	return nil, retAdded
}

func retrievePortainereeHelmUserRepositoryIDFlags(depth int, m *models.PortainereeHelmUserRepository, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereeHelmUserRepositoryURLFlags(depth int, m *models.PortainereeHelmUserRepository, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urlFlagName := fmt.Sprintf("%v.URL", cmdPrefix)
	if cmd.Flags().Changed(urlFlagName) {

		var urlFlagName string
		if cmdPrefix == "" {
			urlFlagName = "URL"
		} else {
			urlFlagName = fmt.Sprintf("%v.URL", cmdPrefix)
		}

		urlFlagValue, err := cmd.Flags().GetString(urlFlagName)
		if err != nil {
			return err, false
		}
		m.URL = urlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeHelmUserRepositoryUserIDFlags(depth int, m *models.PortainereeHelmUserRepository, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
