// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for BackupRestoreS3Settings

// register flags to command
func registerModelBackupRestoreS3SettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBackupRestoreS3SettingsAccessKeyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestoreS3SettingsBucketName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestoreS3SettingsFilename(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestoreS3SettingsPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestoreS3SettingsRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestoreS3SettingsS3CompatibleHost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBackupRestoreS3SettingsSecretAccessKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBackupRestoreS3SettingsAccessKeyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessKeyIdDescription := `AWS access key id`

	var accessKeyIdFlagName string
	if cmdPrefix == "" {
		accessKeyIdFlagName = "accessKeyID"
	} else {
		accessKeyIdFlagName = fmt.Sprintf("%v.accessKeyID", cmdPrefix)
	}

	var accessKeyIdFlagDefault string

	_ = cmd.PersistentFlags().String(accessKeyIdFlagName, accessKeyIdFlagDefault, accessKeyIdDescription)

	return nil
}

func registerBackupRestoreS3SettingsBucketName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bucketNameDescription := `AWS S3 bucket name`

	var bucketNameFlagName string
	if cmdPrefix == "" {
		bucketNameFlagName = "bucketName"
	} else {
		bucketNameFlagName = fmt.Sprintf("%v.bucketName", cmdPrefix)
	}

	var bucketNameFlagDefault string

	_ = cmd.PersistentFlags().String(bucketNameFlagName, bucketNameFlagDefault, bucketNameDescription)

	return nil
}

func registerBackupRestoreS3SettingsFilename(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	filenameDescription := `AWS S3 filename in the bucket`

	var filenameFlagName string
	if cmdPrefix == "" {
		filenameFlagName = "filename"
	} else {
		filenameFlagName = fmt.Sprintf("%v.filename", cmdPrefix)
	}

	var filenameFlagDefault string

	_ = cmd.PersistentFlags().String(filenameFlagName, filenameFlagDefault, filenameDescription)

	return nil
}

func registerBackupRestoreS3SettingsPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := ``

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

func registerBackupRestoreS3SettingsRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionDescription := `AWS S3 region. Default to "us-east-1"`

	var regionFlagName string
	if cmdPrefix == "" {
		regionFlagName = "region"
	} else {
		regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
	}

	var regionFlagDefault string

	_ = cmd.PersistentFlags().String(regionFlagName, regionFlagDefault, regionDescription)

	return nil
}

func registerBackupRestoreS3SettingsS3CompatibleHost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	s3CompatibleHostDescription := `S3 compatible host`

	var s3CompatibleHostFlagName string
	if cmdPrefix == "" {
		s3CompatibleHostFlagName = "s3CompatibleHost"
	} else {
		s3CompatibleHostFlagName = fmt.Sprintf("%v.s3CompatibleHost", cmdPrefix)
	}

	var s3CompatibleHostFlagDefault string

	_ = cmd.PersistentFlags().String(s3CompatibleHostFlagName, s3CompatibleHostFlagDefault, s3CompatibleHostDescription)

	return nil
}

func registerBackupRestoreS3SettingsSecretAccessKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secretAccessKeyDescription := `AWS secret access key`

	var secretAccessKeyFlagName string
	if cmdPrefix == "" {
		secretAccessKeyFlagName = "secretAccessKey"
	} else {
		secretAccessKeyFlagName = fmt.Sprintf("%v.secretAccessKey", cmdPrefix)
	}

	var secretAccessKeyFlagDefault string

	_ = cmd.PersistentFlags().String(secretAccessKeyFlagName, secretAccessKeyFlagDefault, secretAccessKeyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBackupRestoreS3SettingsFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accessKeyIdAdded := retrieveBackupRestoreS3SettingsAccessKeyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessKeyIdAdded

	err, bucketNameAdded := retrieveBackupRestoreS3SettingsBucketNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bucketNameAdded

	err, filenameAdded := retrieveBackupRestoreS3SettingsFilenameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || filenameAdded

	err, passwordAdded := retrieveBackupRestoreS3SettingsPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, regionAdded := retrieveBackupRestoreS3SettingsRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionAdded

	err, s3CompatibleHostAdded := retrieveBackupRestoreS3SettingsS3CompatibleHostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || s3CompatibleHostAdded

	err, secretAccessKeyAdded := retrieveBackupRestoreS3SettingsSecretAccessKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretAccessKeyAdded

	return nil, retAdded
}

