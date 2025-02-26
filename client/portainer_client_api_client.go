// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/client/auth"
	"github.com/portainer/client-api-go/v2/client/backup"
	"github.com/portainer/client-api-go/v2/client/chat"
	"github.com/portainer/client-api-go/v2/client/cloud_credentials"
	"github.com/portainer/client-api-go/v2/client/custom_templates"
	"github.com/portainer/client-api-go/v2/client/docker"
	"github.com/portainer/client-api-go/v2/client/edge"
	"github.com/portainer/client-api-go/v2/client/edge_configs"
	"github.com/portainer/client-api-go/v2/client/edge_groups"
	"github.com/portainer/client-api-go/v2/client/edge_jobs"
	"github.com/portainer/client-api-go/v2/client/edge_stacks"
	"github.com/portainer/client-api-go/v2/client/edge_update_schedules"
	"github.com/portainer/client-api-go/v2/client/endpoint_groups"
	"github.com/portainer/client-api-go/v2/client/endpoints"
	"github.com/portainer/client-api-go/v2/client/gitops"
	"github.com/portainer/client-api-go/v2/client/helm"
	"github.com/portainer/client-api-go/v2/client/intel"
	"github.com/portainer/client-api-go/v2/client/kaas"
	"github.com/portainer/client-api-go/v2/client/kubernetes"
	"github.com/portainer/client-api-go/v2/client/ldap"
	"github.com/portainer/client-api-go/v2/client/license"
	"github.com/portainer/client-api-go/v2/client/motd"
	"github.com/portainer/client-api-go/v2/client/registries"
	"github.com/portainer/client-api-go/v2/client/resource_controls"
	"github.com/portainer/client-api-go/v2/client/roles"
	"github.com/portainer/client-api-go/v2/client/settings"
	"github.com/portainer/client-api-go/v2/client/ssl"
	"github.com/portainer/client-api-go/v2/client/stacks"
	"github.com/portainer/client-api-go/v2/client/status"
	"github.com/portainer/client-api-go/v2/client/support"
	"github.com/portainer/client-api-go/v2/client/system"
	"github.com/portainer/client-api-go/v2/client/tags"
	"github.com/portainer/client-api-go/v2/client/team_memberships"
	"github.com/portainer/client-api-go/v2/client/teams"
	"github.com/portainer/client-api-go/v2/client/templates"
	"github.com/portainer/client-api-go/v2/client/upload"
	"github.com/portainer/client-api-go/v2/client/useractivity"
	"github.com/portainer/client-api-go/v2/client/users"
	"github.com/portainer/client-api-go/v2/client/webhooks"
	"github.com/portainer/client-api-go/v2/client/websocket"
)

// Default portainer client API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new portainer client API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PortainerClientAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new portainer client API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PortainerClientAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new portainer client API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PortainerClientAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PortainerClientAPI)
	cli.Transport = transport
	cli.Auth = auth.New(transport, formats)
	cli.Backup = backup.New(transport, formats)
	cli.Chat = chat.New(transport, formats)
	cli.CloudCredentials = cloud_credentials.New(transport, formats)
	cli.CustomTemplates = custom_templates.New(transport, formats)
	cli.Docker = docker.New(transport, formats)
	cli.Edge = edge.New(transport, formats)
	cli.EdgeConfigs = edge_configs.New(transport, formats)
	cli.EdgeGroups = edge_groups.New(transport, formats)
	cli.EdgeJobs = edge_jobs.New(transport, formats)
	cli.EdgeStacks = edge_stacks.New(transport, formats)
	cli.EdgeUpdateSchedules = edge_update_schedules.New(transport, formats)
	cli.EndpointGroups = endpoint_groups.New(transport, formats)
	cli.Endpoints = endpoints.New(transport, formats)
	cli.Gitops = gitops.New(transport, formats)
	cli.Helm = helm.New(transport, formats)
	cli.Intel = intel.New(transport, formats)
	cli.Kaas = kaas.New(transport, formats)
	cli.Kubernetes = kubernetes.New(transport, formats)
	cli.Ldap = ldap.New(transport, formats)
	cli.License = license.New(transport, formats)
	cli.Motd = motd.New(transport, formats)
	cli.Registries = registries.New(transport, formats)
	cli.ResourceControls = resource_controls.New(transport, formats)
	cli.Roles = roles.New(transport, formats)
	cli.Settings = settings.New(transport, formats)
	cli.Ssl = ssl.New(transport, formats)
	cli.Stacks = stacks.New(transport, formats)
	cli.Status = status.New(transport, formats)
	cli.Support = support.New(transport, formats)
	cli.System = system.New(transport, formats)
	cli.Tags = tags.New(transport, formats)
	cli.TeamMemberships = team_memberships.New(transport, formats)
	cli.Teams = teams.New(transport, formats)
	cli.Templates = templates.New(transport, formats)
	cli.Upload = upload.New(transport, formats)
	cli.Useractivity = useractivity.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Webhooks = webhooks.New(transport, formats)
	cli.Websocket = websocket.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PortainerClientAPI is a client for portainer client API
