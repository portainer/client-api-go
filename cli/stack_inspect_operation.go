// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/stacks"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStacksStackInspectCmd returns a cmd to handle operation stackInspect
func makeOperationStacksStackInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "StackInspect",
		Short: `Retrieve details about a stack.
**Access policy**: restricted`,
		RunE: runOperationStacksStackInspect,
	}

	if err := registerOperationStacksStackInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStacksStackInspect uses cmd flags to call endpoint api
func runOperationStacksStackInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stacks.NewStackInspectParams()
	if err, _ := retrieveOperationStacksStackInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStacksStackInspectResult(appCli.Stacks.StackInspect(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStacksStackInspectParamFlags registers all flags needed to fill params
func registerOperationStacksStackInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStacksStackInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStacksStackInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Stack identifier`

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

func retrieveOperationStacksStackInspectIDFlag(m *stacks.StackInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationStacksStackInspectResult parses request result and return the string content
func parseOperationStacksStackInspectResult(resp0 *stacks.StackInspectOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*stacks.StackInspectOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning stackInspectBadRequest is not supported

		// Non schema case: warning stackInspectForbidden is not supported

		// Non schema case: warning stackInspectNotFound is not supported

		// Non schema case: warning stackInspectInternalServerError is not supported

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
