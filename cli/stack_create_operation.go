// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/portainer/client-api-go/v2/client/stacks"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationStacksStackCreateCmd returns a cmd to handle operation stackCreate
func makeOperationStacksStackCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "StackCreate",
		Short: `Deploy a new stack into a Docker environment(endpoint) specified via the environment(endpoint) identifier.
**Access policy**: authenticated`,
		RunE: runOperationStacksStackCreate,
	}

	if err := registerOperationStacksStackCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationStacksStackCreate uses cmd flags to call endpoint api
func runOperationStacksStackCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := stacks.NewStackCreateParams()
	if err, _ := retrieveOperationStacksStackCreateEnvFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateSwarmIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodyComposeRepositoryFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodyComposeStringFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodyKubernetesRepositoryFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodyKubernetesStringFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodyKubernetesURLFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodySwarmRepositoryFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateBodySwarmStringFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateEndpointIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateFileFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateMethodFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationStacksStackCreateTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationStacksStackCreateResult(appCli.Stacks.StackCreate(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationStacksStackCreateParamFlags registers all flags needed to fill params
func registerOperationStacksStackCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationStacksStackCreateEnvParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateSwarmIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodyComposeRepositoryParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodyComposeStringParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodyKubernetesRepositoryParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodyKubernetesStringParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodyKubernetesURLParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodySwarmRepositoryParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateBodySwarmStringParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateEndpointIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateFileParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateMethodParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationStacksStackCreateTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationStacksStackCreateEnvParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	envDescription := `Environment(Endpoint) variables passed during deployment, represented as a JSON array [{'name': 'name', 'value': 'value'}]. Optional, used when method equals file and type equals 1.`

	var envFlagName string
	if cmdPrefix == "" {
		envFlagName = "Env"
	} else {
		envFlagName = fmt.Sprintf("%v.Env", cmdPrefix)
	}

	var envFlagDefault string

	_ = cmd.PersistentFlags().String(envFlagName, envFlagDefault, envDescription)

	return nil
}
func registerOperationStacksStackCreateNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Name of the stack. required when method is file`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "Name"
	} else {
		nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}
func registerOperationStacksStackCreateSwarmIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	swarmIdDescription := `Swarm cluster identifier. Required when method equals file and type equals 1. required when method is file`

	var swarmIdFlagName string
	if cmdPrefix == "" {
		swarmIdFlagName = "SwarmID"
	} else {
		swarmIdFlagName = fmt.Sprintf("%v.SwarmID", cmdPrefix)
	}

	var swarmIdFlagDefault string

	_ = cmd.PersistentFlags().String(swarmIdFlagName, swarmIdFlagDefault, swarmIdDescription)

	return nil
}
func registerOperationStacksStackCreateBodyComposeRepositoryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyComposeRepositoryFlagName string
	if cmdPrefix == "" {
		bodyComposeRepositoryFlagName = "body_compose_repository"
	} else {
		bodyComposeRepositoryFlagName = fmt.Sprintf("%v.body_compose_repository", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyComposeRepositoryFlagName, "", "Optional json string for [body_compose_repository]. Required when using method=repository and type=2")

	// add flags for body
	if err := registerModelStacksComposeStackFromGitRepositoryPayloadFlags(0, "stacksComposeStackFromGitRepositoryPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateBodyComposeStringParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyComposeStringFlagName string
	if cmdPrefix == "" {
		bodyComposeStringFlagName = "body_compose_string"
	} else {
		bodyComposeStringFlagName = fmt.Sprintf("%v.body_compose_string", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyComposeStringFlagName, "", "Optional json string for [body_compose_string]. Required when using method=string and type=2")

	// add flags for body
	if err := registerModelStacksComposeStackFromFileContentPayloadFlags(0, "stacksComposeStackFromFileContentPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateBodyKubernetesRepositoryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyKubernetesRepositoryFlagName string
	if cmdPrefix == "" {
		bodyKubernetesRepositoryFlagName = "body_kubernetes_repository"
	} else {
		bodyKubernetesRepositoryFlagName = fmt.Sprintf("%v.body_kubernetes_repository", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyKubernetesRepositoryFlagName, "", "Optional json string for [body_kubernetes_repository]. Required when using method=repository and type=3")

	// add flags for body
	if err := registerModelStacksKubernetesGitDeploymentPayloadFlags(0, "stacksKubernetesGitDeploymentPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateBodyKubernetesStringParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyKubernetesStringFlagName string
	if cmdPrefix == "" {
		bodyKubernetesStringFlagName = "body_kubernetes_string"
	} else {
		bodyKubernetesStringFlagName = fmt.Sprintf("%v.body_kubernetes_string", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyKubernetesStringFlagName, "", "Optional json string for [body_kubernetes_string]. Required when using method=string and type=3")

	// add flags for body
	if err := registerModelStacksKubernetesStringDeploymentPayloadFlags(0, "stacksKubernetesStringDeploymentPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateBodyKubernetesURLParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyKubernetesUrlFlagName string
	if cmdPrefix == "" {
		bodyKubernetesUrlFlagName = "body_kubernetes_url"
	} else {
		bodyKubernetesUrlFlagName = fmt.Sprintf("%v.body_kubernetes_url", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyKubernetesUrlFlagName, "", "Optional json string for [body_kubernetes_url]. Required when using method=url and type=3")

	// add flags for body
	if err := registerModelStacksKubernetesManifestURLDeploymentPayloadFlags(0, "stacksKubernetesManifestUrlDeploymentPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateBodySwarmRepositoryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodySwarmRepositoryFlagName string
	if cmdPrefix == "" {
		bodySwarmRepositoryFlagName = "body_swarm_repository"
	} else {
		bodySwarmRepositoryFlagName = fmt.Sprintf("%v.body_swarm_repository", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodySwarmRepositoryFlagName, "", "Optional json string for [body_swarm_repository]. Required when using method=repository and type=1")

	// add flags for body
	if err := registerModelStacksSwarmStackFromGitRepositoryPayloadFlags(0, "stacksSwarmStackFromGitRepositoryPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateBodySwarmStringParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodySwarmStringFlagName string
	if cmdPrefix == "" {
		bodySwarmStringFlagName = "body_swarm_string"
	} else {
		bodySwarmStringFlagName = fmt.Sprintf("%v.body_swarm_string", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodySwarmStringFlagName, "", "Optional json string for [body_swarm_string]. Required when using method=string and type=1")

	// add flags for body
	if err := registerModelStacksSwarmStackFromFileContentPayloadFlags(0, "stacksSwarmStackFromFileContentPayload", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateEndpointIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	endpointIdDescription := `Required. Identifier of the environment(endpoint) that will be used to deploy the stack`

	var endpointIdFlagName string
	if cmdPrefix == "" {
		endpointIdFlagName = "endpointId"
	} else {
		endpointIdFlagName = fmt.Sprintf("%v.endpointId", cmdPrefix)
	}

	var endpointIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(endpointIdFlagName, endpointIdFlagDefault, endpointIdDescription)

	return nil
}
func registerOperationStacksStackCreateFileParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	// warning: primitive file io.ReadCloser is not supported by go-swagger cli yet

	return nil
}
func registerOperationStacksStackCreateMethodParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	methodDescription := `Enum: ["string","file","repository","url"]. Required. Stack deployment method. Possible values: file, string, repository or url.`

	var methodFlagName string
	if cmdPrefix == "" {
		methodFlagName = "method"
	} else {
		methodFlagName = fmt.Sprintf("%v.method", cmdPrefix)
	}

	var methodFlagDefault string

	_ = cmd.PersistentFlags().String(methodFlagName, methodFlagDefault, methodDescription)

	if err := cmd.RegisterFlagCompletionFunc(methodFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["string","file","repository","url"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationStacksStackCreateTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `Enum: [1,2,3]. Required. Stack deployment type. Possible values: 1 (Swarm stack), 2 (Compose stack) or 3 (Kubernetes stack).`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationStacksStackCreateEnvFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("Env") {

		var envFlagName string
		if cmdPrefix == "" {
			envFlagName = "Env"
		} else {
			envFlagName = fmt.Sprintf("%v.Env", cmdPrefix)
		}

		envFlagValue, err := cmd.Flags().GetString(envFlagName)
		if err != nil {
			return err, false
		}
		m.Env = &envFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStacksStackCreateNameFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("Name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "Name"
		} else {
			nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStacksStackCreateSwarmIDFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("SwarmID") {

		var swarmIdFlagName string
		if cmdPrefix == "" {
			swarmIdFlagName = "SwarmID"
		} else {
			swarmIdFlagName = fmt.Sprintf("%v.SwarmID", cmdPrefix)
		}

		swarmIdFlagValue, err := cmd.Flags().GetString(swarmIdFlagName)
		if err != nil {
			return err, false
		}
		m.SwarmID = &swarmIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodyComposeRepositoryFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_compose_repository") {
		// Read body_compose_repository string from cmd and unmarshal
		bodyComposeRepositoryValueStr, err := cmd.Flags().GetString("body_compose_repository")
		if err != nil {
			return err, false
		}

		bodyComposeRepositoryValue := models.StacksComposeStackFromGitRepositoryPayload{}
		if err := json.Unmarshal([]byte(bodyComposeRepositoryValueStr), &bodyComposeRepositoryValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_compose_repository string in models.StacksComposeStackFromGitRepositoryPayload: %v", err), false
		}
		m.BodyComposeRepository = &bodyComposeRepositoryValue
	}
	bodyComposeRepositoryValueModel := m.BodyComposeRepository
	if swag.IsZero(bodyComposeRepositoryValueModel) {
		bodyComposeRepositoryValueModel = &models.StacksComposeStackFromGitRepositoryPayload{}
	}
	err, added := retrieveModelStacksComposeStackFromGitRepositoryPayloadFlags(0, bodyComposeRepositoryValueModel, "stacksComposeStackFromGitRepositoryPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyComposeRepository = bodyComposeRepositoryValueModel
	}
	if dryRun && debug {

		bodyComposeRepositoryValueDebugBytes, err := json.Marshal(m.BodyComposeRepository)
		if err != nil {
			return err, false
		}
		logDebugf("BodyComposeRepository dry-run payload: %v", string(bodyComposeRepositoryValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodyComposeStringFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_compose_string") {
		// Read body_compose_string string from cmd and unmarshal
		bodyComposeStringValueStr, err := cmd.Flags().GetString("body_compose_string")
		if err != nil {
			return err, false
		}

		bodyComposeStringValue := models.StacksComposeStackFromFileContentPayload{}
		if err := json.Unmarshal([]byte(bodyComposeStringValueStr), &bodyComposeStringValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_compose_string string in models.StacksComposeStackFromFileContentPayload: %v", err), false
		}
		m.BodyComposeString = &bodyComposeStringValue
	}
	bodyComposeStringValueModel := m.BodyComposeString
	if swag.IsZero(bodyComposeStringValueModel) {
		bodyComposeStringValueModel = &models.StacksComposeStackFromFileContentPayload{}
	}
	err, added := retrieveModelStacksComposeStackFromFileContentPayloadFlags(0, bodyComposeStringValueModel, "stacksComposeStackFromFileContentPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyComposeString = bodyComposeStringValueModel
	}
	if dryRun && debug {

		bodyComposeStringValueDebugBytes, err := json.Marshal(m.BodyComposeString)
		if err != nil {
			return err, false
		}
		logDebugf("BodyComposeString dry-run payload: %v", string(bodyComposeStringValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodyKubernetesRepositoryFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_kubernetes_repository") {
		// Read body_kubernetes_repository string from cmd and unmarshal
		bodyKubernetesRepositoryValueStr, err := cmd.Flags().GetString("body_kubernetes_repository")
		if err != nil {
			return err, false
		}

		bodyKubernetesRepositoryValue := models.StacksKubernetesGitDeploymentPayload{}
		if err := json.Unmarshal([]byte(bodyKubernetesRepositoryValueStr), &bodyKubernetesRepositoryValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_kubernetes_repository string in models.StacksKubernetesGitDeploymentPayload: %v", err), false
		}
		m.BodyKubernetesRepository = &bodyKubernetesRepositoryValue
	}
	bodyKubernetesRepositoryValueModel := m.BodyKubernetesRepository
	if swag.IsZero(bodyKubernetesRepositoryValueModel) {
		bodyKubernetesRepositoryValueModel = &models.StacksKubernetesGitDeploymentPayload{}
	}
	err, added := retrieveModelStacksKubernetesGitDeploymentPayloadFlags(0, bodyKubernetesRepositoryValueModel, "stacksKubernetesGitDeploymentPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyKubernetesRepository = bodyKubernetesRepositoryValueModel
	}
	if dryRun && debug {

		bodyKubernetesRepositoryValueDebugBytes, err := json.Marshal(m.BodyKubernetesRepository)
		if err != nil {
			return err, false
		}
		logDebugf("BodyKubernetesRepository dry-run payload: %v", string(bodyKubernetesRepositoryValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodyKubernetesStringFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_kubernetes_string") {
		// Read body_kubernetes_string string from cmd and unmarshal
		bodyKubernetesStringValueStr, err := cmd.Flags().GetString("body_kubernetes_string")
		if err != nil {
			return err, false
		}

		bodyKubernetesStringValue := models.StacksKubernetesStringDeploymentPayload{}
		if err := json.Unmarshal([]byte(bodyKubernetesStringValueStr), &bodyKubernetesStringValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_kubernetes_string string in models.StacksKubernetesStringDeploymentPayload: %v", err), false
		}
		m.BodyKubernetesString = &bodyKubernetesStringValue
	}
	bodyKubernetesStringValueModel := m.BodyKubernetesString
	if swag.IsZero(bodyKubernetesStringValueModel) {
		bodyKubernetesStringValueModel = &models.StacksKubernetesStringDeploymentPayload{}
	}
	err, added := retrieveModelStacksKubernetesStringDeploymentPayloadFlags(0, bodyKubernetesStringValueModel, "stacksKubernetesStringDeploymentPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyKubernetesString = bodyKubernetesStringValueModel
	}
	if dryRun && debug {

		bodyKubernetesStringValueDebugBytes, err := json.Marshal(m.BodyKubernetesString)
		if err != nil {
			return err, false
		}
		logDebugf("BodyKubernetesString dry-run payload: %v", string(bodyKubernetesStringValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodyKubernetesURLFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_kubernetes_url") {
		// Read body_kubernetes_url string from cmd and unmarshal
		bodyKubernetesUrlValueStr, err := cmd.Flags().GetString("body_kubernetes_url")
		if err != nil {
			return err, false
		}

		bodyKubernetesUrlValue := models.StacksKubernetesManifestURLDeploymentPayload{}
		if err := json.Unmarshal([]byte(bodyKubernetesUrlValueStr), &bodyKubernetesUrlValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_kubernetes_url string in models.StacksKubernetesManifestURLDeploymentPayload: %v", err), false
		}
		m.BodyKubernetesURL = &bodyKubernetesUrlValue
	}
	bodyKubernetesUrlValueModel := m.BodyKubernetesURL
	if swag.IsZero(bodyKubernetesUrlValueModel) {
		bodyKubernetesUrlValueModel = &models.StacksKubernetesManifestURLDeploymentPayload{}
	}
	err, added := retrieveModelStacksKubernetesManifestURLDeploymentPayloadFlags(0, bodyKubernetesUrlValueModel, "stacksKubernetesManifestUrlDeploymentPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodyKubernetesURL = bodyKubernetesUrlValueModel
	}
	if dryRun && debug {

		bodyKubernetesUrlValueDebugBytes, err := json.Marshal(m.BodyKubernetesURL)
		if err != nil {
			return err, false
		}
		logDebugf("BodyKubernetesURL dry-run payload: %v", string(bodyKubernetesUrlValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodySwarmRepositoryFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_swarm_repository") {
		// Read body_swarm_repository string from cmd and unmarshal
		bodySwarmRepositoryValueStr, err := cmd.Flags().GetString("body_swarm_repository")
		if err != nil {
			return err, false
		}

		bodySwarmRepositoryValue := models.StacksSwarmStackFromGitRepositoryPayload{}
		if err := json.Unmarshal([]byte(bodySwarmRepositoryValueStr), &bodySwarmRepositoryValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_swarm_repository string in models.StacksSwarmStackFromGitRepositoryPayload: %v", err), false
		}
		m.BodySwarmRepository = &bodySwarmRepositoryValue
	}
	bodySwarmRepositoryValueModel := m.BodySwarmRepository
	if swag.IsZero(bodySwarmRepositoryValueModel) {
		bodySwarmRepositoryValueModel = &models.StacksSwarmStackFromGitRepositoryPayload{}
	}
	err, added := retrieveModelStacksSwarmStackFromGitRepositoryPayloadFlags(0, bodySwarmRepositoryValueModel, "stacksSwarmStackFromGitRepositoryPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodySwarmRepository = bodySwarmRepositoryValueModel
	}
	if dryRun && debug {

		bodySwarmRepositoryValueDebugBytes, err := json.Marshal(m.BodySwarmRepository)
		if err != nil {
			return err, false
		}
		logDebugf("BodySwarmRepository dry-run payload: %v", string(bodySwarmRepositoryValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateBodySwarmStringFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body_swarm_string") {
		// Read body_swarm_string string from cmd and unmarshal
		bodySwarmStringValueStr, err := cmd.Flags().GetString("body_swarm_string")
		if err != nil {
			return err, false
		}

		bodySwarmStringValue := models.StacksSwarmStackFromFileContentPayload{}
		if err := json.Unmarshal([]byte(bodySwarmStringValueStr), &bodySwarmStringValue); err != nil {
			return fmt.Errorf("cannot unmarshal body_swarm_string string in models.StacksSwarmStackFromFileContentPayload: %v", err), false
		}
		m.BodySwarmString = &bodySwarmStringValue
	}
	bodySwarmStringValueModel := m.BodySwarmString
	if swag.IsZero(bodySwarmStringValueModel) {
		bodySwarmStringValueModel = &models.StacksSwarmStackFromFileContentPayload{}
	}
	err, added := retrieveModelStacksSwarmStackFromFileContentPayloadFlags(0, bodySwarmStringValueModel, "stacksSwarmStackFromFileContentPayload", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.BodySwarmString = bodySwarmStringValueModel
	}
	if dryRun && debug {

		bodySwarmStringValueDebugBytes, err := json.Marshal(m.BodySwarmString)
		if err != nil {
			return err, false
		}
		logDebugf("BodySwarmString dry-run payload: %v", string(bodySwarmStringValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationStacksStackCreateEndpointIDFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("endpointId") {

		var endpointIdFlagName string
		if cmdPrefix == "" {
			endpointIdFlagName = "endpointId"
		} else {
			endpointIdFlagName = fmt.Sprintf("%v.endpointId", cmdPrefix)
		}

		endpointIdFlagValue, err := cmd.Flags().GetInt64(endpointIdFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointID = endpointIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStacksStackCreateFileFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("file") {

		// warning: primitive file io.ReadCloser is not supported by go-swagger cli yet

	}
	return nil, retAdded
}
func retrieveOperationStacksStackCreateMethodFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("method") {

		var methodFlagName string
		if cmdPrefix == "" {
			methodFlagName = "method"
		} else {
			methodFlagName = fmt.Sprintf("%v.method", cmdPrefix)
		}

		methodFlagValue, err := cmd.Flags().GetString(methodFlagName)
		if err != nil {
			return err, false
		}
		m.Method = methodFlagValue

	}
	return nil, retAdded
}
func retrieveOperationStacksStackCreateTypeFlag(m *stacks.StackCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type") {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetInt64(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

	}
	return nil, retAdded
}

// parseOperationStacksStackCreateResult parses request result and return the string content
func parseOperationStacksStackCreateResult(resp0 *stacks.StackCreateOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*stacks.StackCreateOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning stackCreateBadRequest is not supported

		// Non schema case: warning stackCreateInternalServerError is not supported

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
