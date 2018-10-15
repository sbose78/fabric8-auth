// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "tenant": Application User Types
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-tenant/design
// --notool=true
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-auth/account
// --pkg=tenant
// --version=v1.3.0

package tenant

import (
	"github.com/goadesign/goa"
	uuid "github.com/goadesign/goa/uuid"
	"time"
)

// genericData user type.
type genericData struct {
	// UUID of the object
	ID    *string       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *genericLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Type  *string       `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Publicize creates GenericData from genericData
func (ut *genericData) Publicize() *GenericData {
	var pub GenericData
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Type != nil {
		pub.Type = ut.Type
	}
	return &pub
}

// GenericData user type.
type GenericData struct {
	// UUID of the object
	ID    *string       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *GenericLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Type  *string       `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// genericLinks user type.
type genericLinks struct {
	Meta    map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Related *string                `form:"related,omitempty" json:"related,omitempty" xml:"related,omitempty"`
	Self    *string                `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
}

// Publicize creates GenericLinks from genericLinks
func (ut *genericLinks) Publicize() *GenericLinks {
	var pub GenericLinks
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	if ut.Related != nil {
		pub.Related = ut.Related
	}
	if ut.Self != nil {
		pub.Self = ut.Self
	}
	return &pub
}

// GenericLinks user type.
type GenericLinks struct {
	Meta    map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Related *string                `form:"related,omitempty" json:"related,omitempty" xml:"related,omitempty"`
	Self    *string                `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
}

// Error objects provide additional information about problems encountered while
// performing an operation. Error objects MUST be returned as an array keyed by errors in the
// top level of a JSON API document.
//
// See. also http://jsonapi.org/format/#error-objects.
type jSONAPIError struct {
	// an application-specific error code, expressed as a string value.
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// a human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
	// a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// a links object containing the following members:
	// * about: a link that leads to further details about this particular occurrence of the problem.
	Links map[string]*jSONAPILink `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// a meta object containing non-standard meta-information about the error
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// an object containing references to the source of the error,
	// optionally including any of the following members
	//
	// * pointer: a JSON Pointer [RFC6901] to the associated entity in the request document [e.g. "/data" for a primary data object,
	//            or "/data/attributes/title" for a specific attribute].
	// * parameter: a string indicating which URI query parameter caused the error.
	Source map[string]interface{} `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// the HTTP status code applicable to this problem, expressed as a string value.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// a short, human-readable summary of the problem that SHOULD NOT
	// change from occurrence to occurrence of the problem, except for purposes of localization.
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the jSONAPIError type instance.
func (ut *jSONAPIError) Validate() (err error) {
	if ut.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "detail"))
	}
	return
}

