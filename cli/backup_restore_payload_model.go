// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for BackupRestorePayload

// register flags to command
func registerModelBackupRestorePayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBackupRestorePayloadFileContent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestorePayloadFileName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestorePayloadPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBackupRestorePayloadFileContent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: fileContent []int64 array type is not supported by go-swagger cli yet

	return nil
}

func registerBackupRestorePayloadFileName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileNameDescription := `Required. File name`

	var fileNameFlagName string
	if cmdPrefix == "" {
		fileNameFlagName = "fileName"
	} else {
		fileNameFlagName = fmt.Sprintf("%v.fileName", cmdPrefix)
	}

	var fileNameFlagDefault string

	_ = cmd.PersistentFlags().String(fileNameFlagName, fileNameFlagDefault, fileNameDescription)

	return nil
}

func registerBackupRestorePayloadPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Password to decrypt the backup with`

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBackupRestorePayloadFlags(depth int, m *models.BackupRestorePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, fileContentAdded := retrieveBackupRestorePayloadFileContentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileContentAdded

	err, fileNameAdded := retrieveBackupRestorePayloadFileNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileNameAdded

	err, passwordAdded := retrieveBackupRestorePayloadPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	return nil, retAdded
}

func retrieveBackupRestorePayloadFileContentFlags(depth int, m *models.BackupRestorePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileContentFlagName := fmt.Sprintf("%v.fileContent", cmdPrefix)
	if cmd.Flags().Changed(fileContentFlagName) {
		// warning: fileContent array type []int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveBackupRestorePayloadFileNameFlags(depth int, m *models.BackupRestorePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileNameFlagName := fmt.Sprintf("%v.fileName", cmdPrefix)
	if cmd.Flags().Changed(fileNameFlagName) {

		var fileNameFlagName string
		if cmdPrefix == "" {
			fileNameFlagName = "fileName"
		} else {
			fileNameFlagName = fmt.Sprintf("%v.fileName", cmdPrefix)
		}

		fileNameFlagValue, err := cmd.Flags().GetString(fileNameFlagName)
		if err != nil {
			return err, false
		}
		m.FileName = &fileNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBackupRestorePayloadPasswordFlags(depth int, m *models.BackupRestorePayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
		}

		passwordFlagValue, err := cmd.Flags().GetString(passwordFlagName)
		if err != nil {
			return err, false
		}
		m.Password = passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}
