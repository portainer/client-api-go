// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PortainereeCloudProvider

// register flags to command
func registerModelPortainereeCloudProviderFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPortainereeCloudProviderAmiType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderCPU(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderCredentialID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderHDD(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderInstanceType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderNetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderNodeCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderNodeVolumeSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderRAM(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderDnsprefix(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderKubernetesVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderPoolName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderResourceGroup(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPortainereeCloudProviderTier(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPortainereeCloudProviderAmiType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	amiTypeDescription := `Amazon specific fields`

	var amiTypeFlagName string
	if cmdPrefix == "" {
		amiTypeFlagName = "AmiType"
	} else {
		amiTypeFlagName = fmt.Sprintf("%v.AmiType", cmdPrefix)
	}

	var amiTypeFlagDefault string

	_ = cmd.PersistentFlags().String(amiTypeFlagName, amiTypeFlagDefault, amiTypeDescription)

	return nil
}

func registerPortainereeCloudProviderCPU(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpuDescription := ``

	var cpuFlagName string
	if cmdPrefix == "" {
		cpuFlagName = "CPU"
	} else {
		cpuFlagName = fmt.Sprintf("%v.CPU", cmdPrefix)
	}

	var cpuFlagDefault int64

	_ = cmd.PersistentFlags().Int64(cpuFlagName, cpuFlagDefault, cpuDescription)

	return nil
}

func registerPortainereeCloudProviderCredentialID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	credentialIdDescription := `CredentialID holds an ID of the credential used to create the cluster`

	var credentialIdFlagName string
	if cmdPrefix == "" {
		credentialIdFlagName = "CredentialID"
	} else {
		credentialIdFlagName = fmt.Sprintf("%v.CredentialID", cmdPrefix)
	}

	var credentialIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(credentialIdFlagName, credentialIdFlagDefault, credentialIdDescription)

	return nil
}

func registerPortainereeCloudProviderHDD(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hDDDescription := ``

	var hDDFlagName string
	if cmdPrefix == "" {
		hDDFlagName = "HDD"
	} else {
		hDDFlagName = fmt.Sprintf("%v.HDD", cmdPrefix)
	}

	var hDDFlagDefault int64

	_ = cmd.PersistentFlags().Int64(hDDFlagName, hDDFlagDefault, hDDDescription)

	return nil
}

func registerPortainereeCloudProviderInstanceType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	instanceTypeDescription := ``

	var instanceTypeFlagName string
	if cmdPrefix == "" {
		instanceTypeFlagName = "InstanceType"
	} else {
		instanceTypeFlagName = fmt.Sprintf("%v.InstanceType", cmdPrefix)
	}

	var instanceTypeFlagDefault string

	_ = cmd.PersistentFlags().String(instanceTypeFlagName, instanceTypeFlagDefault, instanceTypeDescription)

	return nil
}

func registerPortainereeCloudProviderName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerPortainereeCloudProviderNetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	networkIdDescription := `Pointer will hide this field for providers other than civo which do
not use this field.`

	var networkIdFlagName string
	if cmdPrefix == "" {
		networkIdFlagName = "NetworkID"
	} else {
		networkIdFlagName = fmt.Sprintf("%v.NetworkID", cmdPrefix)
	}

	var networkIdFlagDefault string

	_ = cmd.PersistentFlags().String(networkIdFlagName, networkIdFlagDefault, networkIdDescription)

	return nil
}

func registerPortainereeCloudProviderNodeCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeCountDescription := ``

	var nodeCountFlagName string
	if cmdPrefix == "" {
		nodeCountFlagName = "NodeCount"
	} else {
		nodeCountFlagName = fmt.Sprintf("%v.NodeCount", cmdPrefix)
	}

	var nodeCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(nodeCountFlagName, nodeCountFlagDefault, nodeCountDescription)

	return nil
}

func registerPortainereeCloudProviderNodeVolumeSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeVolumeSizeDescription := ``

	var nodeVolumeSizeFlagName string
	if cmdPrefix == "" {
		nodeVolumeSizeFlagName = "NodeVolumeSize"
	} else {
		nodeVolumeSizeFlagName = fmt.Sprintf("%v.NodeVolumeSize", cmdPrefix)
	}

	var nodeVolumeSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(nodeVolumeSizeFlagName, nodeVolumeSizeFlagDefault, nodeVolumeSizeDescription)

	return nil
}

func registerPortainereeCloudProviderRAM(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ramDescription := ``

	var ramFlagName string
	if cmdPrefix == "" {
		ramFlagName = "RAM"
	} else {
		ramFlagName = fmt.Sprintf("%v.RAM", cmdPrefix)
	}

	var ramFlagDefault float64

	_ = cmd.PersistentFlags().Float64(ramFlagName, ramFlagDefault, ramDescription)

	return nil
}

func registerPortainereeCloudProviderRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionDescription := ``

	var regionFlagName string
	if cmdPrefix == "" {
		regionFlagName = "Region"
	} else {
		regionFlagName = fmt.Sprintf("%v.Region", cmdPrefix)
	}

	var regionFlagDefault string

	_ = cmd.PersistentFlags().String(regionFlagName, regionFlagDefault, regionDescription)

	return nil
}

func registerPortainereeCloudProviderSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := ``

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "Size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
	}

	var sizeFlagDefault string

	_ = cmd.PersistentFlags().String(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerPortainereeCloudProviderURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := ``

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

func registerPortainereeCloudProviderDnsprefix(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dnsprefixDescription := ``

	var dnsprefixFlagName string
	if cmdPrefix == "" {
		dnsprefixFlagName = "dnsprefix"
	} else {
		dnsprefixFlagName = fmt.Sprintf("%v.dnsprefix", cmdPrefix)
	}

	var dnsprefixFlagDefault string

	_ = cmd.PersistentFlags().String(dnsprefixFlagName, dnsprefixFlagDefault, dnsprefixDescription)

	return nil
}

func registerPortainereeCloudProviderKubernetesVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	kubernetesVersionDescription := ``

	var kubernetesVersionFlagName string
	if cmdPrefix == "" {
		kubernetesVersionFlagName = "kubernetesVersion"
	} else {
		kubernetesVersionFlagName = fmt.Sprintf("%v.kubernetesVersion", cmdPrefix)
	}

	var kubernetesVersionFlagDefault string

	_ = cmd.PersistentFlags().String(kubernetesVersionFlagName, kubernetesVersionFlagDefault, kubernetesVersionDescription)

	return nil
}

func registerPortainereeCloudProviderPoolName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	poolNameDescription := ``

	var poolNameFlagName string
	if cmdPrefix == "" {
		poolNameFlagName = "poolName"
	} else {
		poolNameFlagName = fmt.Sprintf("%v.poolName", cmdPrefix)
	}

	var poolNameFlagDefault string

	_ = cmd.PersistentFlags().String(poolNameFlagName, poolNameFlagDefault, poolNameDescription)

	return nil
}

func registerPortainereeCloudProviderResourceGroup(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	resourceGroupDescription := `Azure specific fields`

	var resourceGroupFlagName string
	if cmdPrefix == "" {
		resourceGroupFlagName = "resourceGroup"
	} else {
		resourceGroupFlagName = fmt.Sprintf("%v.resourceGroup", cmdPrefix)
	}

	var resourceGroupFlagDefault string

	_ = cmd.PersistentFlags().String(resourceGroupFlagName, resourceGroupFlagDefault, resourceGroupDescription)

	return nil
}

