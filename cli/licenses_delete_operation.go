// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/license"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLicenseLicensesDeleteCmd returns a cmd to handle operation licensesDelete
func makeOperationLicenseLicensesDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "licensesDelete",
		Short: `**Access policy**: administrator`,
		RunE:  runOperationLicenseLicensesDelete,
	}

	if err := registerOperationLicenseLicensesDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLicenseLicensesDelete uses cmd flags to call endpoint api
func runOperationLicenseLicensesDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := license.NewLicensesDeleteParams()
	if err, _ := retrieveOperationLicenseLicensesDeleteBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLicenseLicensesDeleteResult(appCli.License.LicensesDelete(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLicenseLicensesDeleteParamFlags registers all flags needed to fill params
func registerOperationLicenseLicensesDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLicenseLicensesDeleteBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLicenseLicensesDeleteBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. list of license keys to remove")

	// add flags for body
	if err := registerModelLicensesDeletePayloadFlags(0, "licensesDeletePayload", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationLicenseLicensesDeleteBodyFlag(m *license.LicensesDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.LicensesDeletePayload{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.LicensesDeletePayload: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.LicensesDeletePayload{}
	}
	err, added := retrieveModelLicensesDeletePayloadFlags(0, bodyValueModel, "licensesDeletePayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationLicenseLicensesDeleteResult parses request result and return the string content
func parseOperationLicenseLicensesDeleteResult(resp0 *license.LicensesDeleteOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*license.LicensesDeleteOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

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
