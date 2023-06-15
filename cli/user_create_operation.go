// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/users"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUsersUserCreateCmd returns a cmd to handle operation userCreate
func makeOperationUsersUserCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "UserCreate",
		Short: `Create a new Portainer user.
Only administrators can create users.
**Access policy**: restricted`,
		RunE: runOperationUsersUserCreate,
	}

	if err := registerOperationUsersUserCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUsersUserCreate uses cmd flags to call endpoint api
func runOperationUsersUserCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := users.NewUserCreateParams()
	if err, _ := retrieveOperationUsersUserCreateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUsersUserCreateResult(appCli.Users.UserCreate(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUsersUserCreateParamFlags registers all flags needed to fill params
func registerOperationUsersUserCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUsersUserCreateBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUsersUserCreateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. User details")

	// add flags for body
	if err := registerModelUsersUserCreatePayloadFlags(0, "usersUserCreatePayload", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationUsersUserCreateBodyFlag(m *users.UserCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.UsersUserCreatePayload{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.UsersUserCreatePayload: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.UsersUserCreatePayload{}
	}
	err, added := retrieveModelUsersUserCreatePayloadFlags(0, bodyValueModel, "usersUserCreatePayload", cmd)
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

// parseOperationUsersUserCreateResult parses request result and return the string content
func parseOperationUsersUserCreateResult(resp0 *users.UserCreateOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*users.UserCreateOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning userCreateBadRequest is not supported

		// Non schema case: warning userCreateForbidden is not supported

		// Non schema case: warning userCreateConflict is not supported

		// Non schema case: warning userCreateInternalServerError is not supported

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