// Publicize creates JSONAPIError from jSONAPIError
func (ut *jSONAPIError) Publicize() *JSONAPIError {
	var pub JSONAPIError
	if ut.Code != nil {
		pub.Code = ut.Code
	}
	if ut.Detail != nil {
		pub.Detail = *ut.Detail
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Links != nil {
		pub.Links = make(map[string]*JSONAPILink, len(ut.Links))
		for k2, v2 := range ut.Links {
			pubk2 := k2
			var pubv2 *JSONAPILink
			if v2 != nil {
				pubv2 = v2.Publicize()
			}
			pub.Links[pubk2] = pubv2
		}
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	if ut.Source != nil {
		pub.Source = ut.Source
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// Error objects provide additional information about problems encountered while
// performing an operation. Error objects MUST be returned as an array keyed by errors in the
// top level of a JSON API document.
//
// See. also http://jsonapi.org/format/#error-objects.
type JSONAPIError struct {
	// an application-specific error code, expressed as a string value.
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// a human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail string `form:"detail" json:"detail" xml:"detail"`
	// a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// a links object containing the following members:
	// * about: a link that leads to further details about this particular occurrence of the problem.
	Links map[string]*JSONAPILink `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// a meta object containing non-standard meta-information about the error
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// an object containing references to the source of the error,
	// optionally including any of the following members
	//
	// * pointer: a JSON Pointer [RFC6901] to the associated entity in the request document [e.g. "/data" for a primary data object,
	//            or "/data/attributes/title" for a specific attribute].
	// * parameter: a string indicating which URI query parameter caused the error.
	Source map[string]interface{} `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// the HTTP status code applicable to this problem, expressed as a string value.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// a short, human-readable summary of the problem that SHOULD NOT
	// change from occurrence to occurrence of the problem, except for purposes of localization.
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the JSONAPIError type instance.
func (ut *JSONAPIError) Validate() (err error) {
	if ut.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "detail"))
	}
	return
}

// See also http://jsonapi.org/format/#document-links.
type jSONAPILink struct {
	// a string containing the link's URL.
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// a meta object containing non-standard meta-information about the link.
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Publicize creates JSONAPILink from jSONAPILink
func (ut *jSONAPILink) Publicize() *JSONAPILink {
	var pub JSONAPILink
	if ut.Href != nil {
		pub.Href = ut.Href
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	return &pub
}

// See also http://jsonapi.org/format/#document-links.
type JSONAPILink struct {
	// a string containing the link's URL.
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// a meta object containing non-standard meta-information about the link.
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// JSONAPI store for all the "attributes" of a Tenant namespace. See also see http://jsonapi.org/format/#document-resource-object-attributes
type namespaceAttributes struct {
	// The cluster app domain
	ClusterAppDomain *string `form:"cluster-app-domain,omitempty" json:"cluster-app-domain,omitempty" xml:"cluster-app-domain,omitempty"`
	// Whether cluster hosting this namespace exhausted it's capacity
	ClusterCapacityExhausted *bool `form:"cluster-capacity-exhausted,omitempty" json:"cluster-capacity-exhausted,omitempty" xml:"cluster-capacity-exhausted,omitempty"`
	// The cluster console url
	ClusterConsoleURL *string `form:"cluster-console-url,omitempty" json:"cluster-console-url,omitempty" xml:"cluster-console-url,omitempty"`
	// The cluster logging url
	ClusterLoggingURL *string `form:"cluster-logging-url,omitempty" json:"cluster-logging-url,omitempty" xml:"cluster-logging-url,omitempty"`
	// The cluster metrics url
	ClusterMetricsURL *string `form:"cluster-metrics-url,omitempty" json:"cluster-metrics-url,omitempty" xml:"cluster-metrics-url,omitempty"`
	// The cluster url
	ClusterURL *string `form:"cluster-url,omitempty" json:"cluster-url,omitempty" xml:"cluster-url,omitempty"`
	// When the tenant was created
	CreatedAt *time.Time `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	// The namespace name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The namespaces state
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// The tenant namespaces
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// When the tenant was updated
	UpdatedAt *time.Time `form:"updated-at,omitempty" json:"updated-at,omitempty" xml:"updated-at,omitempty"`
	// The namespaces version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Validate validates the namespaceAttributes type instance.
func (ut *namespaceAttributes) Validate() (err error) {
	if ut.Type != nil {
		if !(*ut.Type == "user" || *ut.Type == "che" || *ut.Type == "jenkins" || *ut.Type == "stage" || *ut.Type == "test" || *ut.Type == "run") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.type`, *ut.Type, []interface{}{"user", "che", "jenkins", "stage", "test", "run"}))
		}
	}
	return
}

// Publicize creates NamespaceAttributes from namespaceAttributes
func (ut *namespaceAttributes) Publicize() *NamespaceAttributes {
	var pub NamespaceAttributes
	if ut.ClusterAppDomain != nil {
		pub.ClusterAppDomain = ut.ClusterAppDomain
	}
	if ut.ClusterCapacityExhausted != nil {
		pub.ClusterCapacityExhausted = ut.ClusterCapacityExhausted
	}
	if ut.ClusterConsoleURL != nil {
		pub.ClusterConsoleURL = ut.ClusterConsoleURL
	}
	if ut.ClusterLoggingURL != nil {
		pub.ClusterLoggingURL = ut.ClusterLoggingURL
	}
	if ut.ClusterMetricsURL != nil {
		pub.ClusterMetricsURL = ut.ClusterMetricsURL
	}
	if ut.ClusterURL != nil {
		pub.ClusterURL = ut.ClusterURL
	}
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.State != nil {
		pub.State = ut.State
	}
	if ut.Type != nil {
		pub.Type = ut.Type
	}
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	if ut.Version != nil {
		pub.Version = ut.Version
	}
	return &pub
}

// JSONAPI store for all the "attributes" of a Tenant namespace. See also see http://jsonapi.org/format/#document-resource-object-attributes
type NamespaceAttributes struct {
	// The cluster app domain
	ClusterAppDomain *string `form:"cluster-app-domain,omitempty" json:"cluster-app-domain,omitempty" xml:"cluster-app-domain,omitempty"`
	// Whether cluster hosting this namespace exhausted it's capacity
	ClusterCapacityExhausted *bool `form:"cluster-capacity-exhausted,omitempty" json:"cluster-capacity-exhausted,omitempty" xml:"cluster-capacity-exhausted,omitempty"`
	// The cluster console url
	ClusterConsoleURL *string `form:"cluster-console-url,omitempty" json:"cluster-console-url,omitempty" xml:"cluster-console-url,omitempty"`
	// The cluster logging url
	ClusterLoggingURL *string `form:"cluster-logging-url,omitempty" json:"cluster-logging-url,omitempty" xml:"cluster-logging-url,omitempty"`
	// The cluster metrics url
	ClusterMetricsURL *string `form:"cluster-metrics-url,omitempty" json:"cluster-metrics-url,omitempty" xml:"cluster-metrics-url,omitempty"`
	// The cluster url
	ClusterURL *string `form:"cluster-url,omitempty" json:"cluster-url,omitempty" xml:"cluster-url,omitempty"`
	// When the tenant was created
	CreatedAt *time.Time `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	// The namespace name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The namespaces state
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
	// The tenant namespaces
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// When the tenant was updated
	UpdatedAt *time.Time `form:"updated-at,omitempty" json:"updated-at,omitempty" xml:"updated-at,omitempty"`
	// The namespaces version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Validate validates the NamespaceAttributes type instance.
func (ut *NamespaceAttributes) Validate() (err error) {
	if ut.Type != nil {
		if !(*ut.Type == "user" || *ut.Type == "che" || *ut.Type == "jenkins" || *ut.Type == "stage" || *ut.Type == "test" || *ut.Type == "run") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.type`, *ut.Type, []interface{}{"user", "che", "jenkins", "stage", "test", "run"}))
		}
	}
	return
}

