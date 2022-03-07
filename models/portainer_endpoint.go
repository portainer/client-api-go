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

// PortainerEndpoint portainer endpoint
//
// swagger:model portainer.Endpoint
type PortainerEndpoint struct {

	// authorized teams
	AuthorizedTeams []int64 `json:"AuthorizedTeams"`

	// Deprecated in DBVersion == 18
	AuthorizedUsers []int64 `json:"AuthorizedUsers"`

	// azure credentials
	AzureCredentials *PortainerAzureCredentials `json:"AzureCredentials,omitempty"`

	// The check in interval for edge agent (in seconds)
	// Example: 5
	EdgeCheckinInterval int64 `json:"EdgeCheckinInterval,omitempty"`

	// The identifier of the edge agent associated with this endpoint
	EdgeID string `json:"EdgeID,omitempty"`

	// The key which is used to map the agent to Portainer
	EdgeKey string `json:"EdgeKey,omitempty"`

	// extensions
	Extensions []*PortainerEndpointExtension `json:"Extensions"`

	// Endpoint group identifier
	// Example: 1
	GroupID int64 `json:"GroupId,omitempty"`

	// Endpoint Identifier
	// Example: 1
	ID int64 `json:"Id,omitempty"`

	// Associated Kubernetes data
	Kubernetes *PortainerKubernetesData `json:"Kubernetes,omitempty"`

	// Endpoint name
	// Example: my-endpoint
	Name string `json:"Name,omitempty"`

	// URL or IP address where exposed containers will be reachable
	// Example: docker.mydomain.tld:2375
	PublicURL string `json:"PublicURL,omitempty"`

	// List of snapshots
	Snapshots []*PortainerDockerSnapshot `json:"Snapshots"`

	// The status of the endpoint (1 - up, 2 - down)
	// Example: 1
	Status int64 `json:"Status,omitempty"`

	// Deprecated fields
	// Deprecated in DBVersion == 4
	TLS bool `json:"TLS,omitempty"`

	// TLS c a cert
	TLSCACert string `json:"TLSCACert,omitempty"`

	// TLS cert
	TLSCert string `json:"TLSCert,omitempty"`

	// TLS config
	TLSConfig *PortainerTLSConfiguration `json:"TLSConfig,omitempty"`

	// TLS key
	TLSKey string `json:"TLSKey,omitempty"`

	// List of tag identifiers to which this endpoint is associated
	TagIds []int64 `json:"TagIds"`

	// Deprecated in DBVersion == 22
	Tags []string `json:"Tags"`

	// List of team identifiers authorized to connect to this endpoint
	TeamAccessPolicies PortainerTeamAccessPolicies `json:"TeamAccessPolicies,omitempty"`

	// Endpoint environment type. 1 for a Docker environment, 2 for an agent on Docker environment or 3 for an Azure environment.
	// Example: 1
	Type int64 `json:"Type,omitempty"`

	// URL or IP address of the Docker host associated to this endpoint
	// Example: docker.mydomain.tld:2375
	URL string `json:"URL,omitempty"`

	// List of user identifiers authorized to connect to this endpoint
	UserAccessPolicies PortainerUserAccessPolicies `json:"UserAccessPolicies,omitempty"`
}

// Validate validates this portainer endpoint
func (m *PortainerEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerEndpoint) validateAzureCredentials(formats strfmt.Registry) error {
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

func (m *PortainerEndpoint) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Extensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerEndpoint) validateKubernetes(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerEndpoint) validateSnapshots(formats strfmt.Registry) error {
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

func (m *PortainerEndpoint) validateTLSConfig(formats strfmt.Registry) error {
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

func (m *PortainerEndpoint) validateTeamAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamAccessPolicies) { // not required
		return nil
	}

	if m.TeamAccessPolicies != nil {
		if err := m.TeamAccessPolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TeamAccessPolicies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TeamAccessPolicies")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerEndpoint) validateUserAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.UserAccessPolicies) { // not required
		return nil
	}

	if m.UserAccessPolicies != nil {
		if err := m.UserAccessPolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UserAccessPolicies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UserAccessPolicies")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portainer endpoint based on the context it is used
func (m *PortainerEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerEndpoint) contextValidateAzureCredentials(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainerEndpoint) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Extensions); i++ {

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Extensions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerEndpoint) contextValidateKubernetes(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubernetes != nil {
		if err := m.Kubernetes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Kubernetes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerEndpoint) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainerEndpoint) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortainerEndpoint) contextValidateTeamAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TeamAccessPolicies.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TeamAccessPolicies")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TeamAccessPolicies")
		}
		return err
	}

	return nil
}

func (m *PortainerEndpoint) contextValidateUserAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UserAccessPolicies.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("UserAccessPolicies")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("UserAccessPolicies")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerEndpoint) UnmarshalBinary(b []byte) error {
	var res PortainerEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
