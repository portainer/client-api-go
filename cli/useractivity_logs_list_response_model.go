// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for UseractivityLogsListResponse

// register flags to command
func registerModelUseractivityLogsListResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUseractivityLogsListResponseLogs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUseractivityLogsListResponseTotalCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUseractivityLogsListResponseLogs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: logs []*PortainereeUserActivityLog array type is not supported by go-swagger cli yet

	return nil
}

func registerUseractivityLogsListResponseTotalCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalCountDescription := ``

	var totalCountFlagName string
	if cmdPrefix == "" {
		totalCountFlagName = "totalCount"
	} else {
		totalCountFlagName = fmt.Sprintf("%v.totalCount", cmdPrefix)
	}

	var totalCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalCountFlagName, totalCountFlagDefault, totalCountDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUseractivityLogsListResponseFlags(depth int, m *models.UseractivityLogsListResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, logsAdded := retrieveUseractivityLogsListResponseLogsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || logsAdded

	err, totalCountAdded := retrieveUseractivityLogsListResponseTotalCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalCountAdded

	return nil, retAdded
}

func retrieveUseractivityLogsListResponseLogsFlags(depth int, m *models.UseractivityLogsListResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	logsFlagName := fmt.Sprintf("%v.logs", cmdPrefix)
	if cmd.Flags().Changed(logsFlagName) {
		// warning: logs array type []*PortainereeUserActivityLog is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveUseractivityLogsListResponseTotalCountFlags(depth int, m *models.UseractivityLogsListResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalCountFlagName := fmt.Sprintf("%v.totalCount", cmdPrefix)
	if cmd.Flags().Changed(totalCountFlagName) {

		var totalCountFlagName string
		if cmdPrefix == "" {
			totalCountFlagName = "totalCount"
		} else {
			totalCountFlagName = fmt.Sprintf("%v.totalCount", cmdPrefix)
		}

		totalCountFlagValue, err := cmd.Flags().GetInt64(totalCountFlagName)
		if err != nil {
			return err, false
		}
		m.TotalCount = totalCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}
