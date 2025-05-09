// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesK8sPersistentVolume kubernetes k8s persistent volume
//
// swagger:model kubernetes.K8sPersistentVolume
type KubernetesK8sPersistentVolume struct {

	// access modes
	AccessModes []string `json:"accessModes"`

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// capacity
	Capacity V1ResourceList `json:"capacity,omitempty"`

	// claim ref
	ClaimRef *V1ObjectReference `json:"claimRef,omitempty"`

	// csi
	Csi *V1CSIPersistentVolumeSource `json:"csi,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// persistent volume reclaim policy
	PersistentVolumeReclaimPolicy string `json:"persistentVolumeReclaimPolicy,omitempty"`

	// storage class name
	StorageClassName string `json:"storageClassName,omitempty"`

	// volume mode
	VolumeMode string `json:"volumeMode,omitempty"`
}

// Validate validates this kubernetes k8s persistent volume
func (m *KubernetesK8sPersistentVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesK8sPersistentVolume) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesK8sPersistentVolume) validateClaimRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ClaimRef) { // not required
		return nil
	}

	if m.ClaimRef != nil {
		if err := m.ClaimRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claimRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("claimRef")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesK8sPersistentVolume) validateCsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Csi) { // not required
		return nil
	}

	if m.Csi != nil {
		if err := m.Csi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kubernetes k8s persistent volume based on the context it is used
func (m *KubernetesK8sPersistentVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClaimRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesK8sPersistentVolume) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capacity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("capacity")
		}
		return err
	}

	return nil
}

func (m *KubernetesK8sPersistentVolume) contextValidateClaimRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ClaimRef != nil {

		if swag.IsZero(m.ClaimRef) { // not required
			return nil
		}

		if err := m.ClaimRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claimRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("claimRef")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesK8sPersistentVolume) contextValidateCsi(ctx context.Context, formats strfmt.Registry) error {

	if m.Csi != nil {

		if swag.IsZero(m.Csi) { // not required
			return nil
		}

		if err := m.Csi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesK8sPersistentVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesK8sPersistentVolume) UnmarshalBinary(b []byte) error {
	var res KubernetesK8sPersistentVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
