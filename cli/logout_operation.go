// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/client/auth"

	"github.com/spf13/cobra"
)

// makeOperationAuthLogoutCmd returns a cmd to handle operation logout
func makeOperationAuthLogoutCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "Logout",
		Short: `**Access policy**: authenticated`,
		RunE:  runOperationAuthLogout,
	}

	if err := registerOperationAuthLogoutParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthLogout uses cmd flags to call endpoint api
func runOperationAuthLogout(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := auth.NewLogoutParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthLogoutResult(appCli.Auth.Logout(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthLogoutParamFlags registers all flags needed to fill params
func registerOperationAuthLogoutParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationAuthLogoutResult parses request result and return the string content
func parseOperationAuthLogoutResult(resp0 *auth.LogoutNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning logoutNoContent is not supported

		// Non schema case: warning logoutInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response logoutNoContent is not supported by go-swagger cli yet.

	return "", nil
}
