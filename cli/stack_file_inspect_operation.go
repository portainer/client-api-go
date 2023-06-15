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

// makeOperationStacksStackFileInspectCmd returns a cmd to handle operation stackFileInspect
func makeOperationStacksStackFileInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "StackFileInspect",
		Short: `Get Stack file content.
**Access policy**: restricted`,
		RunE: runOperationStacksStackFileInspect,
	}

	if err := registerOperationStacksStackFileInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStacksStackFileInspect uses cmd flags to call endpoint api
func runOperationStacksStackFileInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stacks.NewStackFileInspectParams()
	if err, _ := retrieveOperationStacksStackFileInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStacksStackFileInspectResult(appCli.Stacks.StackFileInspect(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStacksStackFileInspectParamFlags registers all flags needed to fill params
func registerOperationStacksStackFileInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStacksStackFileInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStacksStackFileInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationStacksStackFileInspectIDFlag(m *stacks.StackFileInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationStacksStackFileInspectResult parses request result and return the string content
func parseOperationStacksStackFileInspectResult(resp0 *stacks.StackFileInspectOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*stacks.StackFileInspectOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning stackFileInspectBadRequest is not supported

		// Non schema case: warning stackFileInspectForbidden is not supported

		// Non schema case: warning stackFileInspectNotFound is not supported

		// Non schema case: warning stackFileInspectInternalServerError is not supported

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
