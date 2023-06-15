// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/client/webhooks"

	"github.com/spf13/cobra"
)

// makeOperationWebhooksDeleteWebhooksIDCmd returns a cmd to handle operation deleteWebhooksId
func makeOperationWebhooksDeleteWebhooksIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "DeleteWebhooksID",
		Short: `**Access policy**: authenticated`,
		RunE:  runOperationWebhooksDeleteWebhooksID,
	}

	if err := registerOperationWebhooksDeleteWebhooksIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationWebhooksDeleteWebhooksID uses cmd flags to call endpoint api
func runOperationWebhooksDeleteWebhooksID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := webhooks.NewDeleteWebhooksIDParams()
	if err, _ := retrieveOperationWebhooksDeleteWebhooksIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationWebhooksDeleteWebhooksIDResult(appCli.Webhooks.DeleteWebhooksID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationWebhooksDeleteWebhooksIDParamFlags registers all flags needed to fill params
func registerOperationWebhooksDeleteWebhooksIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationWebhooksDeleteWebhooksIDIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationWebhooksDeleteWebhooksIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Webhook id`

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

func retrieveOperationWebhooksDeleteWebhooksIDIDFlag(m *webhooks.DeleteWebhooksIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationWebhooksDeleteWebhooksIDResult parses request result and return the string content
func parseOperationWebhooksDeleteWebhooksIDResult(resp0 *webhooks.DeleteWebhooksIDAccepted, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteWebhooksIdAccepted is not supported

		// Non schema case: warning deleteWebhooksIdBadRequest is not supported

		// Non schema case: warning deleteWebhooksIdInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response deleteWebhooksIdAccepted is not supported by go-swagger cli yet.

	return "", nil
}
