// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainereePair

// register flags to command
func registerModelPortainereePairFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereePairName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereePairValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereePairName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPortainereePairValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	valueDescription := ``

	var valueFlagName string
	if cmdPrefix == "" {
		valueFlagName = "value"
	} else {
		valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
	}

	var valueFlagDefault string

	_ = cmd.PersistentFlags().String(valueFlagName, valueFlagDefault, valueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereePairFlags(depth int, m *models.PortainereePair, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrievePortainereePairNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, valueAdded := retrievePortainereePairValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrievePortainereePairNameFlags(depth int, m *models.PortainereePair, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereePairValueFlags(depth int, m *models.PortainereePair, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {

		var valueFlagName string
		if cmdPrefix == "" {
			valueFlagName = "value"
		} else {
			valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
		}

		valueFlagValue, err := cmd.Flags().GetString(valueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = valueFlagValue

		retAdded = true
	}

	return nil, retAdded
}
