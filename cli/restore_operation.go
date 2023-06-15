// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/backup"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBackupRestoreCmd returns a cmd to handle operation restore
func makeOperationBackupRestoreCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "Restore",
		Short: `Triggers a system restore using provided backup file
**Access policy**: public`,
		RunE: runOperationBackupRestore,
	}

	if err := registerOperationBackupRestoreParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBackupRestore uses cmd flags to call endpoint api
func runOperationBackupRestore(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := backup.NewRestoreParams()
	if err, _ := retrieveOperationBackupRestoreRestorePayloadFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBackupRestoreResult(appCli.Backup.Restore(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBackupRestoreParamFlags registers all flags needed to fill params
func registerOperationBackupRestoreParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBackupRestoreRestorePayloadParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBackupRestoreRestorePayloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var restorePayloadFlagName string
	if cmdPrefix == "" {
		restorePayloadFlagName = "restorePayload"
	} else {
		restorePayloadFlagName = fmt.Sprintf("%v.restorePayload", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(restorePayloadFlagName, "", "Optional json string for [restorePayload]. Restore request payload")

	// add flags for body
	if err := registerModelBackupRestorePayloadFlags(0, "backupRestorePayload", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationBackupRestoreRestorePayloadFlag(m *backup.RestoreParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("restorePayload") {
		// Read restorePayload string from cmd and unmarshal
		restorePayloadValueStr, err := cmd.Flags().GetString("restorePayload")
		if err != nil {
			return err, false
		}

		restorePayloadValue := models.BackupRestorePayload{}
		if err := json.Unmarshal([]byte(restorePayloadValueStr), &restorePayloadValue); err != nil {
			return fmt.Errorf("cannot unmarshal restorePayload string in models.BackupRestorePayload: %v", err), false
		}
		m.RestorePayload = &restorePayloadValue
	}
	restorePayloadValueModel := m.RestorePayload
	if swag.IsZero(restorePayloadValueModel) {
		restorePayloadValueModel = &models.BackupRestorePayload{}
	}
	err, added := retrieveModelBackupRestorePayloadFlags(0, restorePayloadValueModel, "backupRestorePayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.RestorePayload = restorePayloadValueModel
	}
	if dryRun && debug {

		restorePayloadValueDebugBytes, err := json.Marshal(m.RestorePayload)
		if err != nil {
			return err, false
		}
		logDebugf("RestorePayload dry-run payload: %v", string(restorePayloadValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationBackupRestoreResult parses request result and return the string content
func parseOperationBackupRestoreResult(resp0 *backup.RestoreOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning restoreOK is not supported

		// Non schema case: warning restoreBadRequest is not supported

		// Non schema case: warning restoreInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response restoreOK is not supported by go-swagger cli yet.

	return "", nil
}
