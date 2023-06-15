// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PodsecurityPodSecurityAllowedPaths

// register flags to command
func registerModelPodsecurityPodSecurityAllowedPathsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPodsecurityPodSecurityAllowedPathsPathPrefix(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPodsecurityPodSecurityAllowedPathsReadonly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPodsecurityPodSecurityAllowedPathsPathPrefix(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pathPrefixDescription := ``

	var pathPrefixFlagName string
	if cmdPrefix == "" {
		pathPrefixFlagName = "pathPrefix"
	} else {
		pathPrefixFlagName = fmt.Sprintf("%v.pathPrefix", cmdPrefix)
	}

	var pathPrefixFlagDefault string

	_ = cmd.PersistentFlags().String(pathPrefixFlagName, pathPrefixFlagDefault, pathPrefixDescription)

	return nil
}

func registerPodsecurityPodSecurityAllowedPathsReadonly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	readonlyDescription := ``

	var readonlyFlagName string
	if cmdPrefix == "" {
		readonlyFlagName = "readonly"
	} else {
		readonlyFlagName = fmt.Sprintf("%v.readonly", cmdPrefix)
	}

	var readonlyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(readonlyFlagName, readonlyFlagDefault, readonlyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPodsecurityPodSecurityAllowedPathsFlags(depth int, m *models.PodsecurityPodSecurityAllowedPaths, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, pathPrefixAdded := retrievePodsecurityPodSecurityAllowedPathsPathPrefixFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathPrefixAdded

	err, readonlyAdded := retrievePodsecurityPodSecurityAllowedPathsReadonlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || readonlyAdded

	return nil, retAdded
}

func retrievePodsecurityPodSecurityAllowedPathsPathPrefixFlags(depth int, m *models.PodsecurityPodSecurityAllowedPaths, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pathPrefixFlagName := fmt.Sprintf("%v.pathPrefix", cmdPrefix)
	if cmd.Flags().Changed(pathPrefixFlagName) {

		var pathPrefixFlagName string
		if cmdPrefix == "" {
			pathPrefixFlagName = "pathPrefix"
		} else {
			pathPrefixFlagName = fmt.Sprintf("%v.pathPrefix", cmdPrefix)
		}

		pathPrefixFlagValue, err := cmd.Flags().GetString(pathPrefixFlagName)
		if err != nil {
			return err, false
		}
		m.PathPrefix = pathPrefixFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePodsecurityPodSecurityAllowedPathsReadonlyFlags(depth int, m *models.PodsecurityPodSecurityAllowedPaths, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	readonlyFlagName := fmt.Sprintf("%v.readonly", cmdPrefix)
	if cmd.Flags().Changed(readonlyFlagName) {

		var readonlyFlagName string
		if cmdPrefix == "" {
			readonlyFlagName = "readonly"
		} else {
			readonlyFlagName = fmt.Sprintf("%v.readonly", cmdPrefix)
		}

		readonlyFlagValue, err := cmd.Flags().GetBool(readonlyFlagName)
		if err != nil {
			return err, false
		}
		m.Readonly = &readonlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
