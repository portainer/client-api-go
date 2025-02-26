// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new kaas API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new kaas API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new kaas API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for kaas API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddNodes(params *AddNodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddNodesOK, error)

	KaasVersion(params *KaasVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KaasVersionOK, error)

	Microk8sGetAddons(params *Microk8sGetAddonsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Microk8sGetAddonsOK, error)

	Microk8sGetNodeStatus(params *Microk8sGetNodeStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Microk8sGetNodeStatusOK, error)

	Microk8sUpdateAddons(params *Microk8sUpdateAddonsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Microk8sUpdateAddonsOK, error)

	ProviderInfo(params *ProviderInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProviderInfoOK, error)

	ProvisionCluster(params *ProvisionClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterOK, error)

	ProvisionClusterAmazon(params *ProvisionClusterAmazonParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterAmazonOK, error)

	ProvisionClusterAzure(params *ProvisionClusterAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterAzureOK, error)

	ProvisionClusterGKE(params *ProvisionClusterGKEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterGKEOK, error)

	RemoveNodes(params *RemoveNodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveNodesOK, error)

	TestSSH(params *TestSSHParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TestSSHOK, error)

	Upgrade(params *UpgradeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpgradeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AddNodes adds nodes to the cluster scale up

	Add control plane and worker nodes to the cluster.

**Access policy**: authenticated
*/
func (a *Client) AddNodes(params *AddNodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addNodes",
		Method:             "POST",
		PathPattern:        "/cloud/endpoints/{environmentId}/nodes/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addNodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	KaasVersion gets the current cluster version

	Get the current cluster version.

**Access policy**: authenticated
*/
func (a *Client) KaasVersion(params *KaasVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KaasVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKaasVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "kaasVersion",
		Method:             "GET",
		PathPattern:        "/cloud/endpoints/{environmentId}/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KaasVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KaasVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for kaasVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	Microk8sGetAddons gets a list of addons which are installed in a micro k8s cluster

	The information returned can be used to query the MircoK8s cluster.

**Access policy**: authenticated
*/
func (a *Client) Microk8sGetAddons(params *Microk8sGetAddonsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Microk8sGetAddonsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMicrok8sGetAddonsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "microk8sGetAddons",
		Method:             "GET",
		PathPattern:        "/cloud/endpoints/{environmentId}/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Microk8sGetAddonsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Microk8sGetAddonsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for microk8sGetAddons: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	Microk8sGetNodeStatus gets the micro k8s status for a control plane node

	Get the MicroK8s status for a control plane node in a MicroK8s cluster.

**Access policy**: authenticated
*/
func (a *Client) Microk8sGetNodeStatus(params *Microk8sGetNodeStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Microk8sGetNodeStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMicrok8sGetNodeStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "microk8sGetNodeStatus",
		Method:             "GET",
		PathPattern:        "/cloud/endpoints/{environmentId}/nodes/nodestatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Microk8sGetNodeStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Microk8sGetNodeStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for microk8sGetNodeStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	Microk8sUpdateAddons gets a list of addons which are installed in a micro k8s cluster

	The information returned can be used to query the MircoK8s cluster.

**Access policy**: authenticated
*/
func (a *Client) Microk8sUpdateAddons(params *Microk8sUpdateAddonsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Microk8sUpdateAddonsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMicrok8sUpdateAddonsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "microk8sUpdateAddons",
		Method:             "POST",
		PathPattern:        "/cloud/endpoints/{environmentId}/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Microk8sUpdateAddonsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Microk8sUpdateAddonsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for microk8sUpdateAddons: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ProviderInfo gets information about the provisioning options for a cloud provider

	The information returned can be used to provision a KaaS cluster.

**Access policy**: administrator
*/
func (a *Client) ProviderInfo(params *ProviderInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProviderInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProviderInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "providerInfo",
		Method:             "GET",
		PathPattern:        "/cloud/{provider}/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProviderInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProviderInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for providerInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ProvisionCluster provisions a new c i v o linode or digital ocean cluster and create an environment

	Provision a new KaaS cluster and create an environment.

This documentation is specifically for civo, digitial ocean and linode.

For Azure, GKE and Amazon see:
*[*]/cloud/amazon/provision**
*[*]/cloud/azure/provision**
*[*]/cloud/gke/provision**

**Access policy**: administrator
*/
func (a *Client) ProvisionCluster(params *ProvisionClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "provisionCluster",
		Method:             "POST",
		PathPattern:        "/cloud/{provider}/provision",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProvisionClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProvisionClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for provisionCluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ProvisionClusterAmazon provisions a new amazon e k s and create an environment

	Provision a new KaaS cluster and create an environment.

**Access policy**: administrator
*/
func (a *Client) ProvisionClusterAmazon(params *ProvisionClusterAmazonParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterAmazonOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionClusterAmazonParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "provisionClusterAmazon",
		Method:             "POST",
		PathPattern:        "/cloud/amazon/provision",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProvisionClusterAmazonReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProvisionClusterAmazonOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for provisionClusterAmazon: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ProvisionClusterAzure provisions a new microsoft azure cluster and create an environment

	Provision a new KaaS cluster and create an environment.

**Access policy**: administrator
*/
func (a *Client) ProvisionClusterAzure(params *ProvisionClusterAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterAzureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionClusterAzureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "provisionClusterAzure",
		Method:             "POST",
		PathPattern:        "/cloud/azure/provision",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProvisionClusterAzureReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProvisionClusterAzureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for provisionClusterAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ProvisionClusterGKE provisions a new google kubernetes g k e cluster and create an environment

	Provision a new KaaS cluster and create an environment.

**Access policy**: administrator
*/
func (a *Client) ProvisionClusterGKE(params *ProvisionClusterGKEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProvisionClusterGKEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionClusterGKEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "provisionClusterGKE",
		Method:             "POST",
		PathPattern:        "/cloud/gke/provision",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProvisionClusterGKEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProvisionClusterGKEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for provisionClusterGKE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RemoveNodes removes nodes from the cluster and uninstall micro k8s from them

	Remove nodes from the cluster and uninstall MicroK8s from them.

**Access policy**: authenticated
*/
func (a *Client) RemoveNodes(params *RemoveNodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveNodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeNodes",
		Method:             "POST",
		PathPattern:        "/cloud/endpoints/{environmentId}/nodes/remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeNodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	TestSSH tests SSH connection to nodes

	Test SSH connection to nodes.

**Access policy**: administrator
*/
func (a *Client) TestSSH(params *TestSSHParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TestSSHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestSSHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "testSSH",
		Method:             "POST",
		PathPattern:        "/cloud/testssh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestSSHReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestSSHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testSSH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	Upgrade upgrades the cluster to the next stable kubernetes version

	Upgrade the cluster to the next stable kubernetes version.

**Access policy**: authenticated
*/
func (a *Client) Upgrade(params *UpgradeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpgradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upgrade",
		Method:             "POST",
		PathPattern:        "/cloud/endpoints/{environmentId}/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpgradeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upgrade: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
