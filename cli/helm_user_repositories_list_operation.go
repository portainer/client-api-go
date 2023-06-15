// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/helm"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHelmHelmUserRepositoriesListCmd returns a cmd to handle operation helmUserRepositoriesList
func makeOperationHelmHelmUserRepositoriesListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "HelmUserRepositoriesList",
		Short: `Inspect a user helm repositories.
**Access policy**: authenticated`,
		RunE: runOperationHelmHelmUserRepositoriesList,
	}

	if err := registerOperationHelmHelmUserRepositoriesListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHelmHelmUserRepositoriesList uses cmd flags to call endpoint api
func runOperationHelmHelmUserRepositoriesList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := helm.NewHelmUserRepositoriesListParams()
	if err, _ := retrieveOperationHelmHelmUserRepositoriesListIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHelmHelmUserRepositoriesListResult(appCli.Helm.HelmUserRepositoriesList(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHelmHelmUserRepositoriesListParamFlags registers all flags needed to fill params
func registerOperationHelmHelmUserRepositoriesListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHelmHelmUserRepositoriesListIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHelmHelmUserRepositoriesListIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Environment(Endpoint) identifier`

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

func retrieveOperationHelmHelmUserRepositoriesListIDFlag(m *helm.HelmUserRepositoriesListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationHelmHelmUserRepositoriesListResult parses request result and return the string content
func parseOperationHelmHelmUserRepositoriesListResult(resp0 *helm.HelmUserRepositoriesListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*helm.HelmUserRepositoriesListOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning helmUserRepositoriesListBadRequest is not supported

		// Non schema case: warning helmUserRepositoriesListForbidden is not supported

		// Non schema case: warning helmUserRepositoriesListInternalServerError is not supported

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
