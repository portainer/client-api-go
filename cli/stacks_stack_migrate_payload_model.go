// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for StacksStackMigratePayload

// register flags to command
func registerModelStacksStackMigratePayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStacksStackMigratePayloadEndpointID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStacksStackMigratePayloadName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStacksStackMigratePayloadSwarmID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStacksStackMigratePayloadEndpointID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endpointIdDescription := `Required. `

	var endpointIdFlagName string
	if cmdPrefix == "" {
		endpointIdFlagName = "endpointID"
	} else {
		endpointIdFlagName = fmt.Sprintf("%v.endpointID", cmdPrefix)
	}

	var endpointIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(endpointIdFlagName, endpointIdFlagDefault, endpointIdDescription)

	return nil
}

func registerStacksStackMigratePayloadName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerStacksStackMigratePayloadSwarmID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	swarmIdDescription := ``

	var swarmIdFlagName string
	if cmdPrefix == "" {
		swarmIdFlagName = "swarmID"
	} else {
		swarmIdFlagName = fmt.Sprintf("%v.swarmID", cmdPrefix)
	}

	var swarmIdFlagDefault string

	_ = cmd.PersistentFlags().String(swarmIdFlagName, swarmIdFlagDefault, swarmIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStacksStackMigratePayloadFlags(depth int, m *models.StacksStackMigratePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endpointIdAdded := retrieveStacksStackMigratePayloadEndpointIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointIdAdded

	err, nameAdded := retrieveStacksStackMigratePayloadNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, swarmIdAdded := retrieveStacksStackMigratePayloadSwarmIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || swarmIdAdded

	return nil, retAdded
}

func retrieveStacksStackMigratePayloadEndpointIDFlags(depth int, m *models.StacksStackMigratePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointIdFlagName := fmt.Sprintf("%v.endpointID", cmdPrefix)
	if cmd.Flags().Changed(endpointIdFlagName) {

		var endpointIdFlagName string
		if cmdPrefix == "" {
			endpointIdFlagName = "endpointID"
		} else {
			endpointIdFlagName = fmt.Sprintf("%v.endpointID", cmdPrefix)
		}

		endpointIdFlagValue, err := cmd.Flags().GetInt64(endpointIdFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointID = &endpointIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStacksStackMigratePayloadNameFlags(depth int, m *models.StacksStackMigratePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStacksStackMigratePayloadSwarmIDFlags(depth int, m *models.StacksStackMigratePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	swarmIdFlagName := fmt.Sprintf("%v.swarmID", cmdPrefix)
	if cmd.Flags().Changed(swarmIdFlagName) {

		var swarmIdFlagName string
		if cmdPrefix == "" {
			swarmIdFlagName = "swarmID"
		} else {
			swarmIdFlagName = fmt.Sprintf("%v.swarmID", cmdPrefix)
		}

		swarmIdFlagValue, err := cmd.Flags().GetString(swarmIdFlagName)
		if err != nil {
			return err, false
		}
		m.SwarmID = swarmIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
