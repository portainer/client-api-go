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

// Schema cli for PortainereeLDAPSettings

// register flags to command
func registerModelPortainereeLDAPSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeLDAPSettingsAdminAutoPopulate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsAdminGroupSearchSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsAdminGroups(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsAnonymousMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsAutoCreateUsers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsGroupSearchSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsReaderDN(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsSearchSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsServerType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsStartTLS(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsTLSConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeLDAPSettingsURLs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeLDAPSettingsAdminAutoPopulate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	adminAutoPopulateDescription := `Whether auto admin population is switched on or not`

	var adminAutoPopulateFlagName string
	if cmdPrefix == "" {
		adminAutoPopulateFlagName = "AdminAutoPopulate"
	} else {
		adminAutoPopulateFlagName = fmt.Sprintf("%v.AdminAutoPopulate", cmdPrefix)
	}

	var adminAutoPopulateFlagDefault bool

	_ = cmd.PersistentFlags().Bool(adminAutoPopulateFlagName, adminAutoPopulateFlagDefault, adminAutoPopulateDescription)

	return nil
}

func registerPortainereeLDAPSettingsAdminGroupSearchSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: AdminGroupSearchSettings []*PortainereeLDAPGroupSearchSettings array type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeLDAPSettingsAdminGroups(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: AdminGroups []string array type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeLDAPSettingsAnonymousMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	anonymousModeDescription := `Enable this option if the server is configured for Anonymous access. When enabled, ReaderDN and Password will not be used`

	var anonymousModeFlagName string
	if cmdPrefix == "" {
		anonymousModeFlagName = "AnonymousMode"
	} else {
		anonymousModeFlagName = fmt.Sprintf("%v.AnonymousMode", cmdPrefix)
	}

	var anonymousModeFlagDefault bool

	_ = cmd.PersistentFlags().Bool(anonymousModeFlagName, anonymousModeFlagDefault, anonymousModeDescription)

	return nil
}

func registerPortainereeLDAPSettingsAutoCreateUsers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	autoCreateUsersDescription := `Automatically provision users and assign them to matching LDAP group names`

	var autoCreateUsersFlagName string
	if cmdPrefix == "" {
		autoCreateUsersFlagName = "AutoCreateUsers"
	} else {
		autoCreateUsersFlagName = fmt.Sprintf("%v.AutoCreateUsers", cmdPrefix)
	}

	var autoCreateUsersFlagDefault bool

	_ = cmd.PersistentFlags().Bool(autoCreateUsersFlagName, autoCreateUsersFlagDefault, autoCreateUsersDescription)

	return nil
}

func registerPortainereeLDAPSettingsGroupSearchSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: GroupSearchSettings []*PortainereeLDAPGroupSearchSettings array type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeLDAPSettingsPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Password of the account that will be used to search users`

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

func registerPortainereeLDAPSettingsReaderDN(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	readerDNDescription := `Account that will be used to search for users`

	var readerDNFlagName string
	if cmdPrefix == "" {
		readerDNFlagName = "ReaderDN"
	} else {
		readerDNFlagName = fmt.Sprintf("%v.ReaderDN", cmdPrefix)
	}

	var readerDNFlagDefault string

	_ = cmd.PersistentFlags().String(readerDNFlagName, readerDNFlagDefault, readerDNDescription)

	return nil
}

func registerPortainereeLDAPSettingsSearchSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: SearchSettings []*PortainereeLDAPSearchSettings array type is not supported by go-swagger cli yet

	return nil
}

func registerPortainereeLDAPSettingsServerType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serverTypeDescription := ``

	var serverTypeFlagName string
	if cmdPrefix == "" {
		serverTypeFlagName = "ServerType"
	} else {
		serverTypeFlagName = fmt.Sprintf("%v.ServerType", cmdPrefix)
	}

	var serverTypeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(serverTypeFlagName, serverTypeFlagDefault, serverTypeDescription)

	return nil
}

func registerPortainereeLDAPSettingsStartTLS(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startTlsDescription := `Whether LDAP connection should use StartTLS`

	var startTlsFlagName string
	if cmdPrefix == "" {
		startTlsFlagName = "StartTLS"
	} else {
		startTlsFlagName = fmt.Sprintf("%v.StartTLS", cmdPrefix)
	}

	var startTlsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(startTlsFlagName, startTlsFlagDefault, startTlsDescription)

	return nil
}

func registerPortainereeLDAPSettingsTLSConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var tlsConfigFlagName string
	if cmdPrefix == "" {
		tlsConfigFlagName = "TLSConfig"
	} else {
		tlsConfigFlagName = fmt.Sprintf("%v.TLSConfig", cmdPrefix)
	}

	if err := registerModelPortainereeTLSConfigurationFlags(depth+1, tlsConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeLDAPSettingsURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := `Deprecated`

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

func registerPortainereeLDAPSettingsURLs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: URLs []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereeLDAPSettingsFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, adminAutoPopulateAdded := retrievePortainereeLDAPSettingsAdminAutoPopulateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminAutoPopulateAdded

	err, adminGroupSearchSettingsAdded := retrievePortainereeLDAPSettingsAdminGroupSearchSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminGroupSearchSettingsAdded

	err, adminGroupsAdded := retrievePortainereeLDAPSettingsAdminGroupsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminGroupsAdded

	err, anonymousModeAdded := retrievePortainereeLDAPSettingsAnonymousModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || anonymousModeAdded

	err, autoCreateUsersAdded := retrievePortainereeLDAPSettingsAutoCreateUsersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || autoCreateUsersAdded

	err, groupSearchSettingsAdded := retrievePortainereeLDAPSettingsGroupSearchSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || groupSearchSettingsAdded

	err, passwordAdded := retrievePortainereeLDAPSettingsPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, readerDNAdded := retrievePortainereeLDAPSettingsReaderDNFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || readerDNAdded

	err, searchSettingsAdded := retrievePortainereeLDAPSettingsSearchSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || searchSettingsAdded

	err, serverTypeAdded := retrievePortainereeLDAPSettingsServerTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverTypeAdded

	err, startTlsAdded := retrievePortainereeLDAPSettingsStartTLSFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startTlsAdded

	err, tlsConfigAdded := retrievePortainereeLDAPSettingsTLSConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tlsConfigAdded

	err, urlAdded := retrievePortainereeLDAPSettingsURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	err, uRLsAdded := retrievePortainereeLDAPSettingsURLsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uRLsAdded

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsAdminAutoPopulateFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adminAutoPopulateFlagName := fmt.Sprintf("%v.AdminAutoPopulate", cmdPrefix)
	if cmd.Flags().Changed(adminAutoPopulateFlagName) {

		var adminAutoPopulateFlagName string
		if cmdPrefix == "" {
			adminAutoPopulateFlagName = "AdminAutoPopulate"
		} else {
			adminAutoPopulateFlagName = fmt.Sprintf("%v.AdminAutoPopulate", cmdPrefix)
		}

		adminAutoPopulateFlagValue, err := cmd.Flags().GetBool(adminAutoPopulateFlagName)
		if err != nil {
			return err, false
		}
		m.AdminAutoPopulate = &adminAutoPopulateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsAdminGroupSearchSettingsFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adminGroupSearchSettingsFlagName := fmt.Sprintf("%v.AdminGroupSearchSettings", cmdPrefix)
	if cmd.Flags().Changed(adminGroupSearchSettingsFlagName) {
		// warning: AdminGroupSearchSettings array type []*PortainereeLDAPGroupSearchSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsAdminGroupsFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adminGroupsFlagName := fmt.Sprintf("%v.AdminGroups", cmdPrefix)
	if cmd.Flags().Changed(adminGroupsFlagName) {
		// warning: AdminGroups array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsAnonymousModeFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	anonymousModeFlagName := fmt.Sprintf("%v.AnonymousMode", cmdPrefix)
	if cmd.Flags().Changed(anonymousModeFlagName) {

		var anonymousModeFlagName string
		if cmdPrefix == "" {
			anonymousModeFlagName = "AnonymousMode"
		} else {
			anonymousModeFlagName = fmt.Sprintf("%v.AnonymousMode", cmdPrefix)
		}

		anonymousModeFlagValue, err := cmd.Flags().GetBool(anonymousModeFlagName)
		if err != nil {
			return err, false
		}
		m.AnonymousMode = &anonymousModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsAutoCreateUsersFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	autoCreateUsersFlagName := fmt.Sprintf("%v.AutoCreateUsers", cmdPrefix)
	if cmd.Flags().Changed(autoCreateUsersFlagName) {

		var autoCreateUsersFlagName string
		if cmdPrefix == "" {
			autoCreateUsersFlagName = "AutoCreateUsers"
		} else {
			autoCreateUsersFlagName = fmt.Sprintf("%v.AutoCreateUsers", cmdPrefix)
		}

		autoCreateUsersFlagValue, err := cmd.Flags().GetBool(autoCreateUsersFlagName)
		if err != nil {
			return err, false
		}
		m.AutoCreateUsers = &autoCreateUsersFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsGroupSearchSettingsFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	groupSearchSettingsFlagName := fmt.Sprintf("%v.GroupSearchSettings", cmdPrefix)
	if cmd.Flags().Changed(groupSearchSettingsFlagName) {
		// warning: GroupSearchSettings array type []*PortainereeLDAPGroupSearchSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsPasswordFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereeLDAPSettingsReaderDNFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	readerDNFlagName := fmt.Sprintf("%v.ReaderDN", cmdPrefix)
	if cmd.Flags().Changed(readerDNFlagName) {

		var readerDNFlagName string
		if cmdPrefix == "" {
			readerDNFlagName = "ReaderDN"
		} else {
			readerDNFlagName = fmt.Sprintf("%v.ReaderDN", cmdPrefix)
		}

		readerDNFlagValue, err := cmd.Flags().GetString(readerDNFlagName)
		if err != nil {
			return err, false
		}
		m.ReaderDN = readerDNFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsSearchSettingsFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	searchSettingsFlagName := fmt.Sprintf("%v.SearchSettings", cmdPrefix)
	if cmd.Flags().Changed(searchSettingsFlagName) {
		// warning: SearchSettings array type []*PortainereeLDAPSearchSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsServerTypeFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverTypeFlagName := fmt.Sprintf("%v.ServerType", cmdPrefix)
	if cmd.Flags().Changed(serverTypeFlagName) {

		var serverTypeFlagName string
		if cmdPrefix == "" {
			serverTypeFlagName = "ServerType"
		} else {
			serverTypeFlagName = fmt.Sprintf("%v.ServerType", cmdPrefix)
		}

		serverTypeFlagValue, err := cmd.Flags().GetInt64(serverTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ServerType = serverTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsStartTLSFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startTlsFlagName := fmt.Sprintf("%v.StartTLS", cmdPrefix)
	if cmd.Flags().Changed(startTlsFlagName) {

		var startTlsFlagName string
		if cmdPrefix == "" {
			startTlsFlagName = "StartTLS"
		} else {
			startTlsFlagName = fmt.Sprintf("%v.StartTLS", cmdPrefix)
		}

		startTlsFlagValue, err := cmd.Flags().GetBool(startTlsFlagName)
		if err != nil {
			return err, false
		}
		m.StartTLS = &startTlsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsTLSConfigFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tlsConfigFlagName := fmt.Sprintf("%v.TLSConfig", cmdPrefix)
	if cmd.Flags().Changed(tlsConfigFlagName) {
		// info: complex object TLSConfig PortainereeTLSConfiguration is retrieved outside this Changed() block
	}
	tlsConfigFlagValue := m.TLSConfig
	if swag.IsZero(tlsConfigFlagValue) {
		tlsConfigFlagValue = &models.PortainereeTLSConfiguration{}
	}

	err, tlsConfigAdded := retrieveModelPortainereeTLSConfigurationFlags(depth+1, tlsConfigFlagValue, tlsConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tlsConfigAdded
	if tlsConfigAdded {
		m.TLSConfig = tlsConfigFlagValue
	}

	return nil, retAdded
}

func retrievePortainereeLDAPSettingsURLFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereeLDAPSettingsURLsFlags(depth int, m *models.PortainereeLDAPSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uRLsFlagName := fmt.Sprintf("%v.URLs", cmdPrefix)
	if cmd.Flags().Changed(uRLsFlagName) {
		// warning: URLs array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