// relationGeneric user type.
type relationGeneric struct {
	Data  *genericData           `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *genericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Publicize creates RelationGeneric from relationGeneric
func (ut *relationGeneric) Publicize() *RelationGeneric {
	var pub RelationGeneric
	if ut.Data != nil {
		pub.Data = ut.Data.Publicize()
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	return &pub
}

// RelationGeneric user type.
type RelationGeneric struct {
	Data  *GenericData           `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *GenericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// relationGenericList user type.
type relationGenericList struct {
	Data  []*genericData         `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *genericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Publicize creates RelationGenericList from relationGenericList
func (ut *relationGenericList) Publicize() *RelationGenericList {
	var pub RelationGenericList
	if ut.Data != nil {
		pub.Data = make([]*GenericData, len(ut.Data))
		for i2, elem2 := range ut.Data {
			pub.Data[i2] = elem2.Publicize()
		}
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	return &pub
}

// RelationGenericList user type.
type RelationGenericList struct {
	Data  []*GenericData         `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *GenericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// JSONAPI for the tenant object. See also http://jsonapi.org/format/#document-resource-object
type tenant_ struct {
	Attributes *tenantAttributes `form:"attributes,omitempty" json:"attributes,omitempty" xml:"attributes,omitempty"`
	// ID of tenant
	ID    *uuid.UUID    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *genericLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Type  *string       `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the tenant_ type instance.
func (ut *tenant_) Validate() (err error) {
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "type"))
	}
	if ut.Attributes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "attributes"))
	}
	if ut.Attributes != nil {
		if err2 := ut.Attributes.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Type != nil {
		if !(*ut.Type == "tenants") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.type`, *ut.Type, []interface{}{"tenants"}))
		}
	}
	return
}

