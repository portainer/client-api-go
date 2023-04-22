// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortainereeEndpoint portaineree endpoint
//
// swagger:model portaineree.Endpoint
type PortainereeEndpoint struct {

	// The identifier of the AMT Device associated with this environment(endpoint)
	// Example: 4c4c4544-004b-3910-8037-b6c04f504633
	AMTDeviceGUID string `json:"AMTDeviceGUID,omitempty"`

	// authorized teams
	AuthorizedTeams []int64 `json:"AuthorizedTeams"`

	// Deprecated in DBVersion == 18
	AuthorizedUsers []int64 `json:"AuthorizedUsers"`

	// azure credentials
	AzureCredentials *PortainereeAzureCredentials `json:"AzureCredentials,omitempty"`

	// Automatic update change window restriction for stacks and apps
	ChangeWindow struct {
		PortainereeEndpointChangeWindow
	} `json:"ChangeWindow,omitempty"`

	// A Kubernetes as a service cloud provider. Only included if this
	// endpoint was created using KaaS provisioning.
	CloudProvider struct {
		PortainereeCloudProvider
	} `json:"CloudProvider,omitempty"`

	// Maximum version of docker-compose
	// Example: 3.8
	ComposeSyntaxMaxVersion string `json:"ComposeSyntaxMaxVersion,omitempty"`

	// Hide manual deployment forms for an environment
	DeploymentOptions struct {
		PortainereeDeploymentOptions
	} `json:"DeploymentOptions,omitempty"`

	// The check in interval for edge agent (in seconds)
	// Example: 5
	EdgeCheckinInterval int64 `json:"EdgeCheckinInterval,omitempty"`

	// The identifier of the edge agent associated with this environment(endpoint)
	EdgeID string `json:"EdgeID,omitempty"`

	// The key which is used to map the agent to Portainer
	EdgeKey string `json:"EdgeKey,omitempty"`

	// enable g p u management
	EnableGPUManagement bool `json:"EnableGPUManagement,omitempty"`

	// enable image notification
	EnableImageNotification bool `json:"EnableImageNotification,omitempty"`

	// gpus
	Gpus []*PortainereePair `json:"Gpus"`

	// Environment(Endpoint) group identifier
	// Example: 1
	GroupID int64 `json:"GroupId,omitempty"`

	// Heartbeat indicates the heartbeat status of an edge environment
	// Example: true
	Heartbeat bool `json:"Heartbeat,omitempty"`

	// Environment(Endpoint) Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// Associated Kubernetes data
	Kubernetes struct {
		PortainereeKubernetesData
	} `json:"Kubernetes,omitempty"`

	// Environment(Endpoint) name
	// Example: my-environment
	Name string `json:"Name,omitempty"`

	// Associated Nomad data
	Nomad struct {
		PortainereeNomadData
	} `json:"Nomad,omitempty"`

	// Whether we need to run any "post init migrations".
	PostInitMigrations struct {
		PortainereeEndpointPostInitMigrations
	} `json:"PostInitMigrations,omitempty"`

	// URL or IP address where exposed containers will be reachable
	// Example: docker.mydomain.tld:2375
	PublicURL string `json:"PublicURL,omitempty"`

	// List of snapshots
	Snapshots []*PortainerDockerSnapshot `json:"Snapshots"`

	// The status of the environment(endpoint) (1 - up, 2 - down, 3 -
	// provisioning, 4 - error)
	// Example: 1
	Status struct {
		PortainereeEndpointStatus
	} `json:"Status,omitempty"`

	// A message that describes the status. Should be included for Status 3
	// or 4.
	StatusMessage struct {
		PortainereeEndpointStatusMessage
	} `json:"StatusMessage,omitempty"`

	// Deprecated fields
	// Deprecated in DBVersion == 4
	TLS bool `json:"TLS,omitempty"`

	// TLS c a cert
	TLSCACert string `json:"TLSCACert,omitempty"`

	// TLS cert
	TLSCert string `json:"TLSCert,omitempty"`

	// TLS config
	TLSConfig *PortainereeTLSConfiguration `json:"TLSConfig,omitempty"`

	// TLS key
	TLSKey string `json:"TLSKey,omitempty"`

	// List of tag identifiers to which this environment(endpoint) is associated
	TagIds []int64 `json:"TagIds"`

	// Deprecated in DBVersion == 22
	Tags []string `json:"Tags"`

	// List of team identifiers authorized to connect to this environment(endpoint)
	TeamAccessPolicies struct {
		PortainereeTeamAccessPolicies
	} `json:"TeamAccessPolicies,omitempty"`

	// Environment(Endpoint) environment(endpoint) type. 1 for a Docker environment(endpoint), 2 for an agent on Docker environment(endpoint) or 3 for an Azure environment(endpoint).
	// Example: 1
	Type struct {
		PortainereeEndpointType
	} `json:"Type,omitempty"`

	// URL or IP address of the Docker host associated to this environment(endpoint)
	// Example: docker.mydomain.tld:2375
	URL string `json:"URL,omitempty"`

	// List of user identifiers authorized to connect to this environment(endpoint)
	UserAccessPolicies struct {
		PortainereeUserAccessPolicies
	} `json:"UserAccessPolicies,omitempty"`

	// agent
	Agent *PortainereeEndpointAgent `json:"agent,omitempty"`

	// edge
	Edge *PortainereeEnvironmentEdgeSettings `json:"edge,omitempty"`

	// IsEdgeDevice marks if the environment was created as an EdgeDevice
	IsEdgeDevice bool `json:"isEdgeDevice,omitempty"`

	// LastCheckInDate mark last check-in date on checkin
	LastCheckInDate int64 `json:"lastCheckInDate,omitempty"`

	// LocalTimeZone is the local time zone of the endpoint
	LocalTimeZone string `json:"localTimeZone,omitempty"`

	// QueryDate of each query with the endpoints list
	QueryDate int64 `json:"queryDate,omitempty"`

	// Environment(Endpoint) specific security settings
	SecuritySettings struct {
		PortainereeEndpointSecuritySettings
	} `json:"securitySettings,omitempty"`

	// Whether the device has been trusted or not by the user
	UserTrusted bool `json:"userTrusted,omitempty"`
}

