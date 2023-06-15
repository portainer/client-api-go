// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/registries"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationRegistriesRegistryUpdateCmd returns a cmd to handle operation registryUpdate
func makeOperationRegistriesRegistryUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "RegistryUpdate",
		Short: `Update a registry
**Access policy**: restricted`,
		RunE: runOperationRegistriesRegistryUpdate,
	}

	if err := registerOperationRegistriesRegistryUpdateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRegistriesRegistryUpdate uses cmd flags to call endpoint api
func runOperationRegistriesRegistryUpdate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := registries.NewRegistryUpdateParams()
	if err, _ := retrieveOperationRegistriesRegistryUpdateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRegistriesRegistryUpdateIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRegistriesRegistryUpdateResult(appCli.Registries.RegistryUpdate(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRegistriesRegistryUpdateParamFlags registers all flags needed to fill params
func registerOperationRegistriesRegistryUpdateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRegistriesRegistryUpdateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRegistriesRegistryUpdateIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRegistriesRegistryUpdateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. Registry details")

	// add flags for body
	if err := registerModelRegistriesRegistryUpdatePayloadFlags(0, "registriesRegistryUpdatePayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationRegistriesRegistryUpdateIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Registry identifier`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationRegistriesRegistryUpdateBodyFlag(m *registries.RegistryUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.RegistriesRegistryUpdatePayload{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.RegistriesRegistryUpdatePayload: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.RegistriesRegistryUpdatePayload{}
	}
	err, added := retrieveModelRegistriesRegistryUpdatePayloadFlags(0, bodyValueModel, "registriesRegistryUpdatePayload", cmd)
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
func retrieveOperationRegistriesRegistryUpdateIDFlag(m *registries.RegistryUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationRegistriesRegistryUpdateResult parses request result and return the string content
func parseOperationRegistriesRegistryUpdateResult(resp0 *registries.RegistryUpdateOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*registries.RegistryUpdateOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning registryUpdateBadRequest is not supported

		// Non schema case: warning registryUpdateNotFound is not supported

		// Non schema case: warning registryUpdateConflict is not supported

		// Non schema case: warning registryUpdateInternalServerError is not supported

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
