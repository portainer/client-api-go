// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for TypesSummaryNetworkSettings

// register flags to command
func registerModelTypesSummaryNetworkSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTypesSummaryNetworkSettingsNetworks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTypesSummaryNetworkSettingsNetworks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: networks map[string]NetworkEndpointSettings map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTypesSummaryNetworkSettingsFlags(depth int, m *models.TypesSummaryNetworkSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, networksAdded := retrieveTypesSummaryNetworkSettingsNetworksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networksAdded

	return nil, retAdded
}

func retrieveTypesSummaryNetworkSettingsNetworksFlags(depth int, m *models.TypesSummaryNetworkSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networksFlagName := fmt.Sprintf("%v.networks", cmdPrefix)
	if cmd.Flags().Changed(networksFlagName) {
		// warning: networks map type map[string]NetworkEndpointSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
