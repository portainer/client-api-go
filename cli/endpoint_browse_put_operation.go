// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/client/endpoints"

	"github.com/spf13/cobra"
)

// makeOperationEndpointsEndpointBrowsePutCmd returns a cmd to handle operation endpointBrowsePut
func makeOperationEndpointsEndpointBrowsePutCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "EndpointBrowsePut",
		Short: `Upload a file under a specific path on the file system of an environment (endpoint).
**Access policy**: authenticated`,
		RunE: runOperationEndpointsEndpointBrowsePut,
	}

	if err := registerOperationEndpointsEndpointBrowsePutParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEndpointsEndpointBrowsePut uses cmd flags to call endpoint api
func runOperationEndpointsEndpointBrowsePut(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := endpoints.NewEndpointBrowsePutParams()
	if err, _ := retrieveOperationEndpointsEndpointBrowsePutFileFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEndpointsEndpointBrowsePutIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEndpointsEndpointBrowsePutPathFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEndpointsEndpointBrowsePutVolumeIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEndpointsEndpointBrowsePutResult(appCli.Endpoints.EndpointBrowsePut(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEndpointsEndpointBrowsePutParamFlags registers all flags needed to fill params
func registerOperationEndpointsEndpointBrowsePutParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEndpointsEndpointBrowsePutFileParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEndpointsEndpointBrowsePutIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEndpointsEndpointBrowsePutPathParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEndpointsEndpointBrowsePutVolumeIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEndpointsEndpointBrowsePutFileParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	// warning: primitive file io.ReadCloser is not supported by go-swagger cli yet

	return nil
}
func registerOperationEndpointsEndpointBrowsePutIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationEndpointsEndpointBrowsePutPathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pathDescription := `Required. Path to upload the file`

	var pathFlagName string
	if cmdPrefix == "" {
		pathFlagName = "path"
	} else {
		pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var pathFlagDefault string

	_ = cmd.PersistentFlags().String(pathFlagName, pathFlagDefault, pathDescription)

	return nil
}
func registerOperationEndpointsEndpointBrowsePutVolumeIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	volumeIdDescription := `Optional volume identifier to upload the file`

	var volumeIdFlagName string
	if cmdPrefix == "" {
		volumeIdFlagName = "volumeID"
	} else {
		volumeIdFlagName = fmt.Sprintf("%v.volumeID", cmdPrefix)
	}

	var volumeIdFlagDefault string

	_ = cmd.PersistentFlags().String(volumeIdFlagName, volumeIdFlagDefault, volumeIdDescription)

	return nil
}

func retrieveOperationEndpointsEndpointBrowsePutFileFlag(m *endpoints.EndpointBrowsePutParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("file") {

		// warning: primitive file io.ReadCloser is not supported by go-swagger cli yet

	}
	return nil, retAdded
}
func retrieveOperationEndpointsEndpointBrowsePutIDFlag(m *endpoints.EndpointBrowsePutParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationEndpointsEndpointBrowsePutPathFlag(m *endpoints.EndpointBrowsePutParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("path") {

		var pathFlagName string
		if cmdPrefix == "" {
			pathFlagName = "path"
		} else {
			pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		pathFlagValue, err := cmd.Flags().GetString(pathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = pathFlagValue

	}
	return nil, retAdded
}
func retrieveOperationEndpointsEndpointBrowsePutVolumeIDFlag(m *endpoints.EndpointBrowsePutParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("volumeID") {

		var volumeIdFlagName string
		if cmdPrefix == "" {
			volumeIdFlagName = "volumeID"
		} else {
			volumeIdFlagName = fmt.Sprintf("%v.volumeID", cmdPrefix)
		}

		volumeIdFlagValue, err := cmd.Flags().GetString(volumeIdFlagName)
		if err != nil {
			return err, false
		}
		m.VolumeID = &volumeIdFlagValue

	}
	return nil, retAdded
}

// parseOperationEndpointsEndpointBrowsePutResult parses request result and return the string content
func parseOperationEndpointsEndpointBrowsePutResult(resp0 *endpoints.EndpointBrowsePutOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning endpointBrowsePutOK is not supported

		// Non schema case: warning endpointBrowsePutBadRequest is not supported

		// Non schema case: warning endpointBrowsePutInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response endpointBrowsePutOK is not supported by go-swagger cli yet.

	return "", nil
}
