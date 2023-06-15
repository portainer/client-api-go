// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/nomad"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNomadGetTaskEventsCmd returns a cmd to handle operation getTaskEvents
func makeOperationNomadGetTaskEventsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "GetTaskEvents",
		Short: `Allocation ID, namespace and task name params are required
**Access policy**: administrator`,
		RunE: runOperationNomadGetTaskEvents,
	}

	if err := registerOperationNomadGetTaskEventsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNomadGetTaskEvents uses cmd flags to call endpoint api
func runOperationNomadGetTaskEvents(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := nomad.NewGetTaskEventsParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationNomadGetTaskEventsResult(appCli.Nomad.GetTaskEvents(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationNomadGetTaskEventsParamFlags registers all flags needed to fill params
func registerOperationNomadGetTaskEventsParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationNomadGetTaskEventsResult parses request result and return the string content
func parseOperationNomadGetTaskEventsResult(resp0 *nomad.GetTaskEventsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*nomad.GetTaskEventsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getTaskEventsInternalServerError is not supported

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
