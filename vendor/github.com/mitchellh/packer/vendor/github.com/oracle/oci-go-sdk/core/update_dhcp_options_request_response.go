// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateDhcpOptionsRequest wrapper for the UpdateDhcpOptions operation
type UpdateDhcpOptionsRequest struct {

	// The OCID for the set of DHCP options.
	DhcpId *string `mandatory:"true" contributesTo:"path" name:"dhcpId"`

	// Request object for updating a set of DHCP options.
	UpdateDhcpDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateDhcpOptionsRequest) String() string {
	return common.PointerString(request)
}

// UpdateDhcpOptionsResponse wrapper for the UpdateDhcpOptions operation
type UpdateDhcpOptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DhcpOptions instance
	DhcpOptions `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateDhcpOptionsResponse) String() string {
	return common.PointerString(response)
}
