// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewEndpointListParams creates a new EndpointListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEndpointListParams() *EndpointListParams {
	return &EndpointListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEndpointListParamsWithTimeout creates a new EndpointListParams object
// with the ability to set a timeout on a request.
func NewEndpointListParamsWithTimeout(timeout time.Duration) *EndpointListParams {
	return &EndpointListParams{
		timeout: timeout,
	}
}

// NewEndpointListParamsWithContext creates a new EndpointListParams object
// with the ability to set a context for a request.
func NewEndpointListParamsWithContext(ctx context.Context) *EndpointListParams {
	return &EndpointListParams{
		Context: ctx,
	}
}

// NewEndpointListParamsWithHTTPClient creates a new EndpointListParams object
// with the ability to set a custom HTTPClient for a request.
func NewEndpointListParamsWithHTTPClient(client *http.Client) *EndpointListParams {
	return &EndpointListParams{
		HTTPClient: client,
	}
}

/*
EndpointListParams contains all the parameters to send to the API endpoint

	for the endpoint list operation.

	Typically these are written to a http.Request.
*/
type EndpointListParams struct {

	/* AgentVersions.

	   will return only environments with on of these agent versions
	*/
	AgentVersions []string

	/* EdgeDevice.

	   if exists true show only edge devices, false show only regular edge endpoints. if missing, will show both types (relevant only for edge endpoints)
	*/
	EdgeDevice *bool

	/* EdgeDeviceUntrusted.

	   if true, show only untrusted endpoints, if false show only trusted (relevant only for edge devices, and if edgeDevice is true)
	*/
	EdgeDeviceUntrusted *bool

	/* EndpointIds.

	   will return only these environments(endpoints)
	*/
	EndpointIds []int64

	/* GroupIds.

	   List environments(endpoints) of these groups
	*/
	GroupIds []int64

	/* Limit.

	   Limit results to this value
	*/
	Limit *int64

	/* Name.

	   will return only environments(endpoints) with this name
	*/
	Name *string

	/* Order.

	   Order sorted results by desc/asc
	*/
	Order *int64

	/* Provisioned.

	   If true, will return environment(endpoint) that were provisioned
	*/
	Provisioned *bool

	/* Search.

	   Search query
	*/
	Search *string

	/* Sort.

	   Sort results by this value
	*/
	Sort *int64

	/* Start.

	   Start searching from
	*/
	Start *int64

	/* Status.

	   List environments(endpoints) by this status
	*/
	Status []int64

	/* TagIds.

	   search environments(endpoints) with these tags (depends on tagsPartialMatch)
	*/
	TagIds []int64

	/* TagsPartialMatch.

	   If true, will return environment(endpoint) which has one of tagIds, if false (or missing) will return only environments(endpoints) that has all the tags
	*/
	TagsPartialMatch *bool

	/* Types.

	   List environments(endpoints) of this type
	*/
	Types []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the endpoint list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointListParams) WithDefaults() *EndpointListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the endpoint list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EndpointListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the endpoint list params
