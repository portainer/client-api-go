// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/portainer/client-api-go/v2/models"

	"github.com/spf13/cobra"
)

// Schema cli for PortainereeRegistry

// register flags to command
func registerModelPortainereeRegistryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeRegistryAccessToken(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryAccessTokenExpiry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryAuthentication(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryAuthorizedTeams(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryAuthorizedUsers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryBaseURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryEcr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryGithub(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryGitlab(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryManagementConfiguration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryQuay(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryRegistryAccesses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryTeamAccessPolicies(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryUserAccessPolicies(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeRegistryUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryAccessToken(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessTokenDescription := `Stores temporary access token`

	var accessTokenFlagName string
	if cmdPrefix == "" {
		accessTokenFlagName = "AccessToken"
	} else {
		accessTokenFlagName = fmt.Sprintf("%v.AccessToken", cmdPrefix)
	}

	var accessTokenFlagDefault string

	_ = cmd.PersistentFlags().String(accessTokenFlagName, accessTokenFlagDefault, accessTokenDescription)

	return nil
}

func registerPortainereeRegistryAccessTokenExpiry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessTokenExpiryDescription := ``

	var accessTokenExpiryFlagName string
	if cmdPrefix == "" {
		accessTokenExpiryFlagName = "AccessTokenExpiry"
	} else {
		accessTokenExpiryFlagName = fmt.Sprintf("%v.AccessTokenExpiry", cmdPrefix)
	}

	var accessTokenExpiryFlagDefault int64

	_ = cmd.PersistentFlags().Int64(accessTokenExpiryFlagName, accessTokenExpiryFlagDefault, accessTokenExpiryDescription)

	return nil
}

func registerPortainereeRegistryAuthentication(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authenticationDescription := `Is authentication against this registry enabled`

	var authenticationFlagName string
	if cmdPrefix == "" {
		authenticationFlagName = "Authentication"
	} else {
		authenticationFlagName = fmt.Sprintf("%v.Authentication", cmdPrefix)
	}

	var authenticationFlagDefault bool

	_ = cmd.PersistentFlags().Bool(authenticationFlagName, authenticationFlagDefault, authenticationDescription)

	return nil
}

func registerPortainereeRegistryAuthorizedTeams(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: AuthorizedTeams []int64 array type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeRegistryAuthorizedUsers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: AuthorizedUsers []int64 array type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeRegistryBaseURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	baseUrlDescription := `Base URL, introduced for ProGet registry`

	var baseUrlFlagName string
	if cmdPrefix == "" {
		baseUrlFlagName = "BaseURL"
	} else {
		baseUrlFlagName = fmt.Sprintf("%v.BaseURL", cmdPrefix)
	}

	var baseUrlFlagDefault string

	_ = cmd.PersistentFlags().String(baseUrlFlagName, baseUrlFlagDefault, baseUrlDescription)

	return nil
}

func registerPortainereeRegistryEcr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ecrFlagName string
	if cmdPrefix == "" {
		ecrFlagName = "Ecr"
	} else {
		ecrFlagName = fmt.Sprintf("%v.Ecr", cmdPrefix)
	}

	if err := registerModelPortainereeEcrDataFlags(depth+1, ecrFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryGithub(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var githubFlagName string
	if cmdPrefix == "" {
		githubFlagName = "Github"
	} else {
		githubFlagName = fmt.Sprintf("%v.Github", cmdPrefix)
	}

	if err := registerModelPortainereeGithubRegistryDataFlags(depth+1, githubFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryGitlab(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var gitlabFlagName string
	if cmdPrefix == "" {
		gitlabFlagName = "Gitlab"
	} else {
		gitlabFlagName = fmt.Sprintf("%v.Gitlab", cmdPrefix)
	}

	if err := registerModelPortainereeGitlabRegistryDataFlags(depth+1, gitlabFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Registry Identifier`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "Id"
	} else {
		idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerPortainereeRegistryManagementConfiguration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var managementConfigurationFlagName string
	if cmdPrefix == "" {
		managementConfigurationFlagName = "ManagementConfiguration"
	} else {
		managementConfigurationFlagName = fmt.Sprintf("%v.ManagementConfiguration", cmdPrefix)
	}

	if err := registerModelPortainereeRegistryManagementConfigurationFlags(depth+1, managementConfigurationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Registry Name`

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

func registerPortainereeRegistryPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Password or SecretAccessKey used to authenticate against this registry`

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "Password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.Password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

func registerPortainereeRegistryQuay(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var quayFlagName string
	if cmdPrefix == "" {
		quayFlagName = "Quay"
	} else {
		quayFlagName = fmt.Sprintf("%v.Quay", cmdPrefix)
	}

	if err := registerModelPortainereeQuayRegistryDataFlags(depth+1, quayFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryRegistryAccesses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: RegistryAccesses PortainereeRegistryAccesses map type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeRegistryTeamAccessPolicies(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: TeamAccessPolicies PortainereeTeamAccessPolicies map type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeRegistryType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: [1,2,3,4,5,6,7,8]. Registry Type (1 - Quay, 2 - Azure, 3 - Custom, 4 - Gitlab, 5 - ProGet, 6 - DockerHub, 7 - ECR, 8 - Github)`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "Type"
	} else {
		typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
	}

	var typeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerPortainereeRegistryURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := `URL or IP address of the Docker registry`

	var urlFlagName string
	if cmdPrefix == "" {
		urlFlagName = "URL"
	} else {
		urlFlagName = fmt.Sprintf("%v.URL", cmdPrefix)
	}

	var urlFlagDefault string

	_ = cmd.PersistentFlags().String(urlFlagName, urlFlagDefault, urlDescription)

	return nil
}

func registerPortainereeRegistryUserAccessPolicies(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: UserAccessPolicies PortainereeUserAccessPolicies map type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeRegistryUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := `Username or AccessKeyID used to authenticate against this registry`

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "Username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.Username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereeRegistryFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accessTokenAdded := retrievePortainereeRegistryAccessTokenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessTokenAdded

	err, accessTokenExpiryAdded := retrievePortainereeRegistryAccessTokenExpiryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessTokenExpiryAdded

	err, authenticationAdded := retrievePortainereeRegistryAuthenticationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authenticationAdded

	err, authorizedTeamsAdded := retrievePortainereeRegistryAuthorizedTeamsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorizedTeamsAdded

	err, authorizedUsersAdded := retrievePortainereeRegistryAuthorizedUsersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorizedUsersAdded

	err, baseUrlAdded := retrievePortainereeRegistryBaseURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || baseUrlAdded

	err, ecrAdded := retrievePortainereeRegistryEcrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ecrAdded

	err, githubAdded := retrievePortainereeRegistryGithubFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || githubAdded

	err, gitlabAdded := retrievePortainereeRegistryGitlabFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gitlabAdded

	err, idAdded := retrievePortainereeRegistryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, managementConfigurationAdded := retrievePortainereeRegistryManagementConfigurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || managementConfigurationAdded

	err, nameAdded := retrievePortainereeRegistryNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, passwordAdded := retrievePortainereeRegistryPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, quayAdded := retrievePortainereeRegistryQuayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quayAdded

	err, registryAccessesAdded := retrievePortainereeRegistryRegistryAccessesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || registryAccessesAdded

	err, teamAccessPoliciesAdded := retrievePortainereeRegistryTeamAccessPoliciesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || teamAccessPoliciesAdded

	err, typeAdded := retrievePortainereeRegistryTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, urlAdded := retrievePortainereeRegistryURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	err, userAccessPoliciesAdded := retrievePortainereeRegistryUserAccessPoliciesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAccessPoliciesAdded

	err, usernameAdded := retrievePortainereeRegistryUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrievePortainereeRegistryAccessTokenFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessTokenFlagName := fmt.Sprintf("%v.AccessToken", cmdPrefix)
	if cmd.Flags().Changed(accessTokenFlagName) {

		var accessTokenFlagName string
		if cmdPrefix == "" {
			accessTokenFlagName = "AccessToken"
		} else {
			accessTokenFlagName = fmt.Sprintf("%v.AccessToken", cmdPrefix)
		}

		accessTokenFlagValue, err := cmd.Flags().GetString(accessTokenFlagName)
		if err != nil {
			return err, false
		}
		m.AccessToken = accessTokenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryAccessTokenExpiryFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessTokenExpiryFlagName := fmt.Sprintf("%v.AccessTokenExpiry", cmdPrefix)
	if cmd.Flags().Changed(accessTokenExpiryFlagName) {

		var accessTokenExpiryFlagName string
		if cmdPrefix == "" {
			accessTokenExpiryFlagName = "AccessTokenExpiry"
		} else {
			accessTokenExpiryFlagName = fmt.Sprintf("%v.AccessTokenExpiry", cmdPrefix)
		}

		accessTokenExpiryFlagValue, err := cmd.Flags().GetInt64(accessTokenExpiryFlagName)
		if err != nil {
			return err, false
		}
		m.AccessTokenExpiry = accessTokenExpiryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryAuthenticationFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authenticationFlagName := fmt.Sprintf("%v.Authentication", cmdPrefix)
	if cmd.Flags().Changed(authenticationFlagName) {

		var authenticationFlagName string
		if cmdPrefix == "" {
			authenticationFlagName = "Authentication"
		} else {
			authenticationFlagName = fmt.Sprintf("%v.Authentication", cmdPrefix)
		}

		authenticationFlagValue, err := cmd.Flags().GetBool(authenticationFlagName)
		if err != nil {
			return err, false
		}
		m.Authentication = &authenticationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryAuthorizedTeamsFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authorizedTeamsFlagName := fmt.Sprintf("%v.AuthorizedTeams", cmdPrefix)
	if cmd.Flags().Changed(authorizedTeamsFlagName) {
		// warning: AuthorizedTeams array type []int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeRegistryAuthorizedUsersFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authorizedUsersFlagName := fmt.Sprintf("%v.AuthorizedUsers", cmdPrefix)
	if cmd.Flags().Changed(authorizedUsersFlagName) {
		// warning: AuthorizedUsers array type []int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeRegistryBaseURLFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	baseUrlFlagName := fmt.Sprintf("%v.BaseURL", cmdPrefix)
	if cmd.Flags().Changed(baseUrlFlagName) {

		var baseUrlFlagName string
		if cmdPrefix == "" {
			baseUrlFlagName = "BaseURL"
		} else {
			baseUrlFlagName = fmt.Sprintf("%v.BaseURL", cmdPrefix)
		}

		baseUrlFlagValue, err := cmd.Flags().GetString(baseUrlFlagName)
		if err != nil {
			return err, false
		}
		m.BaseURL = baseUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryEcrFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ecrFlagName := fmt.Sprintf("%v.Ecr", cmdPrefix)
	if cmd.Flags().Changed(ecrFlagName) {
		// info: complex object Ecr PortainereeEcrData is retrieved outside this Changed() block
	}
	ecrFlagValue := m.Ecr
	if swag.IsZero(ecrFlagValue) {
		ecrFlagValue = &models.PortainereeEcrData{}
	}

	err, ecrAdded := retrieveModelPortainereeEcrDataFlags(depth+1, ecrFlagValue, ecrFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ecrAdded
	if ecrAdded {
		m.Ecr = ecrFlagValue
	}

	return nil, retAdded
}

func retrievePortainereeRegistryGithubFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	githubFlagName := fmt.Sprintf("%v.Github", cmdPrefix)
	if cmd.Flags().Changed(githubFlagName) {
		// info: complex object Github PortainereeGithubRegistryData is retrieved outside this Changed() block
	}
	githubFlagValue := m.Github
	if swag.IsZero(githubFlagValue) {
		githubFlagValue = &models.PortainereeGithubRegistryData{}
	}

	err, githubAdded := retrieveModelPortainereeGithubRegistryDataFlags(depth+1, githubFlagValue, githubFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || githubAdded
	if githubAdded {
		m.Github = githubFlagValue
	}

	return nil, retAdded
}

func retrievePortainereeRegistryGitlabFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gitlabFlagName := fmt.Sprintf("%v.Gitlab", cmdPrefix)
	if cmd.Flags().Changed(gitlabFlagName) {
		// info: complex object Gitlab PortainereeGitlabRegistryData is retrieved outside this Changed() block
	}
	gitlabFlagValue := m.Gitlab
	if swag.IsZero(gitlabFlagValue) {
		gitlabFlagValue = &models.PortainereeGitlabRegistryData{}
	}

	err, gitlabAdded := retrieveModelPortainereeGitlabRegistryDataFlags(depth+1, gitlabFlagValue, gitlabFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gitlabAdded
	if gitlabAdded {
		m.Gitlab = gitlabFlagValue
	}

	return nil, retAdded
}

func retrievePortainereeRegistryIDFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.Id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "Id"
		} else {
			idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryManagementConfigurationFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	managementConfigurationFlagName := fmt.Sprintf("%v.ManagementConfiguration", cmdPrefix)
	if cmd.Flags().Changed(managementConfigurationFlagName) {
		// info: complex object ManagementConfiguration PortainereeRegistryManagementConfiguration is retrieved outside this Changed() block
	}
	managementConfigurationFlagValue := m.ManagementConfiguration
	if swag.IsZero(managementConfigurationFlagValue) {
		managementConfigurationFlagValue = &models.PortainereeRegistryManagementConfiguration{}
	}

	err, managementConfigurationAdded := retrieveModelPortainereeRegistryManagementConfigurationFlags(depth+1, managementConfigurationFlagValue, managementConfigurationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || managementConfigurationAdded
	if managementConfigurationAdded {
		m.ManagementConfiguration = managementConfigurationFlagValue
	}

	return nil, retAdded
}

func retrievePortainereeRegistryNameFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

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
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryPasswordFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.Password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "Password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.Password", cmdPrefix)
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

func retrievePortainereeRegistryQuayFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	quayFlagName := fmt.Sprintf("%v.Quay", cmdPrefix)
	if cmd.Flags().Changed(quayFlagName) {
		// info: complex object Quay PortainereeQuayRegistryData is retrieved outside this Changed() block
	}
	quayFlagValue := m.Quay
	if swag.IsZero(quayFlagValue) {
		quayFlagValue = &models.PortainereeQuayRegistryData{}
	}

	err, quayAdded := retrieveModelPortainereeQuayRegistryDataFlags(depth+1, quayFlagValue, quayFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quayAdded
	if quayAdded {
		m.Quay = quayFlagValue
	}

	return nil, retAdded
}

func retrievePortainereeRegistryRegistryAccessesFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	registryAccessesFlagName := fmt.Sprintf("%v.RegistryAccesses", cmdPrefix)
	if cmd.Flags().Changed(registryAccessesFlagName) {
		// warning: RegistryAccesses map type PortainereeRegistryAccesses is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeRegistryTeamAccessPoliciesFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	teamAccessPoliciesFlagName := fmt.Sprintf("%v.TeamAccessPolicies", cmdPrefix)
	if cmd.Flags().Changed(teamAccessPoliciesFlagName) {
		// warning: TeamAccessPolicies map type PortainereeTeamAccessPolicies is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeRegistryTypeFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.Type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "Type"
		} else {
			typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetInt64(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryURLFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urlFlagName := fmt.Sprintf("%v.URL", cmdPrefix)
	if cmd.Flags().Changed(urlFlagName) {

		var urlFlagName string
		if cmdPrefix == "" {
			urlFlagName = "URL"
		} else {
			urlFlagName = fmt.Sprintf("%v.URL", cmdPrefix)
		}

		urlFlagValue, err := cmd.Flags().GetString(urlFlagName)
		if err != nil {
			return err, false
		}
		m.URL = urlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeRegistryUserAccessPoliciesFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userAccessPoliciesFlagName := fmt.Sprintf("%v.UserAccessPolicies", cmdPrefix)
	if cmd.Flags().Changed(userAccessPoliciesFlagName) {
		// warning: UserAccessPolicies map type PortainereeUserAccessPolicies is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeRegistryUsernameFlags(depth int, m *models.PortainereeRegistry, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.Username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "Username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.Username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