// Publicize creates Tenant from tenant_
func (ut *tenant_) Publicize() *Tenant {
	var pub Tenant
	if ut.Attributes != nil {
		pub.Attributes = ut.Attributes.Publicize()
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	return &pub
}

// JSONAPI for the tenant object. See also http://jsonapi.org/format/#document-resource-object
type Tenant struct {
	Attributes *TenantAttributes `form:"attributes" json:"attributes" xml:"attributes"`
	// ID of tenant
	ID    *uuid.UUID    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *GenericLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Type  string        `form:"type" json:"type" xml:"type"`
}

// Validate validates the Tenant type instance.
func (ut *Tenant) Validate() (err error) {
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "type"))
	}
	if ut.Attributes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "attributes"))
	}
	if ut.Attributes != nil {
		if err2 := ut.Attributes.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if !(ut.Type == "tenants") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.type`, ut.Type, []interface{}{"tenants"}))
	}
	return
}

// JSONAPI store for all the "attributes" of a Tenant. See also see http://jsonapi.org/format/#document-resource-object-attributes
type tenantAttributes struct {
	// When the tenant was created
	CreatedAt *time.Time `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	// The tenant name
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// The tenant namespaces
	Namespaces []*namespaceAttributes `form:"namespaces,omitempty" json:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// User profile type
	Profile *string `form:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
}

// Validate validates the tenantAttributes type instance.
func (ut *tenantAttributes) Validate() (err error) {
	for _, e := range ut.Namespaces {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Publicize creates TenantAttributes from tenantAttributes
func (ut *tenantAttributes) Publicize() *TenantAttributes {
	var pub TenantAttributes
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Namespaces != nil {
		pub.Namespaces = make([]*NamespaceAttributes, len(ut.Namespaces))
		for i2, elem2 := range ut.Namespaces {
			pub.Namespaces[i2] = elem2.Publicize()
		}
	}
	if ut.Profile != nil {
		pub.Profile = ut.Profile
	}
	return &pub
}

// JSONAPI store for all the "attributes" of a Tenant. See also see http://jsonapi.org/format/#document-resource-object-attributes
type TenantAttributes struct {
	// When the tenant was created
	CreatedAt *time.Time `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	// The tenant name
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// The tenant namespaces
	Namespaces []*NamespaceAttributes `form:"namespaces,omitempty" json:"namespaces,omitempty" xml:"namespaces,omitempty"`
	// User profile type
	Profile *string `form:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
}

// Validate validates the TenantAttributes type instance.
func (ut *TenantAttributes) Validate() (err error) {
	for _, e := range ut.Namespaces {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// tenantListMeta user type.
type tenantListMeta struct {
	TotalCount *int `form:"totalCount,omitempty" json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

// Validate validates the tenantListMeta type instance.
func (ut *tenantListMeta) Validate() (err error) {
	if ut.TotalCount == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "totalCount"))
	}
	return
}

// Publicize creates TenantListMeta from tenantListMeta
func (ut *tenantListMeta) Publicize() *TenantListMeta {
	var pub TenantListMeta
	if ut.TotalCount != nil {
		pub.TotalCount = *ut.TotalCount
	}
	return &pub
}

// TenantListMeta user type.
type TenantListMeta struct {
	TotalCount int `form:"totalCount" json:"totalCount" xml:"totalCount"`
}

// pagingLinks user type.
type pagingLinks struct {
	Filters *string `form:"filters,omitempty" json:"filters,omitempty" xml:"filters,omitempty"`
	First   *string `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Last    *string `form:"last,omitempty" json:"last,omitempty" xml:"last,omitempty"`
	Next    *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
	Prev    *string `form:"prev,omitempty" json:"prev,omitempty" xml:"prev,omitempty"`
}

// Publicize creates PagingLinks from pagingLinks
func (ut *pagingLinks) Publicize() *PagingLinks {
	var pub PagingLinks
	if ut.Filters != nil {
		pub.Filters = ut.Filters
	}
	if ut.First != nil {
		pub.First = ut.First
	}
	if ut.Last != nil {
		pub.Last = ut.Last
	}
	if ut.Next != nil {
		pub.Next = ut.Next
	}
	if ut.Prev != nil {
		pub.Prev = ut.Prev
	}
	return &pub
}

// PagingLinks user type.
type PagingLinks struct {
	Filters *string `form:"filters,omitempty" json:"filters,omitempty" xml:"filters,omitempty"`
	First   *string `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Last    *string `form:"last,omitempty" json:"last,omitempty" xml:"last,omitempty"`
	Next    *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
	Prev    *string `form:"prev,omitempty" json:"prev,omitempty" xml:"prev,omitempty"`
}