// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetGenericArtifactByPathRequest wrapper for the GetGenericArtifactByPath operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetGenericArtifactByPath.go.html to see an example of how to use GetGenericArtifactByPathRequest.
type GetGenericArtifactByPathRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
	// Example: `ocid1.artifactrepository.oc1..exampleuniqueID`
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// A user-defined path to describe the location of an artifact. You can use slashes to organize the repository, but slashes do not create a directory structure. An artifact path does not include an artifact version.
	// Example: `project01/my-web-app/artifact-abc`
	ArtifactPath *string `mandatory:"true" contributesTo:"path" name:"artifactPath"`

	// A user-defined string to describe the artifact version.
	// Example: `1.1.2` or `1.2-beta-2`
	Version *string `mandatory:"true" contributesTo:"path" name:"version"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetGenericArtifactByPathRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetGenericArtifactByPathRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetGenericArtifactByPathRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetGenericArtifactByPathRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetGenericArtifactByPathRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetGenericArtifactByPathResponse wrapper for the GetGenericArtifactByPath operation
type GetGenericArtifactByPathResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GenericArtifact instance
	GenericArtifact `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetGenericArtifactByPathResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetGenericArtifactByPathResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