// Validate validates this portaineree endpoint
func (m *PortainereeEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNomad(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostInitMigrations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecuritySettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeEndpoint) validateAzureCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureCredentials) { // not required
		return nil
	}

	if m.AzureCredentials != nil {
		if err := m.AzureCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AzureCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AzureCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) validateChangeWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangeWindow) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateCloudProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateDeploymentOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentOptions) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateGpus(formats strfmt.Registry) error {
	if swag.IsZero(m.Gpus) { // not required
		return nil
	}

	for i := 0; i < len(m.Gpus); i++ {
		if swag.IsZero(m.Gpus[i]) { // not required
			continue
		}

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeEndpoint) validateKubernetes(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateNomad(formats strfmt.Registry) error {
	if swag.IsZero(m.Nomad) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validatePostInitMigrations(formats strfmt.Registry) error {
	if swag.IsZero(m.PostInitMigrations) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeEndpoint) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateStatusMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusMessage) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) validateTeamAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamAccessPolicies) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateUserAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.UserAccessPolicies) { // not required
		return nil
	}

	return nil
}

func (m *PortainereeEndpoint) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) validateEdge(formats strfmt.Registry) error {
	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) validateSecuritySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SecuritySettings) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this portaineree endpoint based on the context it is used
func (m *PortainereeEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChangeWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNomad(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostInitMigrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecuritySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainereeEndpoint) contextValidateAzureCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureCredentials != nil {
		if err := m.AzureCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AzureCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AzureCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) contextValidateChangeWindow(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateCloudProvider(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateDeploymentOptions(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateGpus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Gpus); i++ {

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeEndpoint) contextValidateKubernetes(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateNomad(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidatePostInitMigrations(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Snapshots); i++ {

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainereeEndpoint) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateStatusMessage(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TLSConfig != nil {
		if err := m.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) contextValidateTeamAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateUserAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PortainereeEndpoint) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {
		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) contextValidateEdge(ctx context.Context, formats strfmt.Registry) error {

	if m.Edge != nil {
		if err := m.Edge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *PortainereeEndpoint) contextValidateSecuritySettings(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEndpoint) UnmarshalBinary(b []byte) error {
	var res PortainereeEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortainereeEndpointAgent portaineree endpoint agent
//
// swagger:model PortainereeEndpointAgent
type PortainereeEndpointAgent struct {

	// version
	// Example: 1.0.0
	Version string `json:"version,omitempty"`
}

// Validate validates this portaineree endpoint agent
func (m *PortainereeEndpointAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portaineree endpoint agent based on context it is used
func (m *PortainereeEndpointAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortainereeEndpointAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainereeEndpointAgent) UnmarshalBinary(b []byte) error {
	var res PortainereeEndpointAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
