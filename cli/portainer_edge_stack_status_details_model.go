// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainerEdgeStackStatusDetails

// register flags to command
func registerModelPortainerEdgeStackStatusDetailsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainerEdgeStackStatusDetailsAcknowledged(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainerEdgeStackStatusDetailsError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainerEdgeStackStatusDetailsImagesPulled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainerEdgeStackStatusDetailsOk(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainerEdgeStackStatusDetailsPending(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainerEdgeStackStatusDetailsRemoteUpdateSuccess(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainerEdgeStackStatusDetailsRemove(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainerEdgeStackStatusDetailsAcknowledged(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	acknowledgedDescription := ``

	var acknowledgedFlagName string
	if cmdPrefix == "" {
		acknowledgedFlagName = "acknowledged"
	} else {
		acknowledgedFlagName = fmt.Sprintf("%v.acknowledged", cmdPrefix)
	}

	var acknowledgedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(acknowledgedFlagName, acknowledgedFlagDefault, acknowledgedDescription)

	return nil
}

func registerPortainerEdgeStackStatusDetailsError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := ``

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault bool

	_ = cmd.PersistentFlags().Bool(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerPortainerEdgeStackStatusDetailsImagesPulled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imagesPulledDescription := ``

	var imagesPulledFlagName string
	if cmdPrefix == "" {
		imagesPulledFlagName = "imagesPulled"
	} else {
		imagesPulledFlagName = fmt.Sprintf("%v.imagesPulled", cmdPrefix)
	}

	var imagesPulledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(imagesPulledFlagName, imagesPulledFlagDefault, imagesPulledDescription)

	return nil
}

func registerPortainerEdgeStackStatusDetailsOk(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	okDescription := ``

	var okFlagName string
	if cmdPrefix == "" {
		okFlagName = "ok"
	} else {
		okFlagName = fmt.Sprintf("%v.ok", cmdPrefix)
	}

	var okFlagDefault bool

	_ = cmd.PersistentFlags().Bool(okFlagName, okFlagDefault, okDescription)

	return nil
}

func registerPortainerEdgeStackStatusDetailsPending(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pendingDescription := ``

	var pendingFlagName string
	if cmdPrefix == "" {
		pendingFlagName = "pending"
	} else {
		pendingFlagName = fmt.Sprintf("%v.pending", cmdPrefix)
	}

	var pendingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(pendingFlagName, pendingFlagDefault, pendingDescription)

	return nil
}

func registerPortainerEdgeStackStatusDetailsRemoteUpdateSuccess(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	remoteUpdateSuccessDescription := ``

	var remoteUpdateSuccessFlagName string
	if cmdPrefix == "" {
		remoteUpdateSuccessFlagName = "remoteUpdateSuccess"
	} else {
		remoteUpdateSuccessFlagName = fmt.Sprintf("%v.remoteUpdateSuccess", cmdPrefix)
	}

	var remoteUpdateSuccessFlagDefault bool

	_ = cmd.PersistentFlags().Bool(remoteUpdateSuccessFlagName, remoteUpdateSuccessFlagDefault, remoteUpdateSuccessDescription)

	return nil
}

func registerPortainerEdgeStackStatusDetailsRemove(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	removeDescription := ``

	var removeFlagName string
	if cmdPrefix == "" {
		removeFlagName = "remove"
	} else {
		removeFlagName = fmt.Sprintf("%v.remove", cmdPrefix)
	}

	var removeFlagDefault bool

	_ = cmd.PersistentFlags().Bool(removeFlagName, removeFlagDefault, removeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainerEdgeStackStatusDetailsFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, acknowledgedAdded := retrievePortainerEdgeStackStatusDetailsAcknowledgedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || acknowledgedAdded

	err, errorAdded := retrievePortainerEdgeStackStatusDetailsErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, imagesPulledAdded := retrievePortainerEdgeStackStatusDetailsImagesPulledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imagesPulledAdded

	err, okAdded := retrievePortainerEdgeStackStatusDetailsOkFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || okAdded

	err, pendingAdded := retrievePortainerEdgeStackStatusDetailsPendingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pendingAdded

	err, remoteUpdateSuccessAdded := retrievePortainerEdgeStackStatusDetailsRemoteUpdateSuccessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || remoteUpdateSuccessAdded

	err, removeAdded := retrievePortainerEdgeStackStatusDetailsRemoveFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || removeAdded

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsAcknowledgedFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	acknowledgedFlagName := fmt.Sprintf("%v.acknowledged", cmdPrefix)
	if cmd.Flags().Changed(acknowledgedFlagName) {

		var acknowledgedFlagName string
		if cmdPrefix == "" {
			acknowledgedFlagName = "acknowledged"
		} else {
			acknowledgedFlagName = fmt.Sprintf("%v.acknowledged", cmdPrefix)
		}

		acknowledgedFlagValue, err := cmd.Flags().GetBool(acknowledgedFlagName)
		if err != nil {
			return err, false
		}
		m.Acknowledged = &acknowledgedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsErrorFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetBool(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = &errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsImagesPulledFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imagesPulledFlagName := fmt.Sprintf("%v.imagesPulled", cmdPrefix)
	if cmd.Flags().Changed(imagesPulledFlagName) {

		var imagesPulledFlagName string
		if cmdPrefix == "" {
			imagesPulledFlagName = "imagesPulled"
		} else {
			imagesPulledFlagName = fmt.Sprintf("%v.imagesPulled", cmdPrefix)
		}

		imagesPulledFlagValue, err := cmd.Flags().GetBool(imagesPulledFlagName)
		if err != nil {
			return err, false
		}
		m.ImagesPulled = &imagesPulledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsOkFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	okFlagName := fmt.Sprintf("%v.ok", cmdPrefix)
	if cmd.Flags().Changed(okFlagName) {

		var okFlagName string
		if cmdPrefix == "" {
			okFlagName = "ok"
		} else {
			okFlagName = fmt.Sprintf("%v.ok", cmdPrefix)
		}

		okFlagValue, err := cmd.Flags().GetBool(okFlagName)
		if err != nil {
			return err, false
		}
		m.Ok = &okFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsPendingFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pendingFlagName := fmt.Sprintf("%v.pending", cmdPrefix)
	if cmd.Flags().Changed(pendingFlagName) {

		var pendingFlagName string
		if cmdPrefix == "" {
			pendingFlagName = "pending"
		} else {
			pendingFlagName = fmt.Sprintf("%v.pending", cmdPrefix)
		}

		pendingFlagValue, err := cmd.Flags().GetBool(pendingFlagName)
		if err != nil {
			return err, false
		}
		m.Pending = &pendingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsRemoteUpdateSuccessFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	remoteUpdateSuccessFlagName := fmt.Sprintf("%v.remoteUpdateSuccess", cmdPrefix)
	if cmd.Flags().Changed(remoteUpdateSuccessFlagName) {

		var remoteUpdateSuccessFlagName string
		if cmdPrefix == "" {
			remoteUpdateSuccessFlagName = "remoteUpdateSuccess"
		} else {
			remoteUpdateSuccessFlagName = fmt.Sprintf("%v.remoteUpdateSuccess", cmdPrefix)
		}

		remoteUpdateSuccessFlagValue, err := cmd.Flags().GetBool(remoteUpdateSuccessFlagName)
		if err != nil {
			return err, false
		}
		m.RemoteUpdateSuccess = &remoteUpdateSuccessFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainerEdgeStackStatusDetailsRemoveFlags(depth int, m *models.PortainerEdgeStackStatusDetails, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	removeFlagName := fmt.Sprintf("%v.remove", cmdPrefix)
	if cmd.Flags().Changed(removeFlagName) {

		var removeFlagName string
		if cmdPrefix == "" {
			removeFlagName = "remove"
		} else {
			removeFlagName = fmt.Sprintf("%v.remove", cmdPrefix)
		}

		removeFlagValue, err := cmd.Flags().GetBool(removeFlagName)
		if err != nil {
			return err, false
		}
		m.Remove = &removeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
