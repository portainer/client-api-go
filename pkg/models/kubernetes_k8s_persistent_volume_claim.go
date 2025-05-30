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

// KubernetesK8sPersistentVolumeClaim kubernetes k8s persistent volume claim
//
// swagger:model kubernetes.K8sPersistentVolumeClaim
type KubernetesK8sPersistentVolumeClaim struct {

	// access modes
	AccessModes []string `json:"accessModes"`

	// creation date
	CreationDate string `json:"creationDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// owning applications
	OwningApplications []*KubernetesK8sApplication `json:"owningApplications"`

	// phase
	Phase string `json:"phase,omitempty"`

	// resources requests
	ResourcesRequests V1ResourceList `json:"resourcesRequests,omitempty"`

	// storage
	Storage int64 `json:"storage,omitempty"`

	// storage class
	StorageClass string `json:"storageClass,omitempty"`

	// volume mode
	VolumeMode string `json:"volumeMode,omitempty"`

	// volume name
	VolumeName string `json:"volumeName,omitempty"`
}

// Validate validates this kubernetes k8s persistent volume claim
func (m *KubernetesK8sPersistentVolumeClaim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwningApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcesRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesK8sPersistentVolumeClaim) validateOwningApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.OwningApplications) { // not required
		return nil
	}

	for i := 0; i < len(m.OwningApplications); i++ {
		if swag.IsZero(m.OwningApplications[i]) { // not required
			continue
		}

		if m.OwningApplications[i] != nil {
			if err := m.OwningApplications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owningApplications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("owningApplications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesK8sPersistentVolumeClaim) validateResourcesRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcesRequests) { // not required
		return nil
	}

	if m.ResourcesRequests != nil {
		if err := m.ResourcesRequests.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcesRequests")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcesRequests")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kubernetes k8s persistent volume claim based on the context it is used
func (m *KubernetesK8sPersistentVolumeClaim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOwningApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourcesRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesK8sPersistentVolumeClaim) contextValidateOwningApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OwningApplications); i++ {

		if m.OwningApplications[i] != nil {

			if swag.IsZero(m.OwningApplications[i]) { // not required
				return nil
			}

			if err := m.OwningApplications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owningApplications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("owningApplications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesK8sPersistentVolumeClaim) contextValidateResourcesRequests(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ResourcesRequests) { // not required
		return nil
	}

	if err := m.ResourcesRequests.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resourcesRequests")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("resourcesRequests")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesK8sPersistentVolumeClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesK8sPersistentVolumeClaim) UnmarshalBinary(b []byte) error {
	var res KubernetesK8sPersistentVolumeClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