func registerPortainereeCloudProviderTier(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tierDescription := ``

	var tierFlagName string
	if cmdPrefix == "" {
		tierFlagName = "tier"
	} else {
		tierFlagName = fmt.Sprintf("%v.tier", cmdPrefix)
	}

	var tierFlagDefault string

	_ = cmd.PersistentFlags().String(tierFlagName, tierFlagDefault, tierDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPortainereeCloudProviderFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, amiTypeAdded := retrievePortainereeCloudProviderAmiTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || amiTypeAdded

	err, cpuAdded := retrievePortainereeCloudProviderCPUFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuAdded

	err, credentialIdAdded := retrievePortainereeCloudProviderCredentialIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || credentialIdAdded

	err, hDDAdded := retrievePortainereeCloudProviderHDDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hDDAdded

	err, instanceTypeAdded := retrievePortainereeCloudProviderInstanceTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || instanceTypeAdded

	err, nameAdded := retrievePortainereeCloudProviderNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, networkIdAdded := retrievePortainereeCloudProviderNetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkIdAdded

	err, nodeCountAdded := retrievePortainereeCloudProviderNodeCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeCountAdded

	err, nodeVolumeSizeAdded := retrievePortainereeCloudProviderNodeVolumeSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeVolumeSizeAdded

	err, ramAdded := retrievePortainereeCloudProviderRAMFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ramAdded

	err, regionAdded := retrievePortainereeCloudProviderRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionAdded

	err, sizeAdded := retrievePortainereeCloudProviderSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, urlAdded := retrievePortainereeCloudProviderURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	err, dnsprefixAdded := retrievePortainereeCloudProviderDnsprefixFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dnsprefixAdded

	err, kubernetesVersionAdded := retrievePortainereeCloudProviderKubernetesVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || kubernetesVersionAdded

	err, poolNameAdded := retrievePortainereeCloudProviderPoolNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poolNameAdded

	err, resourceGroupAdded := retrievePortainereeCloudProviderResourceGroupFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || resourceGroupAdded

	err, tierAdded := retrievePortainereeCloudProviderTierFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tierAdded

	return nil, retAdded
}

func retrievePortainereeCloudProviderAmiTypeFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	amiTypeFlagName := fmt.Sprintf("%v.AmiType", cmdPrefix)
	if cmd.Flags().Changed(amiTypeFlagName) {

		var amiTypeFlagName string
		if cmdPrefix == "" {
			amiTypeFlagName = "AmiType"
		} else {
			amiTypeFlagName = fmt.Sprintf("%v.AmiType", cmdPrefix)
		}

		amiTypeFlagValue, err := cmd.Flags().GetString(amiTypeFlagName)
		if err != nil {
			return err, false
		}
		m.AmiType = amiTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderCPUFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuFlagName := fmt.Sprintf("%v.CPU", cmdPrefix)
	if cmd.Flags().Changed(cpuFlagName) {

		var cpuFlagName string
		if cmdPrefix == "" {
			cpuFlagName = "CPU"
		} else {
			cpuFlagName = fmt.Sprintf("%v.CPU", cmdPrefix)
		}

		cpuFlagValue, err := cmd.Flags().GetInt64(cpuFlagName)
		if err != nil {
			return err, false
		}
		m.CPU = cpuFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderCredentialIDFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	credentialIdFlagName := fmt.Sprintf("%v.CredentialID", cmdPrefix)
	if cmd.Flags().Changed(credentialIdFlagName) {

		var credentialIdFlagName string
		if cmdPrefix == "" {
			credentialIdFlagName = "CredentialID"
		} else {
			credentialIdFlagName = fmt.Sprintf("%v.CredentialID", cmdPrefix)
		}

		credentialIdFlagValue, err := cmd.Flags().GetInt64(credentialIdFlagName)
		if err != nil {
			return err, false
		}
		m.CredentialID = credentialIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderHDDFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hDDFlagName := fmt.Sprintf("%v.HDD", cmdPrefix)
	if cmd.Flags().Changed(hDDFlagName) {

		var hDDFlagName string
		if cmdPrefix == "" {
			hDDFlagName = "HDD"
		} else {
			hDDFlagName = fmt.Sprintf("%v.HDD", cmdPrefix)
		}

		hDDFlagValue, err := cmd.Flags().GetInt64(hDDFlagName)
		if err != nil {
			return err, false
		}
		m.HDD = hDDFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderInstanceTypeFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	instanceTypeFlagName := fmt.Sprintf("%v.InstanceType", cmdPrefix)
	if cmd.Flags().Changed(instanceTypeFlagName) {

		var instanceTypeFlagName string
		if cmdPrefix == "" {
			instanceTypeFlagName = "InstanceType"
		} else {
			instanceTypeFlagName = fmt.Sprintf("%v.InstanceType", cmdPrefix)
		}

		instanceTypeFlagValue, err := cmd.Flags().GetString(instanceTypeFlagName)
		if err != nil {
			return err, false
		}
		m.InstanceType = instanceTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderNameFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereeCloudProviderNetworkIDFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkIdFlagName := fmt.Sprintf("%v.NetworkID", cmdPrefix)
	if cmd.Flags().Changed(networkIdFlagName) {

		var networkIdFlagName string
		if cmdPrefix == "" {
			networkIdFlagName = "NetworkID"
		} else {
			networkIdFlagName = fmt.Sprintf("%v.NetworkID", cmdPrefix)
		}

		networkIdFlagValue, err := cmd.Flags().GetString(networkIdFlagName)
		if err != nil {
			return err, false
		}
		m.NetworkID = networkIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderNodeCountFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nodeCountFlagName := fmt.Sprintf("%v.NodeCount", cmdPrefix)
	if cmd.Flags().Changed(nodeCountFlagName) {

		var nodeCountFlagName string
		if cmdPrefix == "" {
			nodeCountFlagName = "NodeCount"
		} else {
			nodeCountFlagName = fmt.Sprintf("%v.NodeCount", cmdPrefix)
		}

		nodeCountFlagValue, err := cmd.Flags().GetInt64(nodeCountFlagName)
		if err != nil {
			return err, false
		}
		m.NodeCount = nodeCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderNodeVolumeSizeFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nodeVolumeSizeFlagName := fmt.Sprintf("%v.NodeVolumeSize", cmdPrefix)
	if cmd.Flags().Changed(nodeVolumeSizeFlagName) {

		var nodeVolumeSizeFlagName string
		if cmdPrefix == "" {
			nodeVolumeSizeFlagName = "NodeVolumeSize"
		} else {
			nodeVolumeSizeFlagName = fmt.Sprintf("%v.NodeVolumeSize", cmdPrefix)
		}

		nodeVolumeSizeFlagValue, err := cmd.Flags().GetInt64(nodeVolumeSizeFlagName)
		if err != nil {
			return err, false
		}
		m.NodeVolumeSize = nodeVolumeSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderRAMFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ramFlagName := fmt.Sprintf("%v.RAM", cmdPrefix)
	if cmd.Flags().Changed(ramFlagName) {

		var ramFlagName string
		if cmdPrefix == "" {
			ramFlagName = "RAM"
		} else {
			ramFlagName = fmt.Sprintf("%v.RAM", cmdPrefix)
		}

		ramFlagValue, err := cmd.Flags().GetFloat64(ramFlagName)
		if err != nil {
			return err, false
		}
		m.RAM = ramFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderRegionFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionFlagName := fmt.Sprintf("%v.Region", cmdPrefix)
	if cmd.Flags().Changed(regionFlagName) {

		var regionFlagName string
		if cmdPrefix == "" {
			regionFlagName = "Region"
		} else {
			regionFlagName = fmt.Sprintf("%v.Region", cmdPrefix)
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

func retrievePortainereeCloudProviderSizeFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeFlagName := fmt.Sprintf("%v.Size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "Size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetString(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderURLFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePortainereeCloudProviderDnsprefixFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dnsprefixFlagName := fmt.Sprintf("%v.dnsprefix", cmdPrefix)
	if cmd.Flags().Changed(dnsprefixFlagName) {

		var dnsprefixFlagName string
		if cmdPrefix == "" {
			dnsprefixFlagName = "dnsprefix"
		} else {
			dnsprefixFlagName = fmt.Sprintf("%v.dnsprefix", cmdPrefix)
		}

		dnsprefixFlagValue, err := cmd.Flags().GetString(dnsprefixFlagName)
		if err != nil {
			return err, false
		}
		m.Dnsprefix = dnsprefixFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderKubernetesVersionFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	kubernetesVersionFlagName := fmt.Sprintf("%v.kubernetesVersion", cmdPrefix)
	if cmd.Flags().Changed(kubernetesVersionFlagName) {

		var kubernetesVersionFlagName string
		if cmdPrefix == "" {
			kubernetesVersionFlagName = "kubernetesVersion"
		} else {
			kubernetesVersionFlagName = fmt.Sprintf("%v.kubernetesVersion", cmdPrefix)
		}

		kubernetesVersionFlagValue, err := cmd.Flags().GetString(kubernetesVersionFlagName)
		if err != nil {
			return err, false
		}
		m.KubernetesVersion = kubernetesVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderPoolNameFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	poolNameFlagName := fmt.Sprintf("%v.poolName", cmdPrefix)
	if cmd.Flags().Changed(poolNameFlagName) {

		var poolNameFlagName string
		if cmdPrefix == "" {
			poolNameFlagName = "poolName"
		} else {
			poolNameFlagName = fmt.Sprintf("%v.poolName", cmdPrefix)
		}

		poolNameFlagValue, err := cmd.Flags().GetString(poolNameFlagName)
		if err != nil {
			return err, false
		}
		m.PoolName = poolNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderResourceGroupFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	resourceGroupFlagName := fmt.Sprintf("%v.resourceGroup", cmdPrefix)
	if cmd.Flags().Changed(resourceGroupFlagName) {

		var resourceGroupFlagName string
		if cmdPrefix == "" {
			resourceGroupFlagName = "resourceGroup"
		} else {
			resourceGroupFlagName = fmt.Sprintf("%v.resourceGroup", cmdPrefix)
		}

		resourceGroupFlagValue, err := cmd.Flags().GetString(resourceGroupFlagName)
		if err != nil {
			return err, false
		}
		m.ResourceGroup = resourceGroupFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePortainereeCloudProviderTierFlags(depth int, m *models.PortainereeCloudProvider, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tierFlagName := fmt.Sprintf("%v.tier", cmdPrefix)
	if cmd.Flags().Changed(tierFlagName) {

		var tierFlagName string
		if cmdPrefix == "" {
			tierFlagName = "tier"
		} else {
			tierFlagName = fmt.Sprintf("%v.tier", cmdPrefix)
		}

		tierFlagValue, err := cmd.Flags().GetString(tierFlagName)
		if err != nil {
			return err, false
		}
		m.Tier = tierFlagValue

		retAdded = true
	}

	return nil, retAdded
}
