// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/kaas"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationKaasProvisionKaaSClusterCmd returns a cmd to handle operation provisionKaaSCluster
func makeOperationKaasProvisionKaaSClusterCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "provisionKaaSCluster",
		Short: `Provision a new KaaS cluster and create an environment.
**Access policy**: administrator`,
		RunE: runOperationKaasProvisionKaaSCluster,
	}

	if err := registerOperationKaasProvisionKaaSClusterParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationKaasProvisionKaaSCluster uses cmd flags to call endpoint api
func runOperationKaasProvisionKaaSCluster(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := kaas.NewProvisionKaaSClusterParams()
	if err, _ := retrieveOperationKaasProvisionKaaSClusterBodyAmazonFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationKaasProvisionKaaSClusterBodyAPIFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationKaasProvisionKaaSClusterBodyAzureFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationKaasProvisionKaaSClusterBodyGkeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationKaasProvisionKaaSClusterProviderFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationKaasProvisionKaaSClusterResult(appCli.Kaas.ProvisionKaaSCluster(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationKaasProvisionKaaSClusterParamFlags registers all flags needed to fill params
func registerOperationKaasProvisionKaaSClusterParamFlags(cmd *cobra.Command) error {
	if err := registerOperationKaasProvisionKaaSClusterBodyAmazonParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationKaasProvisionKaaSClusterBodyAPIParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationKaasProvisionKaaSClusterBodyAzureParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationKaasProvisionKaaSClusterBodyGkeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationKaasProvisionKaaSClusterProviderParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationKaasProvisionKaaSClusterBodyAmazonParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyAmazonFlagName string
	if cmdPrefix == "" {
		bodyAmazonFlagName = "body_amazon"
	} else {
		bodyAmazonFlagName = fmt.Sprintf("%v.body_amazon", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyAmazonFlagName, "", "Optional json string for [body_amazon]. KaaS cluster provisioning details (required when provider is amazon)")

	// add flags for body
	if err := registerModelProvidersAmazonProvisionPayloadFlags(0, "providersAmazonProvisionPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationKaasProvisionKaaSClusterBodyAPIParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyApiFlagName string
	if cmdPrefix == "" {
		bodyApiFlagName = "body_api"
	} else {
		bodyApiFlagName = fmt.Sprintf("%v.body_api", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyApiFlagName, "", "Optional json string for [body_api]. KaaS cluster provisioning details (required when provider is civo, digitalocean or linode)")

	// add flags for body
	if err := registerModelProvidersDefaultProvisionPayloadFlags(0, "providersDefaultProvisionPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationKaasProvisionKaaSClusterBodyAzureParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyAzureFlagName string
	if cmdPrefix == "" {
		bodyAzureFlagName = "body_azure"
	} else {
		bodyAzureFlagName = fmt.Sprintf("%v.body_azure", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyAzureFlagName, "", "Optional json string for [body_azure]. KaaS cluster provisioning details (required when provider is azure)")

	// add flags for body
	if err := registerModelProvidersAzureProvisionPayloadFlags(0, "providersAzureProvisionPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationKaasProvisionKaaSClusterBodyGkeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyGkeFlagName string
	if cmdPrefix == "" {
		bodyGkeFlagName = "body_gke"
	} else {
		bodyGkeFlagName = fmt.Sprintf("%v.body_gke", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyGkeFlagName, "", "Optional json string for [body_gke]. KaaS cluster provisioning details (required when provider is gke)")

	// add flags for body
	if err := registerModelProvidersGKEProvisionPayloadFlags(0, "providersGKEProvisionPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationKaasProvisionKaaSClusterProviderParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	providerDescription := `Required. Provider`

	var providerFlagName string
	if cmdPrefix == "" {
		providerFlagName = "provider"
	} else {
		providerFlagName = fmt.Sprintf("%v.provider", cmdPrefix)
	}

	var providerFlagDefault int64

	_ = cmd.PersistentFlags().Int64(providerFlagName, providerFlagDefault, providerDescription)

	return nil
}

func retrieveOperationKaasProvisionKaaSClusterBodyAmazonFlag(m *kaas.ProvisionKaaSClusterParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_amazon") {
		// Read body_amazon string from cmd and unmarshal
		bodyAmazonValueStr, err := cmd.Flags().GetString("body_amazon")
		if err != nil {
			return err, false
		}

		bodyAmazonValue := models.ProvidersAmazonProvisionPayload{}
		if err := json.Unmarshal([]byte(bodyAmazonValueStr), &bodyAmazonValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_amazon string in models.ProvidersAmazonProvisionPayload: %v", err), false
		}
		m.BodyAmazon = &bodyAmazonValue
	}
	bodyAmazonValueModel := m.BodyAmazon
	if swag.IsZero(bodyAmazonValueModel) {
		bodyAmazonValueModel = &models.ProvidersAmazonProvisionPayload{}
	}
	err, added := retrieveModelProvidersAmazonProvisionPayloadFlags(0, bodyAmazonValueModel, "providersAmazonProvisionPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyAmazon = bodyAmazonValueModel
	}
	if dryRun && debug {

		bodyAmazonValueDebugBytes, err := json.Marshal(m.BodyAmazon)
		if err != nil {
			return err, false
		}
		logDebugf("BodyAmazon dry-run payload: %v", string(bodyAmazonValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationKaasProvisionKaaSClusterBodyAPIFlag(m *kaas.ProvisionKaaSClusterParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_api") {
		// Read body_api string from cmd and unmarshal
		bodyApiValueStr, err := cmd.Flags().GetString("body_api")
		if err != nil {
			return err, false
		}

		bodyApiValue := models.ProvidersDefaultProvisionPayload{}
		if err := json.Unmarshal([]byte(bodyApiValueStr), &bodyApiValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_api string in models.ProvidersDefaultProvisionPayload: %v", err), false
		}
		m.BodyAPI = &bodyApiValue
	}
	bodyApiValueModel := m.BodyAPI
	if swag.IsZero(bodyApiValueModel) {
		bodyApiValueModel = &models.ProvidersDefaultProvisionPayload{}
	}
	err, added := retrieveModelProvidersDefaultProvisionPayloadFlags(0, bodyApiValueModel, "providersDefaultProvisionPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyAPI = bodyApiValueModel
	}
	if dryRun && debug {

		bodyApiValueDebugBytes, err := json.Marshal(m.BodyAPI)
		if err != nil {
			return err, false
		}
		logDebugf("BodyAPI dry-run payload: %v", string(bodyApiValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationKaasProvisionKaaSClusterBodyAzureFlag(m *kaas.ProvisionKaaSClusterParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_azure") {
		// Read body_azure string from cmd and unmarshal
		bodyAzureValueStr, err := cmd.Flags().GetString("body_azure")
		if err != nil {
			return err, false
		}

		bodyAzureValue := models.ProvidersAzureProvisionPayload{}
		if err := json.Unmarshal([]byte(bodyAzureValueStr), &bodyAzureValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_azure string in models.ProvidersAzureProvisionPayload: %v", err), false
		}
		m.BodyAzure = &bodyAzureValue
	}
	bodyAzureValueModel := m.BodyAzure
	if swag.IsZero(bodyAzureValueModel) {
		bodyAzureValueModel = &models.ProvidersAzureProvisionPayload{}
	}
	err, added := retrieveModelProvidersAzureProvisionPayloadFlags(0, bodyAzureValueModel, "providersAzureProvisionPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyAzure = bodyAzureValueModel
	}
	if dryRun && debug {

		bodyAzureValueDebugBytes, err := json.Marshal(m.BodyAzure)
		if err != nil {
			return err, false
		}
		logDebugf("BodyAzure dry-run payload: %v", string(bodyAzureValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationKaasProvisionKaaSClusterBodyGkeFlag(m *kaas.ProvisionKaaSClusterParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_gke") {
		// Read body_gke string from cmd and unmarshal
		bodyGkeValueStr, err := cmd.Flags().GetString("body_gke")
		if err != nil {
			return err, false
		}

		bodyGkeValue := models.ProvidersGKEProvisionPayload{}
		if err := json.Unmarshal([]byte(bodyGkeValueStr), &bodyGkeValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_gke string in models.ProvidersGKEProvisionPayload: %v", err), false
		}
		m.BodyGke = &bodyGkeValue
	}
	bodyGkeValueModel := m.BodyGke
	if swag.IsZero(bodyGkeValueModel) {
		bodyGkeValueModel = &models.ProvidersGKEProvisionPayload{}
	}
	err, added := retrieveModelProvidersGKEProvisionPayloadFlags(0, bodyGkeValueModel, "providersGKEProvisionPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyGke = bodyGkeValueModel
	}
	if dryRun && debug {

		bodyGkeValueDebugBytes, err := json.Marshal(m.BodyGke)
		if err != nil {
			return err, false
		}
		logDebugf("BodyGke dry-run payload: %v", string(bodyGkeValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationKaasProvisionKaaSClusterProviderFlag(m *kaas.ProvisionKaaSClusterParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("provider") {

		var providerFlagName string
		if cmdPrefix == "" {
			providerFlagName = "provider"
		} else {
			providerFlagName = fmt.Sprintf("%v.provider", cmdPrefix)
		}

		providerFlagValue, err := cmd.Flags().GetInt64(providerFlagName)
		if err != nil {
			return err, false
		}
		m.Provider = providerFlagValue

	}
	return nil, retAdded
}

// parseOperationKaasProvisionKaaSClusterResult parses request result and return the string content
func parseOperationKaasProvisionKaaSClusterResult(resp0 *kaas.ProvisionKaaSClusterOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*kaas.ProvisionKaaSClusterOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning provisionKaaSClusterBadRequest is not supported

		// Non schema case: warning provisionKaaSClusterInternalServerError is not supported

		// Non schema case: warning provisionKaaSClusterServiceUnavailable is not supported

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
