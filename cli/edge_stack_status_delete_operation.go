// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/edge_stacks"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEdgeStacksEdgeStackStatusDeleteCmd returns a cmd to handle operation edgeStackStatusDelete
func makeOperationEdgeStacksEdgeStackStatusDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "EdgeStackStatusDelete",
		Short: `Authorized only if the request is done by an Edge Environment(Endpoint)`,
		RunE:  runOperationEdgeStacksEdgeStackStatusDelete,
	}

	if err := registerOperationEdgeStacksEdgeStackStatusDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEdgeStacksEdgeStackStatusDelete uses cmd flags to call endpoint api
func runOperationEdgeStacksEdgeStackStatusDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := edge_stacks.NewEdgeStackStatusDeleteParams()
	if err, _ := retrieveOperationEdgeStacksEdgeStackStatusDeleteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEdgeStacksEdgeStackStatusDeleteResult(appCli.EdgeStacks.EdgeStackStatusDelete(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEdgeStacksEdgeStackStatusDeleteParamFlags registers all flags needed to fill params
func registerOperationEdgeStacksEdgeStackStatusDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEdgeStacksEdgeStackStatusDeleteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEdgeStacksEdgeStackStatusDeleteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. EdgeStack Id`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationEdgeStacksEdgeStackStatusDeleteIDFlag(m *edge_stacks.EdgeStackStatusDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationEdgeStacksEdgeStackStatusDeleteResult parses request result and return the string content
func parseOperationEdgeStacksEdgeStackStatusDeleteResult(resp0 *edge_stacks.EdgeStackStatusDeleteOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*edge_stacks.EdgeStackStatusDeleteOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning edgeStackStatusDeleteBadRequest is not supported

		// Non schema case: warning edgeStackStatusDeleteForbidden is not supported

		// Non schema case: warning edgeStackStatusDeleteNotFound is not supported

		// Non schema case: warning edgeStackStatusDeleteInternalServerError is not supported

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
