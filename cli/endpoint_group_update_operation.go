// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/endpoint_groups"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEndpointGroupsEndpointGroupUpdateCmd returns a cmd to handle operation endpointGroupUpdate
func makeOperationEndpointGroupsEndpointGroupUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "EndpointGroupUpdate",
		Short: `Update an environment(endpoint) group.
**Access policy**: administrator`,
		RunE: runOperationEndpointGroupsEndpointGroupUpdate,
	}

	if err := registerOperationEndpointGroupsEndpointGroupUpdateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEndpointGroupsEndpointGroupUpdate uses cmd flags to call endpoint api
func runOperationEndpointGroupsEndpointGroupUpdate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := endpoint_groups.NewEndpointGroupUpdateParams()
	if err, _ := retrieveOperationEndpointGroupsEndpointGroupUpdateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEndpointGroupsEndpointGroupUpdateIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEndpointGroupsEndpointGroupUpdateResult(appCli.EndpointGroups.EndpointGroupUpdate(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEndpointGroupsEndpointGroupUpdateParamFlags registers all flags needed to fill params
func registerOperationEndpointGroupsEndpointGroupUpdateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEndpointGroupsEndpointGroupUpdateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEndpointGroupsEndpointGroupUpdateIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEndpointGroupsEndpointGroupUpdateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. EndpointGroup details")

	// add flags for body
	if err := registerModelEndpointgroupsEndpointGroupUpdatePayloadFlags(0, "endpointgroupsEndpointGroupUpdatePayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationEndpointGroupsEndpointGroupUpdateIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. EndpointGroup identifier`

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

func retrieveOperationEndpointGroupsEndpointGroupUpdateBodyFlag(m *endpoint_groups.EndpointGroupUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.EndpointgroupsEndpointGroupUpdatePayload{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.EndpointgroupsEndpointGroupUpdatePayload: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.EndpointgroupsEndpointGroupUpdatePayload{}
	}
	err, added := retrieveModelEndpointgroupsEndpointGroupUpdatePayloadFlags(0, bodyValueModel, "endpointgroupsEndpointGroupUpdatePayload", cmd)
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
func retrieveOperationEndpointGroupsEndpointGroupUpdateIDFlag(m *endpoint_groups.EndpointGroupUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationEndpointGroupsEndpointGroupUpdateResult parses request result and return the string content
func parseOperationEndpointGroupsEndpointGroupUpdateResult(resp0 *endpoint_groups.EndpointGroupUpdateOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*endpoint_groups.EndpointGroupUpdateOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning endpointGroupUpdateBadRequest is not supported

		// Non schema case: warning endpointGroupUpdateNotFound is not supported

		// Non schema case: warning endpointGroupUpdateInternalServerError is not supported

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
