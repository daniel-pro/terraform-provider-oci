// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateLoadBalancerDetails Configuration details to update a load balancer.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateLoadBalancerDetails struct {

	// The user-friendly display name for the load balancer. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether or not the load balancer has delete protection enabled.
	// If "true", the loadbalancer will be protected against deletion if configured to accept traffic.
	// If "false", the loadbalancer will not be protected against deletion.
	// If null or unset, the value for delete protection will not be changed.
	// Example: `true`
	IsDeleteProtectionEnabled *bool `mandatory:"false" json:"isDeleteProtectionEnabled"`

	// Request ID is an identifier given to every request that goes through the load balancer.
	// The same request id will be generated for both incoming request and the corresponding outgoing response.
	// The header X-Request-ID (default name) holding the value of the request ID will be added to both request and response.
	// If the header already exists i.e. it was sent by the caller or returned by the backend then its value will not be changed.
	// Request ID header property allows:
	// 1. specifying name of the header holding the request ID;
	// 2. switching this feature off by setting the header name to empty string.
	// **Notes:**
	// * The header name must conform to the
	//   RFC 7230 - Hypertext Transfer Protocol (HTTP/1.1) (https://datatracker.ietf.org/doc/html/rfc7230) standard.
	// * The header name must start with "X-" prefix.
	RequestIdHeader *string `mandatory:"false" json:"requestIdHeader"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
