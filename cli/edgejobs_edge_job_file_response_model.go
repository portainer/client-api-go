// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for EdgejobsEdgeJobFileResponse

// register flags to command
func registerModelEdgejobsEdgeJobFileResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEdgejobsEdgeJobFileResponseFileContent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEdgejobsEdgeJobFileResponseFileContent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileContentDescription := ``

	var fileContentFlagName string
	if cmdPrefix == "" {
		fileContentFlagName = "FileContent"
	} else {
		fileContentFlagName = fmt.Sprintf("%v.FileContent", cmdPrefix)
	}

	var fileContentFlagDefault string

	_ = cmd.PersistentFlags().String(fileContentFlagName, fileContentFlagDefault, fileContentDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEdgejobsEdgeJobFileResponseFlags(depth int, m *models.EdgejobsEdgeJobFileResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, fileContentAdded := retrieveEdgejobsEdgeJobFileResponseFileContentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileContentAdded

	return nil, retAdded
}

func retrieveEdgejobsEdgeJobFileResponseFileContentFlags(depth int, m *models.EdgejobsEdgeJobFileResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileContentFlagName := fmt.Sprintf("%v.FileContent", cmdPrefix)
	if cmd.Flags().Changed(fileContentFlagName) {

		var fileContentFlagName string
		if cmdPrefix == "" {
			fileContentFlagName = "FileContent"
		} else {
			fileContentFlagName = fmt.Sprintf("%v.FileContent", cmdPrefix)
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
