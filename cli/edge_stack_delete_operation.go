// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/client/edge_stacks"

	"github.com/spf13/cobra"
)

// makeOperationEdgeStacksEdgeStackDeleteCmd returns a cmd to handle operation edgeStackDelete
func makeOperationEdgeStacksEdgeStackDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "EdgeStackDelete",
		Short: `**Access policy**: administrator`,
		RunE:  runOperationEdgeStacksEdgeStackDelete,
	}

	if err := registerOperationEdgeStacksEdgeStackDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEdgeStacksEdgeStackDelete uses cmd flags to call endpoint api
func runOperationEdgeStacksEdgeStackDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := edge_stacks.NewEdgeStackDeleteParams()
	if err, _ := retrieveOperationEdgeStacksEdgeStackDeleteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEdgeStacksEdgeStackDeleteResult(appCli.EdgeStacks.EdgeStackDelete(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEdgeStacksEdgeStackDeleteParamFlags registers all flags needed to fill params
func registerOperationEdgeStacksEdgeStackDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEdgeStacksEdgeStackDeleteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEdgeStacksEdgeStackDeleteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationEdgeStacksEdgeStackDeleteIDFlag(m *edge_stacks.EdgeStackDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationEdgeStacksEdgeStackDeleteResult parses request result and return the string content
func parseOperationEdgeStacksEdgeStackDeleteResult(resp0 *edge_stacks.EdgeStackDeleteNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning edgeStackDeleteNoContent is not supported

		// Non schema case: warning edgeStackDeleteBadRequest is not supported

		// Non schema case: warning edgeStackDeleteInternalServerError is not supported

		// Non schema case: warning edgeStackDeleteServiceUnavailable is not supported

		return "", respErr
	}

	// warning: non schema response edgeStackDeleteNoContent is not supported by go-swagger cli yet.

	return "", nil
}
