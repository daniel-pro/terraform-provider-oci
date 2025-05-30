// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DeleteDeploymentBackupRequest wrapper for the DeleteDeploymentBackup operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteDeploymentBackup.go.html to see an example of how to use DeleteDeploymentBackupRequest.
type DeleteDeploymentBackupRequest struct {

	// A unique DeploymentBackup identifier.
	DeploymentBackupId *string `mandatory:"true" contributesTo:"path" name:"deploymentBackupId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the
	// `if-match` parameter to the value of the etag from a previous GET or POST response for that
	// resource.  The resource is updated or deleted only if the etag you provide matches the
	// resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Whether to override locks (if any exist).
	IsLockOverride *bool `mandatory:"false" contributesTo:"query" name:"isLockOverride"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteDeploymentBackupRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteDeploymentBackupRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteDeploymentBackupRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteDeploymentBackupRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteDeploymentBackupRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteDeploymentBackupResponse wrapper for the DeleteDeploymentBackup operation
type DeleteDeploymentBackupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A unique Oracle-assigned identifier for an asynchronous request. You can use this to query
	// status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteDeploymentBackupResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteDeploymentBackupResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