type PortainerClientAPI struct {
	Auth auth.ClientService

	Backup backup.ClientService

	Chat chat.ClientService

	CloudCredentials cloud_credentials.ClientService

	CustomTemplates custom_templates.ClientService

	Docker docker.ClientService

	Edge edge.ClientService

	EdgeConfigs edge_configs.ClientService

	EdgeGroups edge_groups.ClientService

	EdgeJobs edge_jobs.ClientService

	EdgeStacks edge_stacks.ClientService

	EdgeUpdateSchedules edge_update_schedules.ClientService

	EndpointGroups endpoint_groups.ClientService

	Endpoints endpoints.ClientService

	Gitops gitops.ClientService

	Helm helm.ClientService

	Intel intel.ClientService

	Kaas kaas.ClientService

	Kubernetes kubernetes.ClientService

	Ldap ldap.ClientService

	License license.ClientService

	Motd motd.ClientService

	Registries registries.ClientService

	ResourceControls resource_controls.ClientService

	Roles roles.ClientService

	Settings settings.ClientService

	Ssl ssl.ClientService

	Stacks stacks.ClientService

	Status status.ClientService

	Support support.ClientService

	System system.ClientService

	Tags tags.ClientService

	TeamMemberships team_memberships.ClientService

	Teams teams.ClientService

	Templates templates.ClientService

	Upload upload.ClientService

	Useractivity useractivity.ClientService

	Users users.ClientService

	Webhooks webhooks.ClientService

	Websocket websocket.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PortainerClientAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Auth.SetTransport(transport)
	c.Backup.SetTransport(transport)
	c.Chat.SetTransport(transport)
	c.CloudCredentials.SetTransport(transport)
	c.CustomTemplates.SetTransport(transport)
	c.Docker.SetTransport(transport)
	c.Edge.SetTransport(transport)
	c.EdgeConfigs.SetTransport(transport)
	c.EdgeGroups.SetTransport(transport)
	c.EdgeJobs.SetTransport(transport)
	c.EdgeStacks.SetTransport(transport)
	c.EdgeUpdateSchedules.SetTransport(transport)
	c.EndpointGroups.SetTransport(transport)
	c.Endpoints.SetTransport(transport)
	c.Gitops.SetTransport(transport)
	c.Helm.SetTransport(transport)
	c.Intel.SetTransport(transport)
	c.Kaas.SetTransport(transport)
	c.Kubernetes.SetTransport(transport)
	c.Ldap.SetTransport(transport)
	c.License.SetTransport(transport)
	c.Motd.SetTransport(transport)
	c.Registries.SetTransport(transport)
	c.ResourceControls.SetTransport(transport)
	c.Roles.SetTransport(transport)
	c.Settings.SetTransport(transport)
	c.Ssl.SetTransport(transport)
	c.Stacks.SetTransport(transport)
	c.Status.SetTransport(transport)
	c.Support.SetTransport(transport)
	c.System.SetTransport(transport)
	c.Tags.SetTransport(transport)
	c.TeamMemberships.SetTransport(transport)
	c.Teams.SetTransport(transport)
	c.Templates.SetTransport(transport)
	c.Upload.SetTransport(transport)
	c.Useractivity.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Webhooks.SetTransport(transport)
	c.Websocket.SetTransport(transport)
}
