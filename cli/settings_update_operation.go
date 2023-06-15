// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/settings"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSettingsSettingsUpdateCmd returns a cmd to handle operation settingsUpdate
func makeOperationSettingsSettingsUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "SettingsUpdate",
		Short: `Update Portainer settings.
**Access policy**: administrator`,
		RunE: runOperationSettingsSettingsUpdate,
	}

	if err := registerOperationSettingsSettingsUpdateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSettingsSettingsUpdate uses cmd flags to call endpoint api
func runOperationSettingsSettingsUpdate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := settings.NewSettingsUpdateParams()
	if err, _ := retrieveOperationSettingsSettingsUpdateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSettingsSettingsUpdateResult(appCli.Settings.SettingsUpdate(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSettingsSettingsUpdateParamFlags registers all flags needed to fill params
func registerOperationSettingsSettingsUpdateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSettingsSettingsUpdateBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSettingsSettingsUpdateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. New settings")

	// add flags for body
	if err := registerModelSettingsSettingsUpdatePayloadFlags(0, "settingsSettingsUpdatePayload", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationSettingsSettingsUpdateBodyFlag(m *settings.SettingsUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.SettingsSettingsUpdatePayload{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.SettingsSettingsUpdatePayload: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.SettingsSettingsUpdatePayload{}
	}
	err, added := retrieveModelSettingsSettingsUpdatePayloadFlags(0, bodyValueModel, "settingsSettingsUpdatePayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationSettingsSettingsUpdateResult parses request result and return the string content
func parseOperationSettingsSettingsUpdateResult(resp0 *settings.SettingsUpdateOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*settings.SettingsUpdateOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning settingsUpdateBadRequest is not supported

		// Non schema case: warning settingsUpdateInternalServerError is not supported

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
