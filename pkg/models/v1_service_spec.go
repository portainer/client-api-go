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

// V1ServiceSpec v1 service spec
//
// swagger:model v1.ServiceSpec
type V1ServiceSpec struct {

	// allocateLoadBalancerNodePorts defines if NodePorts will be automatically
	// allocated for services with type LoadBalancer.  Default is "true". It
	// may be set to "false" if the cluster load-balancer does not rely on
	// NodePorts.  If the caller requests specific NodePorts (by specifying a
	// value), those requests will be respected, regardless of this field.
	// This field may only be set for services with type LoadBalancer and will
	// be cleared if the type is changed to any other type.
	// +optional
	AllocateLoadBalancerNodePorts bool `json:"allocateLoadBalancerNodePorts,omitempty"`

	// clusterIP is the IP address of the service and is usually assigned
	// randomly. If an address is specified manually, is in-range (as per
	// system configuration), and is not in use, it will be allocated to the
	// service; otherwise creation of the service will fail. This field may not
	// be changed through updates unless the type field is also being changed
	// to ExternalName (which requires this field to be blank) or the type
	// field is being changed from ExternalName (in which case this field may
	// optionally be specified, as describe above).  Valid values are "None",
	// empty string (""), or a valid IP address. Setting this to "None" makes a
	// "headless service" (no virtual IP), which is useful when direct endpoint
	// connections are preferred and proxying is not required.  Only applies to
	// types ClusterIP, NodePort, and LoadBalancer. If this field is specified
	// when creating a Service of type ExternalName, creation will fail. This
	// field will be wiped when updating a Service to type ExternalName.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +optional
	ClusterIP string `json:"clusterIP,omitempty"`

	// ClusterIPs is a list of IP addresses assigned to this service, and are
	// usually assigned randomly.  If an address is specified manually, is
	// in-range (as per system configuration), and is not in use, it will be
	// allocated to the service; otherwise creation of the service will fail.
	// This field may not be changed through updates unless the type field is
	// also being changed to ExternalName (which requires this field to be
	// empty) or the type field is being changed from ExternalName (in which
	// case this field may optionally be specified, as describe above).  Valid
	// values are "None", empty string (""), or a valid IP address.  Setting
	// this to "None" makes a "headless service" (no virtual IP), which is
	// useful when direct endpoint connections are preferred and proxying is
	// not required.  Only applies to types ClusterIP, NodePort, and
	// LoadBalancer. If this field is specified when creating a Service of type
	// ExternalName, creation will fail. This field will be wiped when updating
	// a Service to type ExternalName.  If this field is not specified, it will
	// be initialized from the clusterIP field.  If this field is specified,
	// clients must ensure that clusterIPs[0] and clusterIP have the same
	// value.
	//
	// This field may hold a maximum of two entries (dual-stack IPs, in either order).
	// These IPs must correspond to the values of the ipFamilies field. Both
	// clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +listType=atomic
	// +optional
	ClusterIPs []string `json:"clusterIPs"`

	// externalIPs is a list of IP addresses for which nodes in the cluster
	// will also accept traffic for this service.  These IPs are not managed by
	// Kubernetes.  The user is responsible for ensuring that traffic arrives
	// at a node with this IP.  A common example is external load-balancers
	// that are not part of the Kubernetes system.
	// +optional
	// +listType=atomic
	ExternalIPs []string `json:"externalIPs"`

	// externalName is the external reference that discovery mechanisms will
	// return as an alias for this service (e.g. a DNS CNAME record). No
	// proxying will be involved.  Must be a lowercase RFC-1123 hostname
	// (https://tools.ietf.org/html/rfc1123) and requires `type` to be "ExternalName".
	// +optional
	ExternalName string `json:"externalName,omitempty"`

	// externalTrafficPolicy describes how nodes distribute service traffic they
	// receive on one of the Service's "externally-facing" addresses (NodePorts,
	// ExternalIPs, and LoadBalancer IPs). If set to "Local", the proxy will configure
	// the service in a way that assumes that external load balancers will take care
	// of balancing the service traffic between nodes, and so each node will deliver
	// traffic only to the node-local endpoints of the service, without masquerading
	// the client source IP. (Traffic mistakenly sent to a node with no endpoints will
	// be dropped.) The default value, "Cluster", uses the standard behavior of
	// routing to all endpoints evenly (possibly modified by topology and other
	// features). Note that traffic sent to an External IP or LoadBalancer IP from
	// within the cluster will always get "Cluster" semantics, but clients sending to
	// a NodePort from within the cluster may need to take traffic policy into account
	// when picking a node.
	// +optional
	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	// healthCheckNodePort specifies the healthcheck nodePort for the service.
	// This only applies when type is set to LoadBalancer and
	// externalTrafficPolicy is set to Local. If a value is specified, is
	// in-range, and is not in use, it will be used.  If not specified, a value
	// will be automatically allocated.  External systems (e.g. load-balancers)
	// can use this port to determine if a given node holds endpoints for this
	// service or not.  If this field is specified when creating a Service
	// which does not need it, creation will fail. This field will be wiped
	// when updating a Service to no longer need it (e.g. changing type).
	// This field cannot be updated once set.
	// +optional
	HealthCheckNodePort int64 `json:"healthCheckNodePort,omitempty"`

	// InternalTrafficPolicy describes how nodes distribute service traffic they
	// receive on the ClusterIP. If set to "Local", the proxy will assume that pods
	// only want to talk to endpoints of the service on the same node as the pod,
	// dropping the traffic if there are no local endpoints. The default value,
	// "Cluster", uses the standard behavior of routing to all endpoints evenly
	// (possibly modified by topology and other features).
	// +optional
	InternalTrafficPolicy string `json:"internalTrafficPolicy,omitempty"`

	// IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this
	// service. This field is usually assigned automatically based on cluster
	// configuration and the ipFamilyPolicy field. If this field is specified
	// manually, the requested family is available in the cluster,
	// and ipFamilyPolicy allows it, it will be used; otherwise creation of
	// the service will fail. This field is conditionally mutable: it allows
	// for adding or removing a secondary IP family, but it does not allow
	// changing the primary IP family of the Service. Valid values are "IPv4"
	// and "IPv6".  This field only applies to Services of types ClusterIP,
	// NodePort, and LoadBalancer, and does apply to "headless" services.
	// This field will be wiped when updating a Service to type ExternalName.
	//
	// This field may hold a maximum of two entries (dual-stack families, in
	// either order).  These families must correspond to the values of the
	// clusterIPs field, if specified. Both clusterIPs and ipFamilies are
	// governed by the ipFamilyPolicy field.
	// +listType=atomic
	// +optional
	IPFamilies []string `json:"ipFamilies"`

	// IPFamilyPolicy represents the dual-stack-ness requested or required by
	// this Service. If there is no value provided, then this field will be set
	// to SingleStack. Services can be "SingleStack" (a single IP family),
	// "PreferDualStack" (two IP families on dual-stack configured clusters or
	// a single IP family on single-stack clusters), or "RequireDualStack"
	// (two IP families on dual-stack configured clusters, otherwise fail). The
	// ipFamilies and clusterIPs fields depend on the value of this field. This
	// field will be wiped when updating a service to type ExternalName.
	// +optional
	IPFamilyPolicy string `json:"ipFamilyPolicy,omitempty"`

	// loadBalancerClass is the class of the load balancer implementation this Service belongs to.
	// If specified, the value of this field must be a label-style identifier, with an optional prefix,
	// e.g. "internal-vip" or "example.com/internal-vip". Unprefixed names are reserved for end-users.
	// This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load
	// balancer implementation is used, today this is typically done through the cloud provider integration,
	// but should apply for any default implementation. If set, it is assumed that a load balancer
	// implementation is watching for Services with a matching class. Any default load balancer
	// implementation (e.g. cloud providers) should ignore Services that set this field.
	// This field can only be set when creating or updating a Service to type 'LoadBalancer'.
	// Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.
	// +optional
	LoadBalancerClass string `json:"loadBalancerClass,omitempty"`

	// Only applies to Service Type: LoadBalancer.
	// This feature depends on whether the underlying cloud-provider supports specifying
	// the loadBalancerIP when a load balancer is created.
	// This field will be ignored if the cloud-provider does not support the feature.
	// Deprecated: This field was under-specified and its meaning varies across implementations.
	// Using it is non-portable and it may not support dual-stack.
	// Users are encouraged to use implementation-specific annotations when available.
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`

	// If specified and supported by the platform, this will restrict traffic through the cloud-provider
	// load-balancer will be restricted to the specified client IPs. This field will be ignored if the
	// cloud-provider does not support the feature."
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/
	// +optional
	// +listType=atomic
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges"`

	// The list of ports that are exposed by this service.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +patchMergeKey=port
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=port
	// +listMapKey=protocol
	Ports []*V1ServicePort `json:"ports"`

	// publishNotReadyAddresses indicates that any agent which deals with endpoints for this
	// Service should disregard any indications of ready/not-ready.
	// The primary use case for setting this field is for a StatefulSet's Headless Service to
	// propagate SRV DNS records for its Pods for the purpose of peer discovery.
	// The Kubernetes controllers that generate Endpoints and EndpointSlice resources for
	// Services interpret this to mean that all endpoints are considered "ready" even if the
	// Pods themselves are not. Agents which consume only Kubernetes generated endpoints
	// through the Endpoints or EndpointSlice resources can safely assume this behavior.
	// +optional
	PublishNotReadyAddresses bool `json:"publishNotReadyAddresses,omitempty"`

	// Route service traffic to pods with label keys and values matching this
	// selector. If empty or not present, the service is assumed to have an
	// external process managing its endpoints, which Kubernetes will not
	// modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
	// Ignored if type is ExternalName.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/
	// +optional
	// +mapType=atomic
	Selector map[string]string `json:"selector,omitempty"`

	// Supports "ClientIP" and "None". Used to maintain session affinity.
	// Enable client IP based session affinity.
	// Must be ClientIP or None.
	// Defaults to None.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +optional
	SessionAffinity string `json:"sessionAffinity,omitempty"`

	// sessionAffinityConfig contains the configurations of session affinity.
	// +optional
	SessionAffinityConfig *V1SessionAffinityConfig `json:"sessionAffinityConfig,omitempty"`

	// TrafficDistribution offers a way to express preferences for how traffic is
	// distributed to Service endpoints. Implementations can use this field as a
	// hint, but are not required to guarantee strict adherence. If the field is
	// not set, the implementation will apply its default routing strategy. If set
	// to "PreferClose", implementations should prioritize endpoints that are
	// topologically close (e.g., same zone).
	// This is a beta field and requires enabling ServiceTrafficDistribution feature.
	// +featureGate=ServiceTrafficDistribution
	// +optional
	TrafficDistribution string `json:"trafficDistribution,omitempty"`

	// type determines how the Service is exposed. Defaults to ClusterIP. Valid
	// options are ExternalName, ClusterIP, NodePort, and LoadBalancer.
	// "ClusterIP" allocates a cluster-internal IP address for load-balancing
	// to endpoints. Endpoints are determined by the selector or if that is not
	// specified, by manual construction of an Endpoints object or
	// EndpointSlice objects. If clusterIP is "None", no virtual IP is
	// allocated and the endpoints are published as a set of endpoints rather
	// than a virtual IP.
	// "NodePort" builds on ClusterIP and allocates a port on every node which
	// routes to the same endpoints as the clusterIP.
	// "LoadBalancer" builds on NodePort and creates an external load-balancer
	// (if supported in the current cloud) which routes to the same endpoints
	// as the clusterIP.
	// "ExternalName" aliases this service to the specified externalName.
	// Several other fields do not apply to ExternalName services.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// +optional
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 service spec
func (m *V1ServiceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionAffinityConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServiceSpec) validatePorts(formats strfmt.Registry) error {
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

func (m *V1ServiceSpec) validateSessionAffinityConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionAffinityConfig) { // not required
		return nil
	}

	if m.SessionAffinityConfig != nil {
		if err := m.SessionAffinityConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sessionAffinityConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sessionAffinityConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 service spec based on the context it is used
func (m *V1ServiceSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionAffinityConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServiceSpec) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {

			if swag.IsZero(m.Ports[i]) { // not required
				return nil
			}

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

func (m *V1ServiceSpec) contextValidateSessionAffinityConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SessionAffinityConfig != nil {

		if swag.IsZero(m.SessionAffinityConfig) { // not required
			return nil
		}

		if err := m.SessionAffinityConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sessionAffinityConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sessionAffinityConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ServiceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServiceSpec) UnmarshalBinary(b []byte) error {
	var res V1ServiceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