func retrieveBackupRestoreS3SettingsAccessKeyIDFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessKeyIdFlagName := fmt.Sprintf("%v.accessKeyID", cmdPrefix)
	if cmd.Flags().Changed(accessKeyIdFlagName) {

		var accessKeyIdFlagName string
		if cmdPrefix == "" {
			accessKeyIdFlagName = "accessKeyID"
		} else {
			accessKeyIdFlagName = fmt.Sprintf("%v.accessKeyID", cmdPrefix)
		}

		accessKeyIdFlagValue, err := cmd.Flags().GetString(accessKeyIdFlagName)
		if err != nil {
			return err, false
		}
		m.AccessKeyID = accessKeyIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBackupRestoreS3SettingsBucketNameFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bucketNameFlagName := fmt.Sprintf("%v.bucketName", cmdPrefix)
	if cmd.Flags().Changed(bucketNameFlagName) {

		var bucketNameFlagName string
		if cmdPrefix == "" {
			bucketNameFlagName = "bucketName"
		} else {
			bucketNameFlagName = fmt.Sprintf("%v.bucketName", cmdPrefix)
		}

		bucketNameFlagValue, err := cmd.Flags().GetString(bucketNameFlagName)
		if err != nil {
			return err, false
		}
		m.BucketName = bucketNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBackupRestoreS3SettingsFilenameFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	filenameFlagName := fmt.Sprintf("%v.filename", cmdPrefix)
	if cmd.Flags().Changed(filenameFlagName) {

		var filenameFlagName string
		if cmdPrefix == "" {
			filenameFlagName = "filename"
		} else {
			filenameFlagName = fmt.Sprintf("%v.filename", cmdPrefix)
		}

		filenameFlagValue, err := cmd.Flags().GetString(filenameFlagName)
		if err != nil {
			return err, false
		}
		m.Filename = filenameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBackupRestoreS3SettingsPasswordFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveBackupRestoreS3SettingsRegionFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionFlagName := fmt.Sprintf("%v.region", cmdPrefix)
	if cmd.Flags().Changed(regionFlagName) {

		var regionFlagName string
		if cmdPrefix == "" {
			regionFlagName = "region"
		} else {
			regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
		}

		regionFlagValue, err := cmd.Flags().GetString(regionFlagName)
		if err != nil {
			return err, false
		}
		m.Region = regionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBackupRestoreS3SettingsS3CompatibleHostFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	s3CompatibleHostFlagName := fmt.Sprintf("%v.s3CompatibleHost", cmdPrefix)
	if cmd.Flags().Changed(s3CompatibleHostFlagName) {

		var s3CompatibleHostFlagName string
		if cmdPrefix == "" {
			s3CompatibleHostFlagName = "s3CompatibleHost"
		} else {
			s3CompatibleHostFlagName = fmt.Sprintf("%v.s3CompatibleHost", cmdPrefix)
		}

		s3CompatibleHostFlagValue, err := cmd.Flags().GetString(s3CompatibleHostFlagName)
		if err != nil {
			return err, false
		}
		m.S3CompatibleHost = s3CompatibleHostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBackupRestoreS3SettingsSecretAccessKeyFlags(depth int, m *models.BackupRestoreS3Settings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	secretAccessKeyFlagName := fmt.Sprintf("%v.secretAccessKey", cmdPrefix)
	if cmd.Flags().Changed(secretAccessKeyFlagName) {

		var secretAccessKeyFlagName string
		if cmdPrefix == "" {
			secretAccessKeyFlagName = "secretAccessKey"
		} else {
			secretAccessKeyFlagName = fmt.Sprintf("%v.secretAccessKey", cmdPrefix)
		}

		secretAccessKeyFlagValue, err := cmd.Flags().GetString(secretAccessKeyFlagName)
		if err != nil {
			return err, false
		}
		m.SecretAccessKey = secretAccessKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
