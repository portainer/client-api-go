// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/endpoints"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEndpointsEndpointRegistryAccessCmd returns a cmd to handle operation endpointRegistryAccess
func makeOperationEndpointsEndpointRegistryAccessCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "endpointRegistryAccess",
		Short: `**Access policy**: authenticated`,
		RunE:  runOperationEndpointsEndpointRegistryAccess,
	}

	if err := registerOperationEndpointsEndpointRegistryAccessParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEndpointsEndpointRegistryAccess uses cmd flags to call endpoint api
func runOperationEndpointsEndpointRegistryAccess(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := endpoints.NewEndpointRegistryAccessParams()
	if err, _ := retrieveOperationEndpointsEndpointRegistryAccessBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEndpointsEndpointRegistryAccessIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEndpointsEndpointRegistryAccessRegistryIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEndpointsEndpointRegistryAccessResult(appCli.Endpoints.EndpointRegistryAccess(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEndpointsEndpointRegistryAccessParamFlags registers all flags needed to fill params
func registerOperationEndpointsEndpointRegistryAccessParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEndpointsEndpointRegistryAccessBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEndpointsEndpointRegistryAccessIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEndpointsEndpointRegistryAccessRegistryIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEndpointsEndpointRegistryAccessBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. details")

	// add flags for body
	if err := registerModelEndpointsRegistryAccessPayloadFlags(0, "endpointsRegistryAccessPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationEndpointsEndpointRegistryAccessIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Environment(Endpoint) identifier`

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
func registerOperationEndpointsEndpointRegistryAccessRegistryIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	registryIdDescription := `Required. Registry identifier`

	var registryIdFlagName string
	if cmdPrefix == "" {
		registryIdFlagName = "registryId"
	} else {
		registryIdFlagName = fmt.Sprintf("%v.registryId", cmdPrefix)
	}

	var registryIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(registryIdFlagName, registryIdFlagDefault, registryIdDescription)

	return nil
}

func retrieveOperationEndpointsEndpointRegistryAccessBodyFlag(m *endpoints.EndpointRegistryAccessParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.EndpointsRegistryAccessPayload{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.EndpointsRegistryAccessPayload: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.EndpointsRegistryAccessPayload{}
	}
	err, added := retrieveModelEndpointsRegistryAccessPayloadFlags(0, bodyValueModel, "endpointsRegistryAccessPayload", cmd)
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
func retrieveOperationEndpointsEndpointRegistryAccessIDFlag(m *endpoints.EndpointRegistryAccessParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationEndpointsEndpointRegistryAccessRegistryIDFlag(m *endpoints.EndpointRegistryAccessParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("registryId") {

		var registryIdFlagName string
		if cmdPrefix == "" {
			registryIdFlagName = "registryId"
		} else {
			registryIdFlagName = fmt.Sprintf("%v.registryId", cmdPrefix)
		}

		registryIdFlagValue, err := cmd.Flags().GetInt64(registryIdFlagName)
		if err != nil {
			return err, false
		}
		m.RegistryID = registryIdFlagValue

	}
	return nil, retAdded
}

// parseOperationEndpointsEndpointRegistryAccessResult parses request result and return the string content
func parseOperationEndpointsEndpointRegistryAccessResult(resp0 *endpoints.EndpointRegistryAccessNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning endpointRegistryAccessNoContent is not supported

		// Non schema case: warning endpointRegistryAccessBadRequest is not supported

		// Non schema case: warning endpointRegistryAccessForbidden is not supported

		// Non schema case: warning endpointRegistryAccessNotFound is not supported

		// Non schema case: warning endpointRegistryAccessInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response endpointRegistryAccessNoContent is not supported by go-swagger cli yet.

	return "", nil
}
