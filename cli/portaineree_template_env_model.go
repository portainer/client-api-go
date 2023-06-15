// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainereeTemplateEnv

// register flags to command
func registerModelPortainereeTemplateEnvFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeTemplateEnvDefault(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTemplateEnvDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTemplateEnvLabel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTemplateEnvName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTemplateEnvPreset(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeTemplateEnvSelect(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeTemplateEnvDefault(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	defaultDescription := `Default value that will be set for the variable`

	var defaultFlagName string
	if cmdPrefix == "" {
		defaultFlagName = "default"
	} else {
		defaultFlagName = fmt.Sprintf("%v.default", cmdPrefix)
	}

	var defaultFlagDefault string

	_ = cmd.PersistentFlags().String(defaultFlagName, defaultFlagDefault, defaultDescription)

	return nil
}

func registerPortainereeTemplateEnvDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Content of the tooltip that will be generated in the UI`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerPortainereeTemplateEnvLabel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	labelDescription := `Text for the label that will be generated in the UI`

	var labelFlagName string
	if cmdPrefix == "" {
		labelFlagName = "label"
	} else {
		labelFlagName = fmt.Sprintf("%v.label", cmdPrefix)
	}

	var labelFlagDefault string

	_ = cmd.PersistentFlags().String(labelFlagName, labelFlagDefault, labelDescription)

	return nil
}

func registerPortainereeTemplateEnvName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `name of the environment(endpoint) variable`

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

func registerPortainereeTemplateEnvPreset(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	presetDescription := `If set to true, will not generate any input for this variable in the UI`

	var presetFlagName string
	if cmdPrefix == "" {
		presetFlagName = "preset"
	} else {
		presetFlagName = fmt.Sprintf("%v.preset", cmdPrefix)
	}

	var presetFlagDefault bool

	_ = cmd.PersistentFlags().Bool(presetFlagName, presetFlagDefault, presetDescription)

	return nil
}

func registerPortainereeTemplateEnvSelect(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: select []*PortainereeTemplateEnvSelect array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereeTemplateEnvFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, defaultAdded := retrievePortainereeTemplateEnvDefaultFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultAdded

	err, descriptionAdded := retrievePortainereeTemplateEnvDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, labelAdded := retrievePortainereeTemplateEnvLabelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelAdded

	err, nameAdded := retrievePortainereeTemplateEnvNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, presetAdded := retrievePortainereeTemplateEnvPresetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || presetAdded

	err, selectAdded := retrievePortainereeTemplateEnvSelectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || selectAdded

	return nil, retAdded
}

func retrievePortainereeTemplateEnvDefaultFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultFlagName := fmt.Sprintf("%v.default", cmdPrefix)
	if cmd.Flags().Changed(defaultFlagName) {

		var defaultFlagName string
		if cmdPrefix == "" {
			defaultFlagName = "default"
		} else {
			defaultFlagName = fmt.Sprintf("%v.default", cmdPrefix)
		}

		defaultFlagValue, err := cmd.Flags().GetString(defaultFlagName)
		if err != nil {
			return err, false
		}
		m.Default = defaultFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTemplateEnvDescriptionFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTemplateEnvLabelFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	labelFlagName := fmt.Sprintf("%v.label", cmdPrefix)
	if cmd.Flags().Changed(labelFlagName) {

		var labelFlagName string
		if cmdPrefix == "" {
			labelFlagName = "label"
		} else {
			labelFlagName = fmt.Sprintf("%v.label", cmdPrefix)
		}

		labelFlagValue, err := cmd.Flags().GetString(labelFlagName)
		if err != nil {
			return err, false
		}
		m.Label = labelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTemplateEnvNameFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereeTemplateEnvPresetFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	presetFlagName := fmt.Sprintf("%v.preset", cmdPrefix)
	if cmd.Flags().Changed(presetFlagName) {

		var presetFlagName string
		if cmdPrefix == "" {
			presetFlagName = "preset"
		} else {
			presetFlagName = fmt.Sprintf("%v.preset", cmdPrefix)
		}

		presetFlagValue, err := cmd.Flags().GetBool(presetFlagName)
		if err != nil {
			return err, false
		}
		m.Preset = &presetFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeTemplateEnvSelectFlags(depth int, m *models.PortainereeTemplateEnv, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	selectFlagName := fmt.Sprintf("%v.select", cmdPrefix)
	if cmd.Flags().Changed(selectFlagName) {
		// warning: select array type []*PortainereeTemplateEnvSelect is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
