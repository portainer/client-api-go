// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/endpoint_groups"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEndpointGroupsEndpointGroupListCmd returns a cmd to handle operation endpointGroupList
func makeOperationEndpointGroupsEndpointGroupListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "EndpointGroupList",
		Short: `List all environment(endpoint) groups based on the current user authorizations. Will
return all environment(endpoint) groups if using an administrator account otherwise it will
only return authorized environment(endpoint) groups.
**Access policy**: restricted`,
		RunE: runOperationEndpointGroupsEndpointGroupList,
	}

	if err := registerOperationEndpointGroupsEndpointGroupListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEndpointGroupsEndpointGroupList uses cmd flags to call endpoint api
func runOperationEndpointGroupsEndpointGroupList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := endpoint_groups.NewEndpointGroupListParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEndpointGroupsEndpointGroupListResult(appCli.EndpointGroups.EndpointGroupList(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEndpointGroupsEndpointGroupListParamFlags registers all flags needed to fill params
func registerOperationEndpointGroupsEndpointGroupListParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationEndpointGroupsEndpointGroupListResult parses request result and return the string content
func parseOperationEndpointGroupsEndpointGroupListResult(resp0 *endpoint_groups.EndpointGroupListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*endpoint_groups.EndpointGroupListOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning endpointGroupListInternalServerError is not supported

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
