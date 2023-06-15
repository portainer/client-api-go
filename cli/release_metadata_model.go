// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for ReleaseMetadata

// register flags to command
func registerModelReleaseMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerReleaseMetadataAnnotations(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataAPIVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataAppVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataCondition(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataDependencies(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataDeprecated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataHome(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataIcon(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataKeywords(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataKubeVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataMaintainers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataSources(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReleaseMetadataVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerReleaseMetadataAnnotations(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: annotations map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerReleaseMetadataAPIVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apiVersionDescription := `The API Version of this chart. Required.`

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "apiVersion"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.apiVersion", cmdPrefix)
	}

	var apiVersionFlagDefault string

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}

func registerReleaseMetadataAppVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	appVersionDescription := `The version of the application enclosed inside of this chart.`

	var appVersionFlagName string
	if cmdPrefix == "" {
		appVersionFlagName = "appVersion"
	} else {
		appVersionFlagName = fmt.Sprintf("%v.appVersion", cmdPrefix)
	}

	var appVersionFlagDefault string

	_ = cmd.PersistentFlags().String(appVersionFlagName, appVersionFlagDefault, appVersionDescription)

	return nil
}

func registerReleaseMetadataCondition(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	conditionDescription := `The condition to check to enable chart`

	var conditionFlagName string
	if cmdPrefix == "" {
		conditionFlagName = "condition"
	} else {
		conditionFlagName = fmt.Sprintf("%v.condition", cmdPrefix)
	}

	var conditionFlagDefault string

	_ = cmd.PersistentFlags().String(conditionFlagName, conditionFlagDefault, conditionDescription)

	return nil
}

func registerReleaseMetadataDependencies(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: dependencies []*ReleaseDependency array type is not supported by go-swagger cli yet

	return nil
}

func registerReleaseMetadataDeprecated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deprecatedDescription := `Whether or not this chart is deprecated`

	var deprecatedFlagName string
	if cmdPrefix == "" {
		deprecatedFlagName = "deprecated"
	} else {
		deprecatedFlagName = fmt.Sprintf("%v.deprecated", cmdPrefix)
	}

	var deprecatedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(deprecatedFlagName, deprecatedFlagDefault, deprecatedDescription)

	return nil
}

func registerReleaseMetadataDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `A one-sentence description of the chart`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerReleaseMetadataHome(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	homeDescription := `The URL to a relevant project page, git repo, or contact person`

	var homeFlagName string
	if cmdPrefix == "" {
		homeFlagName = "home"
	} else {
		homeFlagName = fmt.Sprintf("%v.home", cmdPrefix)
	}

	var homeFlagDefault string

	_ = cmd.PersistentFlags().String(homeFlagName, homeFlagDefault, homeDescription)

	return nil
}

func registerReleaseMetadataIcon(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	iconDescription := `The URL to an icon file.`

	var iconFlagName string
	if cmdPrefix == "" {
		iconFlagName = "icon"
	} else {
		iconFlagName = fmt.Sprintf("%v.icon", cmdPrefix)
	}

	var iconFlagDefault string

	_ = cmd.PersistentFlags().String(iconFlagName, iconFlagDefault, iconDescription)

	return nil
}

func registerReleaseMetadataKeywords(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: keywords []string array type is not supported by go-swagger cli yet

	return nil
}

func registerReleaseMetadataKubeVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	kubeVersionDescription := `KubeVersion is a SemVer constraint specifying the version of Kubernetes required.`

	var kubeVersionFlagName string
	if cmdPrefix == "" {
		kubeVersionFlagName = "kubeVersion"
	} else {
		kubeVersionFlagName = fmt.Sprintf("%v.kubeVersion", cmdPrefix)
	}

	var kubeVersionFlagDefault string

	_ = cmd.PersistentFlags().String(kubeVersionFlagName, kubeVersionFlagDefault, kubeVersionDescription)

	return nil
}

func registerReleaseMetadataMaintainers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: maintainers []*ReleaseMaintainer array type is not supported by go-swagger cli yet

	return nil
}

func registerReleaseMetadataName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `The name of the chart. Required.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerReleaseMetadataSources(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: sources []string array type is not supported by go-swagger cli yet

	return nil
}

func registerReleaseMetadataTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tagsDescription := `The tags to check to enable chart`

	var tagsFlagName string
	if cmdPrefix == "" {
		tagsFlagName = "tags"
	} else {
		tagsFlagName = fmt.Sprintf("%v.tags", cmdPrefix)
	}

	var tagsFlagDefault string

	_ = cmd.PersistentFlags().String(tagsFlagName, tagsFlagDefault, tagsDescription)

	return nil
}

func registerReleaseMetadataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Specifies the chart type: application or library`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func registerReleaseMetadataVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `A SemVer 2 conformant version string of the chart. Required.`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelReleaseMetadataFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, annotationsAdded := retrieveReleaseMetadataAnnotationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || annotationsAdded

	err, apiVersionAdded := retrieveReleaseMetadataAPIVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apiVersionAdded

	err, appVersionAdded := retrieveReleaseMetadataAppVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || appVersionAdded

	err, conditionAdded := retrieveReleaseMetadataConditionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || conditionAdded

	err, dependenciesAdded := retrieveReleaseMetadataDependenciesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dependenciesAdded

	err, deprecatedAdded := retrieveReleaseMetadataDeprecatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deprecatedAdded

	err, descriptionAdded := retrieveReleaseMetadataDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, homeAdded := retrieveReleaseMetadataHomeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || homeAdded

	err, iconAdded := retrieveReleaseMetadataIconFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || iconAdded

	err, keywordsAdded := retrieveReleaseMetadataKeywordsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keywordsAdded

	err, kubeVersionAdded := retrieveReleaseMetadataKubeVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || kubeVersionAdded

	err, maintainersAdded := retrieveReleaseMetadataMaintainersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maintainersAdded

	err, nameAdded := retrieveReleaseMetadataNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, sourcesAdded := retrieveReleaseMetadataSourcesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourcesAdded

	err, tagsAdded := retrieveReleaseMetadataTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tagsAdded

	err, typeAdded := retrieveReleaseMetadataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, versionAdded := retrieveReleaseMetadataVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveReleaseMetadataAnnotationsFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	annotationsFlagName := fmt.Sprintf("%v.annotations", cmdPrefix)
	if cmd.Flags().Changed(annotationsFlagName) {
		// warning: annotations map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveReleaseMetadataAPIVersionFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apiVersionFlagName := fmt.Sprintf("%v.apiVersion", cmdPrefix)
	if cmd.Flags().Changed(apiVersionFlagName) {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "apiVersion"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.apiVersion", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = apiVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataAppVersionFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	appVersionFlagName := fmt.Sprintf("%v.appVersion", cmdPrefix)
	if cmd.Flags().Changed(appVersionFlagName) {

		var appVersionFlagName string
		if cmdPrefix == "" {
			appVersionFlagName = "appVersion"
		} else {
			appVersionFlagName = fmt.Sprintf("%v.appVersion", cmdPrefix)
		}

		appVersionFlagValue, err := cmd.Flags().GetString(appVersionFlagName)
		if err != nil {
			return err, false
		}
		m.AppVersion = appVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataConditionFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	conditionFlagName := fmt.Sprintf("%v.condition", cmdPrefix)
	if cmd.Flags().Changed(conditionFlagName) {

		var conditionFlagName string
		if cmdPrefix == "" {
			conditionFlagName = "condition"
		} else {
			conditionFlagName = fmt.Sprintf("%v.condition", cmdPrefix)
		}

		conditionFlagValue, err := cmd.Flags().GetString(conditionFlagName)
		if err != nil {
			return err, false
		}
		m.Condition = conditionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataDependenciesFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dependenciesFlagName := fmt.Sprintf("%v.dependencies", cmdPrefix)
	if cmd.Flags().Changed(dependenciesFlagName) {
		// warning: dependencies array type []*ReleaseDependency is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveReleaseMetadataDeprecatedFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deprecatedFlagName := fmt.Sprintf("%v.deprecated", cmdPrefix)
	if cmd.Flags().Changed(deprecatedFlagName) {

		var deprecatedFlagName string
		if cmdPrefix == "" {
			deprecatedFlagName = "deprecated"
		} else {
			deprecatedFlagName = fmt.Sprintf("%v.deprecated", cmdPrefix)
		}

		deprecatedFlagValue, err := cmd.Flags().GetBool(deprecatedFlagName)
		if err != nil {
			return err, false
		}
		m.Deprecated = &deprecatedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataDescriptionFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataHomeFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	homeFlagName := fmt.Sprintf("%v.home", cmdPrefix)
	if cmd.Flags().Changed(homeFlagName) {

		var homeFlagName string
		if cmdPrefix == "" {
			homeFlagName = "home"
		} else {
			homeFlagName = fmt.Sprintf("%v.home", cmdPrefix)
		}

		homeFlagValue, err := cmd.Flags().GetString(homeFlagName)
		if err != nil {
			return err, false
		}
		m.Home = homeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataIconFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	iconFlagName := fmt.Sprintf("%v.icon", cmdPrefix)
	if cmd.Flags().Changed(iconFlagName) {

		var iconFlagName string
		if cmdPrefix == "" {
			iconFlagName = "icon"
		} else {
			iconFlagName = fmt.Sprintf("%v.icon", cmdPrefix)
		}

		iconFlagValue, err := cmd.Flags().GetString(iconFlagName)
		if err != nil {
			return err, false
		}
		m.Icon = iconFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataKeywordsFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	keywordsFlagName := fmt.Sprintf("%v.keywords", cmdPrefix)
	if cmd.Flags().Changed(keywordsFlagName) {
		// warning: keywords array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveReleaseMetadataKubeVersionFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	kubeVersionFlagName := fmt.Sprintf("%v.kubeVersion", cmdPrefix)
	if cmd.Flags().Changed(kubeVersionFlagName) {

		var kubeVersionFlagName string
		if cmdPrefix == "" {
			kubeVersionFlagName = "kubeVersion"
		} else {
			kubeVersionFlagName = fmt.Sprintf("%v.kubeVersion", cmdPrefix)
		}

		kubeVersionFlagValue, err := cmd.Flags().GetString(kubeVersionFlagName)
		if err != nil {
			return err, false
		}
		m.KubeVersion = kubeVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataMaintainersFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maintainersFlagName := fmt.Sprintf("%v.maintainers", cmdPrefix)
	if cmd.Flags().Changed(maintainersFlagName) {
		// warning: maintainers array type []*ReleaseMaintainer is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveReleaseMetadataNameFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
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

func retrieveReleaseMetadataSourcesFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourcesFlagName := fmt.Sprintf("%v.sources", cmdPrefix)
	if cmd.Flags().Changed(sourcesFlagName) {
		// warning: sources array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveReleaseMetadataTagsFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tagsFlagName := fmt.Sprintf("%v.tags", cmdPrefix)
	if cmd.Flags().Changed(tagsFlagName) {

		var tagsFlagName string
		if cmdPrefix == "" {
			tagsFlagName = "tags"
		} else {
			tagsFlagName = fmt.Sprintf("%v.tags", cmdPrefix)
		}

		tagsFlagValue, err := cmd.Flags().GetString(tagsFlagName)
		if err != nil {
			return err, false
		}
		m.Tags = tagsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataTypeFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReleaseMetadataVersionFlags(depth int, m *models.ReleaseMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
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
