// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for ContainersContainerGpusResponse

// register flags to command
func registerModelContainersContainerGpusResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainersContainerGpusResponseGpus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainersContainerGpusResponseGpus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gpusDescription := ``

	var gpusFlagName string
	if cmdPrefix == "" {
		gpusFlagName = "gpus"
	} else {
		gpusFlagName = fmt.Sprintf("%v.gpus", cmdPrefix)
	}

	var gpusFlagDefault string

	_ = cmd.PersistentFlags().String(gpusFlagName, gpusFlagDefault, gpusDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainersContainerGpusResponseFlags(depth int, m *models.ContainersContainerGpusResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, gpusAdded := retrieveContainersContainerGpusResponseGpusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gpusAdded

	return nil, retAdded
}

func retrieveContainersContainerGpusResponseGpusFlags(depth int, m *models.ContainersContainerGpusResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gpusFlagName := fmt.Sprintf("%v.gpus", cmdPrefix)
	if cmd.Flags().Changed(gpusFlagName) {

		var gpusFlagName string
		if cmdPrefix == "" {
			gpusFlagName = "gpus"
		} else {
			gpusFlagName = fmt.Sprintf("%v.gpus", cmdPrefix)
		}

		gpusFlagValue, err := cmd.Flags().GetString(gpusFlagName)
		if err != nil {
			return err, false
		}
		m.Gpus = gpusFlagValue

		retAdded = true
	}

	return nil, retAdded
}
