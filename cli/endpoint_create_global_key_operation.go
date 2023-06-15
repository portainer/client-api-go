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

// makeOperationEndpointsEndpointCreateGlobalKeyCmd returns a cmd to handle operation endpointCreateGlobalKey
func makeOperationEndpointsEndpointCreateGlobalKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "EndpointCreateGlobalKey",
		Short: ``,
		RunE:  runOperationEndpointsEndpointCreateGlobalKey,
	}

	if err := registerOperationEndpointsEndpointCreateGlobalKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEndpointsEndpointCreateGlobalKey uses cmd flags to call endpoint api
func runOperationEndpointsEndpointCreateGlobalKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := endpoints.NewEndpointCreateGlobalKeyParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEndpointsEndpointCreateGlobalKeyResult(appCli.Endpoints.EndpointCreateGlobalKey(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEndpointsEndpointCreateGlobalKeyParamFlags registers all flags needed to fill params
func registerOperationEndpointsEndpointCreateGlobalKeyParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationEndpointsEndpointCreateGlobalKeyResult parses request result and return the string content
func parseOperationEndpointsEndpointCreateGlobalKeyResult(resp0 *endpoints.EndpointCreateGlobalKeyOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*endpoints.EndpointCreateGlobalKeyOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning endpointCreateGlobalKeyBadRequest is not supported

		// Non schema case: warning endpointCreateGlobalKeyInternalServerError is not supported

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
