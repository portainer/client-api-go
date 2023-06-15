// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/endpoints"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEndpointsSnapshotInspectCmd returns a cmd to handle operation snapshotInspect
func makeOperationEndpointsSnapshotInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "snapshotInspect",
		Short: `**Access policy**:`,
		RunE:  runOperationEndpointsSnapshotInspect,
	}

	if err := registerOperationEndpointsSnapshotInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEndpointsSnapshotInspect uses cmd flags to call endpoint api
func runOperationEndpointsSnapshotInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := endpoints.NewSnapshotInspectParams()
	if err, _ := retrieveOperationEndpointsSnapshotInspectEnvironmentIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEndpointsSnapshotInspectResult(appCli.Endpoints.SnapshotInspect(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEndpointsSnapshotInspectParamFlags registers all flags needed to fill params
func registerOperationEndpointsSnapshotInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEndpointsSnapshotInspectEnvironmentIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEndpointsSnapshotInspectEnvironmentIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	environmentIdDescription := `Required. Environment identifier`

	var environmentIdFlagName string
	if cmdPrefix == "" {
		environmentIdFlagName = "environmentId"
	} else {
		environmentIdFlagName = fmt.Sprintf("%v.environmentId", cmdPrefix)
	}

	var environmentIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(environmentIdFlagName, environmentIdFlagDefault, environmentIdDescription)

	return nil
}

func retrieveOperationEndpointsSnapshotInspectEnvironmentIDFlag(m *endpoints.SnapshotInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("environmentId") {

		var environmentIdFlagName string
		if cmdPrefix == "" {
			environmentIdFlagName = "environmentId"
		} else {
			environmentIdFlagName = fmt.Sprintf("%v.environmentId", cmdPrefix)
		}

		environmentIdFlagValue, err := cmd.Flags().GetInt64(environmentIdFlagName)
		if err != nil {
			return err, false
		}
		m.EnvironmentID = environmentIdFlagValue

	}
	return nil, retAdded
}

// parseOperationEndpointsSnapshotInspectResult parses request result and return the string content
func parseOperationEndpointsSnapshotInspectResult(resp0 *endpoints.SnapshotInspectOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*endpoints.SnapshotInspectOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning snapshotInspectNotFound is not supported

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
