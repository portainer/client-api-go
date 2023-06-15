// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/spf13/cobra"
)

// Schema cli for GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus

// register flags to command
func registerModelGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusDemoEnvironment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusInstanceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Portainer API version`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

func registerGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusDemoEnvironment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var demoEnvironmentFlagName string
	if cmdPrefix == "" {
		demoEnvironmentFlagName = "demoEnvironment"
	} else {
		demoEnvironmentFlagName = fmt.Sprintf("%v.demoEnvironment", cmdPrefix)
	}

	if err := registerModelGithubComPortainerPortainerEeAPIDemoEnvironmentDetailsFlags(depth+1, demoEnvironmentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusInstanceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	instanceIdDescription := `Server Instance ID`

	var instanceIdFlagName string
	if cmdPrefix == "" {
		instanceIdFlagName = "instanceID"
	} else {
		instanceIdFlagName = fmt.Sprintf("%v.instanceID", cmdPrefix)
	}

	var instanceIdFlagDefault string

	_ = cmd.PersistentFlags().String(instanceIdFlagName, instanceIdFlagDefault, instanceIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusFlags(depth int, m *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, versionAdded := retrieveGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	err, demoEnvironmentAdded := retrieveGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusDemoEnvironmentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || demoEnvironmentAdded

	err, instanceIdAdded := retrieveGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusInstanceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || instanceIdAdded

	return nil, retAdded
}

func retrieveGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusVersionFlags(depth int, m *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "Version"
		} else {
			versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusDemoEnvironmentFlags(depth int, m *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	demoEnvironmentFlagName := fmt.Sprintf("%v.demoEnvironment", cmdPrefix)
	if cmd.Flags().Changed(demoEnvironmentFlagName) {
		// info: complex object demoEnvironment GithubComPortainerPortainerEeAPIDemoEnvironmentDetails is retrieved outside this Changed() block
	}
	demoEnvironmentFlagValue := m.DemoEnvironment
	if swag.IsZero(demoEnvironmentFlagValue) {
		demoEnvironmentFlagValue = &models.GithubComPortainerPortainerEeAPIDemoEnvironmentDetails{}
	}

	err, demoEnvironmentAdded := retrieveModelGithubComPortainerPortainerEeAPIDemoEnvironmentDetailsFlags(depth+1, demoEnvironmentFlagValue, demoEnvironmentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || demoEnvironmentAdded
	if demoEnvironmentAdded {
		m.DemoEnvironment = demoEnvironmentFlagValue
	}

	return nil, retAdded
}

func retrieveGithubComPortainerPortainerEeAPIHTTPHandlerSystemStatusInstanceIDFlags(depth int, m *models.GithubComPortainerPortainerEeAPIHTTPHandlerSystemStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	instanceIdFlagName := fmt.Sprintf("%v.instanceID", cmdPrefix)
	if cmd.Flags().Changed(instanceIdFlagName) {

		var instanceIdFlagName string
		if cmdPrefix == "" {
			instanceIdFlagName = "instanceID"
		} else {
			instanceIdFlagName = fmt.Sprintf("%v.instanceID", cmdPrefix)
		}

		instanceIdFlagValue, err := cmd.Flags().GetString(instanceIdFlagName)
		if err != nil {
			return err, false
		}
		m.InstanceID = instanceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
