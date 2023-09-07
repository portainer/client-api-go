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

// ModelsK8sServiceInfo models k8s service info
//
// swagger:model models.K8sServiceInfo
type ModelsK8sServiceInfo struct {

	// allocate load balancer node ports
	AllocateLoadBalancerNodePorts bool `json:"allocateLoadBalancerNodePorts,omitempty"`

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// serviceList screen
	Applications []*ModelsK8sApplication `json:"applications"`

	// cluster i ps
	ClusterIPs []string `json:"clusterIPs"`

	// creation timestamp
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// external i ps
	ExternalIPs []string `json:"externalIPs"`

	// external name
	ExternalName string `json:"externalName,omitempty"`

	// ingress status
	IngressStatus []*ModelsK8sServiceIngress `json:"ingressStatus"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// ports
	Ports []*ModelsK8sServicePort `json:"ports"`

	// selector
	Selector map[string]string `json:"selector,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this models k8s service info
func (m *ModelsK8sServiceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsK8sServiceInfo) validateApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {
		if swag.IsZero(m.Applications[i]) { // not required
			continue
		}

		if m.Applications[i] != nil {
			if err := m.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsK8sServiceInfo) validateIngressStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.IngressStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.IngressStatus); i++ {
		if swag.IsZero(m.IngressStatus[i]) { // not required
			continue
		}

		if m.IngressStatus[i] != nil {
			if err := m.IngressStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingressStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingressStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsK8sServiceInfo) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this models k8s service info based on the context it is used
func (m *ModelsK8sServiceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngressStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsK8sServiceInfo) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Applications); i++ {

		if m.Applications[i] != nil {
			if err := m.Applications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsK8sServiceInfo) contextValidateIngressStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IngressStatus); i++ {

		if m.IngressStatus[i] != nil {
			if err := m.IngressStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingressStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingressStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsK8sServiceInfo) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsK8sServiceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsK8sServiceInfo) UnmarshalBinary(b []byte) error {
	var res ModelsK8sServiceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