func (o *EndpointListParams) WithTimeout(timeout time.Duration) *EndpointListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the endpoint list params
func (o *EndpointListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the endpoint list params
func (o *EndpointListParams) WithContext(ctx context.Context) *EndpointListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the endpoint list params
func (o *EndpointListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the endpoint list params
func (o *EndpointListParams) WithHTTPClient(client *http.Client) *EndpointListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the endpoint list params
func (o *EndpointListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentVersions adds the agentVersions to the endpoint list params
func (o *EndpointListParams) WithAgentVersions(agentVersions []string) *EndpointListParams {
	o.SetAgentVersions(agentVersions)
	return o
}

// SetAgentVersions adds the agentVersions to the endpoint list params
func (o *EndpointListParams) SetAgentVersions(agentVersions []string) {
	o.AgentVersions = agentVersions
}

// WithEdgeDevice adds the edgeDevice to the endpoint list params
func (o *EndpointListParams) WithEdgeDevice(edgeDevice *bool) *EndpointListParams {
	o.SetEdgeDevice(edgeDevice)
	return o
}

// SetEdgeDevice adds the edgeDevice to the endpoint list params
func (o *EndpointListParams) SetEdgeDevice(edgeDevice *bool) {
	o.EdgeDevice = edgeDevice
}

// WithEdgeDeviceUntrusted adds the edgeDeviceUntrusted to the endpoint list params
func (o *EndpointListParams) WithEdgeDeviceUntrusted(edgeDeviceUntrusted *bool) *EndpointListParams {
	o.SetEdgeDeviceUntrusted(edgeDeviceUntrusted)
	return o
}

// SetEdgeDeviceUntrusted adds the edgeDeviceUntrusted to the endpoint list params
func (o *EndpointListParams) SetEdgeDeviceUntrusted(edgeDeviceUntrusted *bool) {
	o.EdgeDeviceUntrusted = edgeDeviceUntrusted
}

// WithEndpointIds adds the endpointIds to the endpoint list params
func (o *EndpointListParams) WithEndpointIds(endpointIds []int64) *EndpointListParams {
	o.SetEndpointIds(endpointIds)
	return o
}

// SetEndpointIds adds the endpointIds to the endpoint list params
func (o *EndpointListParams) SetEndpointIds(endpointIds []int64) {
	o.EndpointIds = endpointIds
}

// WithGroupIds adds the groupIds to the endpoint list params
func (o *EndpointListParams) WithGroupIds(groupIds []int64) *EndpointListParams {
	o.SetGroupIds(groupIds)
	return o
}

// SetGroupIds adds the groupIds to the endpoint list params
func (o *EndpointListParams) SetGroupIds(groupIds []int64) {
	o.GroupIds = groupIds
}

// WithLimit adds the limit to the endpoint list params
func (o *EndpointListParams) WithLimit(limit *int64) *EndpointListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the endpoint list params
func (o *EndpointListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the endpoint list params
func (o *EndpointListParams) WithName(name *string) *EndpointListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the endpoint list params
func (o *EndpointListParams) SetName(name *string) {
	o.Name = name
}

// WithOrder adds the order to the endpoint list params
func (o *EndpointListParams) WithOrder(order *int64) *EndpointListParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the endpoint list params
func (o *EndpointListParams) SetOrder(order *int64) {
	o.Order = order
}

// WithProvisioned adds the provisioned to the endpoint list params
func (o *EndpointListParams) WithProvisioned(provisioned *bool) *EndpointListParams {
	o.SetProvisioned(provisioned)
	return o
}

// SetProvisioned adds the provisioned to the endpoint list params
func (o *EndpointListParams) SetProvisioned(provisioned *bool) {
	o.Provisioned = provisioned
}

// WithSearch adds the search to the endpoint list params
func (o *EndpointListParams) WithSearch(search *string) *EndpointListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the endpoint list params
func (o *EndpointListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSort adds the sort to the endpoint list params
func (o *EndpointListParams) WithSort(sort *int64) *EndpointListParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the endpoint list params
func (o *EndpointListParams) SetSort(sort *int64) {
	o.Sort = sort
}

// WithStart adds the start to the endpoint list params
func (o *EndpointListParams) WithStart(start *int64) *EndpointListParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the endpoint list params
func (o *EndpointListParams) SetStart(start *int64) {
	o.Start = start
}

// WithStatus adds the status to the endpoint list params
func (o *EndpointListParams) WithStatus(status []int64) *EndpointListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the endpoint list params
func (o *EndpointListParams) SetStatus(status []int64) {
	o.Status = status
}

// WithTagIds adds the tagIds to the endpoint list params
func (o *EndpointListParams) WithTagIds(tagIds []int64) *EndpointListParams {
	o.SetTagIds(tagIds)
	return o
}

// SetTagIds adds the tagIds to the endpoint list params
func (o *EndpointListParams) SetTagIds(tagIds []int64) {
	o.TagIds = tagIds
}

// WithTagsPartialMatch adds the tagsPartialMatch to the endpoint list params
func (o *EndpointListParams) WithTagsPartialMatch(tagsPartialMatch *bool) *EndpointListParams {
	o.SetTagsPartialMatch(tagsPartialMatch)
	return o
}

// SetTagsPartialMatch adds the tagsPartialMatch to the endpoint list params
func (o *EndpointListParams) SetTagsPartialMatch(tagsPartialMatch *bool) {
	o.TagsPartialMatch = tagsPartialMatch
}

// WithTypes adds the types to the endpoint list params
func (o *EndpointListParams) WithTypes(types []int64) *EndpointListParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the endpoint list params
func (o *EndpointListParams) SetTypes(types []int64) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *EndpointListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AgentVersions != nil {

		// binding items for agentVersions
		joinedAgentVersions := o.bindParamAgentVersions(reg)

		// query array param agentVersions
		if err := r.SetQueryParam("agentVersions", joinedAgentVersions...); err != nil {
			return err
		}
	}

	if o.EdgeDevice != nil {

		// query param edgeDevice
		var qrEdgeDevice bool

		if o.EdgeDevice != nil {
			qrEdgeDevice = *o.EdgeDevice
		}
		qEdgeDevice := swag.FormatBool(qrEdgeDevice)
		if qEdgeDevice != "" {

			if err := r.SetQueryParam("edgeDevice", qEdgeDevice); err != nil {
				return err
			}
		}
	}

	if o.EdgeDeviceUntrusted != nil {

		// query param edgeDeviceUntrusted
		var qrEdgeDeviceUntrusted bool

		if o.EdgeDeviceUntrusted != nil {
			qrEdgeDeviceUntrusted = *o.EdgeDeviceUntrusted
		}
		qEdgeDeviceUntrusted := swag.FormatBool(qrEdgeDeviceUntrusted)
		if qEdgeDeviceUntrusted != "" {

			if err := r.SetQueryParam("edgeDeviceUntrusted", qEdgeDeviceUntrusted); err != nil {
				return err
			}
		}
	}

	if o.EndpointIds != nil {

		// binding items for endpointIds
		joinedEndpointIds := o.bindParamEndpointIds(reg)

		// query array param endpointIds
		if err := r.SetQueryParam("endpointIds", joinedEndpointIds...); err != nil {
			return err
		}
	}

	if o.GroupIds != nil {

		// binding items for groupIds
		joinedGroupIds := o.bindParamGroupIds(reg)

		// query array param groupIds
		if err := r.SetQueryParam("groupIds", joinedGroupIds...); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Order != nil {

		// query param order
		var qrOrder int64

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := swag.FormatInt64(qrOrder)
		if qOrder != "" {

			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}
	}

	if o.Provisioned != nil {

		// query param provisioned
		var qrProvisioned bool

		if o.Provisioned != nil {
			qrProvisioned = *o.Provisioned
		}
		qProvisioned := swag.FormatBool(qrProvisioned)
		if qProvisioned != "" {

			if err := r.SetQueryParam("provisioned", qProvisioned); err != nil {
				return err
			}
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort int64

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := swag.FormatInt64(qrSort)
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if o.TagIds != nil {

		// binding items for tagIds
		joinedTagIds := o.bindParamTagIds(reg)

		// query array param tagIds
		if err := r.SetQueryParam("tagIds", joinedTagIds...); err != nil {
			return err
		}
	}

	if o.TagsPartialMatch != nil {

		// query param tagsPartialMatch
		var qrTagsPartialMatch bool

		if o.TagsPartialMatch != nil {
			qrTagsPartialMatch = *o.TagsPartialMatch
		}
		qTagsPartialMatch := swag.FormatBool(qrTagsPartialMatch)
		if qTagsPartialMatch != "" {

			if err := r.SetQueryParam("tagsPartialMatch", qTagsPartialMatch); err != nil {
				return err
			}
		}
	}

	if o.Types != nil {

		// binding items for types
		joinedTypes := o.bindParamTypes(reg)

		// query array param types
		if err := r.SetQueryParam("types", joinedTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEndpointList binds the parameter agentVersions
func (o *EndpointListParams) bindParamAgentVersions(formats strfmt.Registry) []string {
	agentVersionsIR := o.AgentVersions

	var agentVersionsIC []string
	for _, agentVersionsIIR := range agentVersionsIR { // explode []string

		agentVersionsIIV := agentVersionsIIR // string as string
		agentVersionsIC = append(agentVersionsIC, agentVersionsIIV)
	}

	// items.CollectionFormat: ""
	agentVersionsIS := swag.JoinByFormat(agentVersionsIC, "")

	return agentVersionsIS
}

// bindParamEndpointList binds the parameter endpointIds
func (o *EndpointListParams) bindParamEndpointIds(formats strfmt.Registry) []string {
	endpointIdsIR := o.EndpointIds

	var endpointIdsIC []string
	for _, endpointIdsIIR := range endpointIdsIR { // explode []int64

		endpointIdsIIV := swag.FormatInt64(endpointIdsIIR) // int64 as string
		endpointIdsIC = append(endpointIdsIC, endpointIdsIIV)
	}

	// items.CollectionFormat: ""
	endpointIdsIS := swag.JoinByFormat(endpointIdsIC, "")

	return endpointIdsIS
}

// bindParamEndpointList binds the parameter groupIds
func (o *EndpointListParams) bindParamGroupIds(formats strfmt.Registry) []string {
	groupIdsIR := o.GroupIds

	var groupIdsIC []string
	for _, groupIdsIIR := range groupIdsIR { // explode []int64

		groupIdsIIV := swag.FormatInt64(groupIdsIIR) // int64 as string
		groupIdsIC = append(groupIdsIC, groupIdsIIV)
	}

	// items.CollectionFormat: ""
	groupIdsIS := swag.JoinByFormat(groupIdsIC, "")

	return groupIdsIS
}

// bindParamEndpointList binds the parameter status
func (o *EndpointListParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []int64

		statusIIV := swag.FormatInt64(statusIIR) // int64 as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: ""
	statusIS := swag.JoinByFormat(statusIC, "")

	return statusIS
}

// bindParamEndpointList binds the parameter tagIds
func (o *EndpointListParams) bindParamTagIds(formats strfmt.Registry) []string {
	tagIdsIR := o.TagIds

	var tagIdsIC []string
	for _, tagIdsIIR := range tagIdsIR { // explode []int64

		tagIdsIIV := swag.FormatInt64(tagIdsIIR) // int64 as string
		tagIdsIC = append(tagIdsIC, tagIdsIIV)
	}

	// items.CollectionFormat: ""
	tagIdsIS := swag.JoinByFormat(tagIdsIC, "")

	return tagIdsIS
}

// bindParamEndpointList binds the parameter types
func (o *EndpointListParams) bindParamTypes(formats strfmt.Registry) []string {
	typesIR := o.Types

	var typesIC []string
	for _, typesIIR := range typesIR { // explode []int64

		typesIIV := swag.FormatInt64(typesIIR) // int64 as string
		typesIC = append(typesIC, typesIIV)
	}

	// items.CollectionFormat: ""
	typesIS := swag.JoinByFormat(typesIC, "")

	return typesIS
}
