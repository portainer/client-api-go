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

// V1ServicePort v1 service port
//
// swagger:model v1.ServicePort
type V1ServicePort struct {

	// The application protocol for this port.
	// This is used as a hint for implementations to offer richer behavior for protocols that they understand.
	// This field follows standard Kubernetes label syntax.
	// Valid values are either:
	//
	// * Un-prefixed protocol names - reserved for IANA standard service names (as per
	// RFC-6335 and https://www.iana.org/assignments/service-names).
	//
	// * Kubernetes-defined prefixed names:
	//   * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
	//   * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
	//   * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
	//
	// * Other protocols should use implementation-defined prefixed names such as
	// mycompany.com/my-custom-protocol.
	// +optional
	AppProtocol string `json:"appProtocol,omitempty"`

	// The name of this port within the service. This must be a DNS_LABEL.
	// All ports within a ServiceSpec must have unique names. When considering
	// the endpoints for a Service, this must match the 'name' field in the
	// EndpointPort.
	// Optional if only one ServicePort is defined on this service.
	// +optional
	Name string `json:"name,omitempty"`

	// The port on each node on which this service is exposed when type is
	// NodePort or LoadBalancer.  Usually assigned by the system. If a value is
	// specified, in-range, and not in use it will be used, otherwise the
	// operation will fail.  If not specified, a port will be allocated if this
	// Service requires one.  If this field is specified when creating a
	// Service which does not need it, creation will fail. This field will be
	// wiped when updating a Service to no longer need it (e.g. changing type
	// from NodePort to ClusterIP).
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	// +optional
	NodePort int64 `json:"nodePort,omitempty"`

	// The port that will be exposed by this service.
	Port int64 `json:"port,omitempty"`

	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".
	// Default is TCP.
	// +default="TCP"
	// +optional
	Protocol string `json:"protocol,omitempty"`

	// Number or name of the port to access on the pods targeted by the service.
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// If this is a string, it will be looked up as a named port in the
	// target Pod's container ports. If this is not specified, the value
	// of the 'port' field is used (an identity map).
	// This field is ignored for services with clusterIP=None, and should be
	// omitted or set equal to the 'port' field.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	// +optional
	TargetPort *IntstrIntOrString `json:"targetPort,omitempty"`
}

// Validate validates this v1 service port
func (m *V1ServicePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServicePort) validateTargetPort(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetPort) { // not required
		return nil
	}

	if m.TargetPort != nil {
		if err := m.TargetPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetPort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetPort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 service port based on the context it is used
func (m *V1ServicePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServicePort) contextValidateTargetPort(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetPort != nil {

		if swag.IsZero(m.TargetPort) { // not required
			return nil
		}

		if err := m.TargetPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetPort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetPort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ServicePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServicePort) UnmarshalBinary(b []byte) error {
	var res V1ServicePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
