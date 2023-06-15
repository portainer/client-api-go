// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/spf13/cobra"
)

// Schema cli for TypesContainer

// register flags to command
func registerModelTypesContainerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTypesContainerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerCommand(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerHostConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerImage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerImageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerMounts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerNetworkSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerPorts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerSizeRootFs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerSizeRw(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTypesContainerStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTypesContainerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "Id"
	} else {
		idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerTypesContainerCommand(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commandDescription := ``

	var commandFlagName string
	if cmdPrefix == "" {
		commandFlagName = "command"
	} else {
		commandFlagName = fmt.Sprintf("%v.command", cmdPrefix)
	}

	var commandFlagDefault string

	_ = cmd.PersistentFlags().String(commandFlagName, commandFlagDefault, commandDescription)

	return nil
}

func registerTypesContainerCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdDescription := ``

	var createdFlagName string
	if cmdPrefix == "" {
		createdFlagName = "created"
	} else {
		createdFlagName = fmt.Sprintf("%v.created", cmdPrefix)
	}

	var createdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(createdFlagName, createdFlagDefault, createdDescription)

	return nil
}

func registerTypesContainerHostConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var hostConfigFlagName string
	if cmdPrefix == "" {
		hostConfigFlagName = "hostConfig"
	} else {
		hostConfigFlagName = fmt.Sprintf("%v.hostConfig", cmdPrefix)
	}

	if err := registerModelTypesContainerHostConfigFlags(depth+1, hostConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTypesContainerImage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imageDescription := ``

	var imageFlagName string
	if cmdPrefix == "" {
		imageFlagName = "image"
	} else {
		imageFlagName = fmt.Sprintf("%v.image", cmdPrefix)
	}

	var imageFlagDefault string

	_ = cmd.PersistentFlags().String(imageFlagName, imageFlagDefault, imageDescription)

	return nil
}

func registerTypesContainerImageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imageIdDescription := ``

	var imageIdFlagName string
	if cmdPrefix == "" {
		imageIdFlagName = "imageID"
	} else {
		imageIdFlagName = fmt.Sprintf("%v.imageID", cmdPrefix)
	}

	var imageIdFlagDefault string

	_ = cmd.PersistentFlags().String(imageIdFlagName, imageIdFlagDefault, imageIdDescription)

	return nil
}

func registerTypesContainerLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerTypesContainerMounts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: mounts []*TypesMountPoint array type is not supported by go-swagger cli yet

	return nil
}

func registerTypesContainerNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: names []string array type is not supported by go-swagger cli yet

	return nil
}

func registerTypesContainerNetworkSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var networkSettingsFlagName string
	if cmdPrefix == "" {
		networkSettingsFlagName = "networkSettings"
	} else {
		networkSettingsFlagName = fmt.Sprintf("%v.networkSettings", cmdPrefix)
	}

	if err := registerModelTypesSummaryNetworkSettingsFlags(depth+1, networkSettingsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTypesContainerPorts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ports []*TypesPort array type is not supported by go-swagger cli yet

	return nil
}

func registerTypesContainerSizeRootFs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeRootFsDescription := ``

	var sizeRootFsFlagName string
	if cmdPrefix == "" {
		sizeRootFsFlagName = "sizeRootFs"
	} else {
		sizeRootFsFlagName = fmt.Sprintf("%v.sizeRootFs", cmdPrefix)
	}

	var sizeRootFsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeRootFsFlagName, sizeRootFsFlagDefault, sizeRootFsDescription)

	return nil
}

func registerTypesContainerSizeRw(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeRwDescription := ``

	var sizeRwFlagName string
	if cmdPrefix == "" {
		sizeRwFlagName = "sizeRw"
	} else {
		sizeRwFlagName = fmt.Sprintf("%v.sizeRw", cmdPrefix)
	}

	var sizeRwFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeRwFlagName, sizeRwFlagDefault, sizeRwDescription)

	return nil
}

func registerTypesContainerState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stateDescription := ``

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	return nil
}

func registerTypesContainerStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTypesContainerFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveTypesContainerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, commandAdded := retrieveTypesContainerCommandFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commandAdded

	err, createdAdded := retrieveTypesContainerCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, hostConfigAdded := retrieveTypesContainerHostConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostConfigAdded

	err, imageAdded := retrieveTypesContainerImageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imageAdded

	err, imageIdAdded := retrieveTypesContainerImageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imageIdAdded

	err, labelsAdded := retrieveTypesContainerLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, mountsAdded := retrieveTypesContainerMountsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mountsAdded

	err, namesAdded := retrieveTypesContainerNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || namesAdded

	err, networkSettingsAdded := retrieveTypesContainerNetworkSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkSettingsAdded

	err, portsAdded := retrieveTypesContainerPortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portsAdded

	err, sizeRootFsAdded := retrieveTypesContainerSizeRootFsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeRootFsAdded

	err, sizeRwAdded := retrieveTypesContainerSizeRwFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeRwAdded

	err, stateAdded := retrieveTypesContainerStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stateAdded

	err, statusAdded := retrieveTypesContainerStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveTypesContainerIDFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerCommandFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	commandFlagName := fmt.Sprintf("%v.command", cmdPrefix)
	if cmd.Flags().Changed(commandFlagName) {

		var commandFlagName string
		if cmdPrefix == "" {
			commandFlagName = "command"
		} else {
			commandFlagName = fmt.Sprintf("%v.command", cmdPrefix)
		}

		commandFlagValue, err := cmd.Flags().GetString(commandFlagName)
		if err != nil {
			return err, false
		}
		m.Command = commandFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerCreatedFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdFlagName := fmt.Sprintf("%v.created", cmdPrefix)
	if cmd.Flags().Changed(createdFlagName) {

		var createdFlagName string
		if cmdPrefix == "" {
			createdFlagName = "created"
		} else {
			createdFlagName = fmt.Sprintf("%v.created", cmdPrefix)
		}

		createdFlagValue, err := cmd.Flags().GetInt64(createdFlagName)
		if err != nil {
			return err, false
		}
		m.Created = createdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerHostConfigFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostConfigFlagName := fmt.Sprintf("%v.hostConfig", cmdPrefix)
	if cmd.Flags().Changed(hostConfigFlagName) {
		// info: complex object hostConfig TypesContainerHostConfig is retrieved outside this Changed() block
	}
	hostConfigFlagValue := m.HostConfig
	if swag.IsZero(hostConfigFlagValue) {
		hostConfigFlagValue = &models.TypesContainerHostConfig{}
	}

	err, hostConfigAdded := retrieveModelTypesContainerHostConfigFlags(depth+1, hostConfigFlagValue, hostConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostConfigAdded
	if hostConfigAdded {
		m.HostConfig = hostConfigFlagValue
	}

	return nil, retAdded
}

func retrieveTypesContainerImageFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imageFlagName := fmt.Sprintf("%v.image", cmdPrefix)
	if cmd.Flags().Changed(imageFlagName) {

		var imageFlagName string
		if cmdPrefix == "" {
			imageFlagName = "image"
		} else {
			imageFlagName = fmt.Sprintf("%v.image", cmdPrefix)
		}

		imageFlagValue, err := cmd.Flags().GetString(imageFlagName)
		if err != nil {
			return err, false
		}
		m.Image = imageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerImageIDFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imageIdFlagName := fmt.Sprintf("%v.imageID", cmdPrefix)
	if cmd.Flags().Changed(imageIdFlagName) {

		var imageIdFlagName string
		if cmdPrefix == "" {
			imageIdFlagName = "imageID"
		} else {
			imageIdFlagName = fmt.Sprintf("%v.imageID", cmdPrefix)
		}

		imageIdFlagValue, err := cmd.Flags().GetString(imageIdFlagName)
		if err != nil {
			return err, false
		}
		m.ImageID = imageIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerLabelsFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	labelsFlagName := fmt.Sprintf("%v.labels", cmdPrefix)
	if cmd.Flags().Changed(labelsFlagName) {
		// warning: labels map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTypesContainerMountsFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mountsFlagName := fmt.Sprintf("%v.mounts", cmdPrefix)
	if cmd.Flags().Changed(mountsFlagName) {
		// warning: mounts array type []*TypesMountPoint is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTypesContainerNamesFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	namesFlagName := fmt.Sprintf("%v.names", cmdPrefix)
	if cmd.Flags().Changed(namesFlagName) {
		// warning: names array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTypesContainerNetworkSettingsFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkSettingsFlagName := fmt.Sprintf("%v.networkSettings", cmdPrefix)
	if cmd.Flags().Changed(networkSettingsFlagName) {
		// info: complex object networkSettings TypesSummaryNetworkSettings is retrieved outside this Changed() block
	}
	networkSettingsFlagValue := m.NetworkSettings
	if swag.IsZero(networkSettingsFlagValue) {
		networkSettingsFlagValue = &models.TypesSummaryNetworkSettings{}
	}

	err, networkSettingsAdded := retrieveModelTypesSummaryNetworkSettingsFlags(depth+1, networkSettingsFlagValue, networkSettingsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkSettingsAdded
	if networkSettingsAdded {
		m.NetworkSettings = networkSettingsFlagValue
	}

	return nil, retAdded
}

func retrieveTypesContainerPortsFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portsFlagName := fmt.Sprintf("%v.ports", cmdPrefix)
	if cmd.Flags().Changed(portsFlagName) {
		// warning: ports array type []*TypesPort is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTypesContainerSizeRootFsFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeRootFsFlagName := fmt.Sprintf("%v.sizeRootFs", cmdPrefix)
	if cmd.Flags().Changed(sizeRootFsFlagName) {

		var sizeRootFsFlagName string
		if cmdPrefix == "" {
			sizeRootFsFlagName = "sizeRootFs"
		} else {
			sizeRootFsFlagName = fmt.Sprintf("%v.sizeRootFs", cmdPrefix)
		}

		sizeRootFsFlagValue, err := cmd.Flags().GetInt64(sizeRootFsFlagName)
		if err != nil {
			return err, false
		}
		m.SizeRootFs = sizeRootFsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerSizeRwFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeRwFlagName := fmt.Sprintf("%v.sizeRw", cmdPrefix)
	if cmd.Flags().Changed(sizeRwFlagName) {

		var sizeRwFlagName string
		if cmdPrefix == "" {
			sizeRwFlagName = "sizeRw"
		} else {
			sizeRwFlagName = fmt.Sprintf("%v.sizeRw", cmdPrefix)
		}

		sizeRwFlagValue, err := cmd.Flags().GetInt64(sizeRwFlagName)
		if err != nil {
			return err, false
		}
		m.SizeRw = sizeRwFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerStateFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stateFlagName := fmt.Sprintf("%v.state", cmdPrefix)
	if cmd.Flags().Changed(stateFlagName) {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = stateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTypesContainerStatusFlags(depth int, m *models.TypesContainer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for TypesContainerHostConfig

// register flags to command
func registerModelTypesContainerHostConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTypesContainerHostConfigNetworkMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTypesContainerHostConfigNetworkMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	networkModeDescription := ``

	var networkModeFlagName string
	if cmdPrefix == "" {
		networkModeFlagName = "networkMode"
	} else {
		networkModeFlagName = fmt.Sprintf("%v.networkMode", cmdPrefix)
	}

	var networkModeFlagDefault string

	_ = cmd.PersistentFlags().String(networkModeFlagName, networkModeFlagDefault, networkModeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTypesContainerHostConfigFlags(depth int, m *models.TypesContainerHostConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, networkModeAdded := retrieveTypesContainerHostConfigNetworkModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkModeAdded

	return nil, retAdded
}

func retrieveTypesContainerHostConfigNetworkModeFlags(depth int, m *models.TypesContainerHostConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkModeFlagName := fmt.Sprintf("%v.networkMode", cmdPrefix)
	if cmd.Flags().Changed(networkModeFlagName) {

		var networkModeFlagName string
		if cmdPrefix == "" {
			networkModeFlagName = "networkMode"
		} else {
			networkModeFlagName = fmt.Sprintf("%v.networkMode", cmdPrefix)
		}

		networkModeFlagValue, err := cmd.Flags().GetString(networkModeFlagName)
		if err != nil {
			return err, false
		}
		m.NetworkMode = networkModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
