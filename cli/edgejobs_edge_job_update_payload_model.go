// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for EdgejobsEdgeJobUpdatePayload

// register flags to command
func registerModelEdgejobsEdgeJobUpdatePayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEdgejobsEdgeJobUpdatePayloadCronExpression(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEdgejobsEdgeJobUpdatePayloadEdgeGroups(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEdgejobsEdgeJobUpdatePayloadEndpoints(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEdgejobsEdgeJobUpdatePayloadFileContent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEdgejobsEdgeJobUpdatePayloadName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEdgejobsEdgeJobUpdatePayloadRecurring(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEdgejobsEdgeJobUpdatePayloadCronExpression(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cronExpressionDescription := ``

	var cronExpressionFlagName string
	if cmdPrefix == "" {
		cronExpressionFlagName = "cronExpression"
	} else {
		cronExpressionFlagName = fmt.Sprintf("%v.cronExpression", cmdPrefix)
	}

	var cronExpressionFlagDefault string

	_ = cmd.PersistentFlags().String(cronExpressionFlagName, cronExpressionFlagDefault, cronExpressionDescription)

	return nil
}

func registerEdgejobsEdgeJobUpdatePayloadEdgeGroups(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: edgeGroups []int64 array type is not supported by go-swagger cli yet

	return nil
}

func registerEdgejobsEdgeJobUpdatePayloadEndpoints(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: endpoints []int64 array type is not supported by go-swagger cli yet

	return nil
}

func registerEdgejobsEdgeJobUpdatePayloadFileContent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileContentDescription := ``

	var fileContentFlagName string
	if cmdPrefix == "" {
		fileContentFlagName = "fileContent"
	} else {
		fileContentFlagName = fmt.Sprintf("%v.fileContent", cmdPrefix)
	}

	var fileContentFlagDefault string

	_ = cmd.PersistentFlags().String(fileContentFlagName, fileContentFlagDefault, fileContentDescription)

	return nil
}

func registerEdgejobsEdgeJobUpdatePayloadName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEdgejobsEdgeJobUpdatePayloadRecurring(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	recurringDescription := ``

	var recurringFlagName string
	if cmdPrefix == "" {
		recurringFlagName = "recurring"
	} else {
		recurringFlagName = fmt.Sprintf("%v.recurring", cmdPrefix)
	}

	var recurringFlagDefault bool

	_ = cmd.PersistentFlags().Bool(recurringFlagName, recurringFlagDefault, recurringDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEdgejobsEdgeJobUpdatePayloadFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, cronExpressionAdded := retrieveEdgejobsEdgeJobUpdatePayloadCronExpressionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cronExpressionAdded

	err, edgeGroupsAdded := retrieveEdgejobsEdgeJobUpdatePayloadEdgeGroupsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || edgeGroupsAdded

	err, endpointsAdded := retrieveEdgejobsEdgeJobUpdatePayloadEndpointsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointsAdded

	err, fileContentAdded := retrieveEdgejobsEdgeJobUpdatePayloadFileContentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileContentAdded

	err, nameAdded := retrieveEdgejobsEdgeJobUpdatePayloadNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, recurringAdded := retrieveEdgejobsEdgeJobUpdatePayloadRecurringFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || recurringAdded

	return nil, retAdded
}

func retrieveEdgejobsEdgeJobUpdatePayloadCronExpressionFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cronExpressionFlagName := fmt.Sprintf("%v.cronExpression", cmdPrefix)
	if cmd.Flags().Changed(cronExpressionFlagName) {

		var cronExpressionFlagName string
		if cmdPrefix == "" {
			cronExpressionFlagName = "cronExpression"
		} else {
			cronExpressionFlagName = fmt.Sprintf("%v.cronExpression", cmdPrefix)
		}

		cronExpressionFlagValue, err := cmd.Flags().GetString(cronExpressionFlagName)
		if err != nil {
			return err, false
		}
		m.CronExpression = cronExpressionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEdgejobsEdgeJobUpdatePayloadEdgeGroupsFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	edgeGroupsFlagName := fmt.Sprintf("%v.edgeGroups", cmdPrefix)
	if cmd.Flags().Changed(edgeGroupsFlagName) {
		// warning: edgeGroups array type []int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveEdgejobsEdgeJobUpdatePayloadEndpointsFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointsFlagName := fmt.Sprintf("%v.endpoints", cmdPrefix)
	if cmd.Flags().Changed(endpointsFlagName) {
		// warning: endpoints array type []int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveEdgejobsEdgeJobUpdatePayloadFileContentFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileContentFlagName := fmt.Sprintf("%v.fileContent", cmdPrefix)
	if cmd.Flags().Changed(fileContentFlagName) {

		var fileContentFlagName string
		if cmdPrefix == "" {
			fileContentFlagName = "fileContent"
		} else {
			fileContentFlagName = fmt.Sprintf("%v.fileContent", cmdPrefix)
		}

		fileContentFlagValue, err := cmd.Flags().GetString(fileContentFlagName)
		if err != nil {
			return err, false
		}
		m.FileContent = fileContentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEdgejobsEdgeJobUpdatePayloadNameFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEdgejobsEdgeJobUpdatePayloadRecurringFlags(depth int, m *models.EdgejobsEdgeJobUpdatePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	recurringFlagName := fmt.Sprintf("%v.recurring", cmdPrefix)
	if cmd.Flags().Changed(recurringFlagName) {

		var recurringFlagName string
		if cmdPrefix == "" {
			recurringFlagName = "recurring"
		} else {
			recurringFlagName = fmt.Sprintf("%v.recurring", cmdPrefix)
		}

		recurringFlagValue, err := cmd.Flags().GetBool(recurringFlagName)
		if err != nil {
			return err, false
		}
		m.Recurring = &recurringFlagValue

		retAdded = true
	}

	return nil, retAdded
}
