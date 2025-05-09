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

// PortainerTemplate portainer template
//
// swagger:model portainer.Template
type PortainerTemplate struct {

	// Whether the template should be available to administrators only
	// Example: true
	AdministratorOnly bool `json:"administrator_only,omitempty"`

	// A list of categories associated to the template
	// Example: ["database"]
	Categories []string `json:"categories"`

	// The command that will be executed in a container template
	// Example: ls -lah
	Command string `json:"command,omitempty"`

	// Description of the template
	// Example: High performance web server
	Description string `json:"description,omitempty"`

	// A list of environment(endpoint) variables used during the template deployment
	Env []*PortainerTemplateEnv `json:"env"`

	// Container hostname
	// Example: mycontainer
	Hostname string `json:"hostname,omitempty"`

	// Mandatory container/stack fields
	// Template Identifier
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Mandatory container fields
	// Image associated to a container template. Mandatory for a container template
	// Example: nginx:latest
	Image string `json:"image,omitempty"`

	// Whether the container should be started in
	// interactive mode (-i -t equivalent on the CLI)
	// Example: true
	Interactive bool `json:"interactive,omitempty"`

	// Container labels
	Labels []*PortainerPair `json:"labels"`

	// URL of the template's logo
	// Example: https://portainer.io/img/logo.svg
	Logo string `json:"logo,omitempty"`

	// Optional stack/container fields
	// Default name for the stack/container to be used on deployment
	// Example: mystackname
	Name string `json:"name,omitempty"`

	// Name of a network that will be used on container deployment if it exists inside the environment(endpoint)
	// Example: mynet
	Network string `json:"network,omitempty"`

	// A note that will be displayed in the UI. Supports HTML content
	// Example: This is my \u003cb\u003ecustom\u003c/b\u003e template
	Note string `json:"note,omitempty"`

	// Platform associated to the template.
	// Valid values are: 'linux', 'windows' or leave empty for multi-platform
	// Example: linux
	Platform string `json:"platform,omitempty"`

	// A list of ports exposed by the container
	// Example: ["8080:80/tcp"]
	Ports []string `json:"ports"`

	// Whether the container should be started in privileged mode
	// Example: true
	Privileged bool `json:"privileged,omitempty"`

	// Optional container fields
	// The URL of a registry associated to the image for a container template
	// Example: quay.io
	Registry string `json:"registry,omitempty"`

	// Mandatory stack fields
	Repository *PortainerTemplateRepository `json:"repository,omitempty"`

	// Container restart policy
	// Example: on-failure
	RestartPolicy string `json:"restart_policy,omitempty"`

	// Mandatory Edge stack fields
	// Stack file used for this template
	StackFile string `json:"stackFile,omitempty"`

	// Title of the template
	// Example: Nginx
	Title string `json:"title,omitempty"`

	// Template type. Valid values are: 1 (container), 2 (Swarm stack), 3 (Compose stack), 4 (Compose edge stack)
	// Example: 1
	Type int64 `json:"type,omitempty"`

	// A list of volumes used during the container template deployment
	Volumes []*PortainerTemplateVolume `json:"volumes"`
}

// Validate validates this portainer template
func (m *PortainerTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerTemplate) validateEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerTemplate) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerTemplate) validateRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerTemplate) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this portainer template based on the context it is used
func (m *PortainerTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortainerTemplate) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {

			if swag.IsZero(m.Env[i]) { // not required
				return nil
			}

			if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerTemplate) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {

			if swag.IsZero(m.Labels[i]) { // not required
				return nil
			}

			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortainerTemplate) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {

		if swag.IsZero(m.Repository) { // not required
			return nil
		}

		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *PortainerTemplate) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {

			if swag.IsZero(m.Volumes[i]) { // not required
				return nil
			}

			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortainerTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortainerTemplate) UnmarshalBinary(b []byte) error {
	var res PortainerTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
