// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCrossConnectsRequest wrapper for the ListCrossConnects operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListCrossConnects.go.html to see an example of how to use ListCrossConnectsRequest.
type ListCrossConnectsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect group.
	CrossConnectGroupId *string `mandatory:"false" contributesTo:"query" name:"crossConnectGroupId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListCrossConnectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListCrossConnectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle
	// state. The value is case insensitive.
	LifecycleState CrossConnectLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCrossConnectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCrossConnectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCrossConnectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCrossConnectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCrossConnectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCrossConnectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCrossConnectsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCrossConnectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCrossConnectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCrossConnectLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCrossConnectLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCrossConnectsResponse wrapper for the ListCrossConnects operation
type ListCrossConnectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CrossConnect instances
	Items []CrossConnect `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCrossConnectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCrossConnectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCrossConnectsSortByEnum Enum with underlying type: string
type ListCrossConnectsSortByEnum string

// Set of constants representing the allowable values for ListCrossConnectsSortByEnum
const (
	ListCrossConnectsSortByTimecreated ListCrossConnectsSortByEnum = "TIMECREATED"
	ListCrossConnectsSortByDisplayname ListCrossConnectsSortByEnum = "DISPLAYNAME"
)

var mappingListCrossConnectsSortByEnum = map[string]ListCrossConnectsSortByEnum{
	"TIMECREATED": ListCrossConnectsSortByTimecreated,
	"DISPLAYNAME": ListCrossConnectsSortByDisplayname,
}

var mappingListCrossConnectsSortByEnumLowerCase = map[string]ListCrossConnectsSortByEnum{
	"timecreated": ListCrossConnectsSortByTimecreated,
	"displayname": ListCrossConnectsSortByDisplayname,
}

// GetListCrossConnectsSortByEnumValues Enumerates the set of values for ListCrossConnectsSortByEnum
func GetListCrossConnectsSortByEnumValues() []ListCrossConnectsSortByEnum {
	values := make([]ListCrossConnectsSortByEnum, 0)
	for _, v := range mappingListCrossConnectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCrossConnectsSortByEnumStringValues Enumerates the set of values in String for ListCrossConnectsSortByEnum
func GetListCrossConnectsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCrossConnectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCrossConnectsSortByEnum(val string) (ListCrossConnectsSortByEnum, bool) {
	enum, ok := mappingListCrossConnectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCrossConnectsSortOrderEnum Enum with underlying type: string
type ListCrossConnectsSortOrderEnum string

// Set of constants representing the allowable values for ListCrossConnectsSortOrderEnum
const (
	ListCrossConnectsSortOrderAsc  ListCrossConnectsSortOrderEnum = "ASC"
	ListCrossConnectsSortOrderDesc ListCrossConnectsSortOrderEnum = "DESC"
)

var mappingListCrossConnectsSortOrderEnum = map[string]ListCrossConnectsSortOrderEnum{
	"ASC":  ListCrossConnectsSortOrderAsc,
	"DESC": ListCrossConnectsSortOrderDesc,
}

var mappingListCrossConnectsSortOrderEnumLowerCase = map[string]ListCrossConnectsSortOrderEnum{
	"asc":  ListCrossConnectsSortOrderAsc,
	"desc": ListCrossConnectsSortOrderDesc,
}

// GetListCrossConnectsSortOrderEnumValues Enumerates the set of values for ListCrossConnectsSortOrderEnum
func GetListCrossConnectsSortOrderEnumValues() []ListCrossConnectsSortOrderEnum {
	values := make([]ListCrossConnectsSortOrderEnum, 0)
	for _, v := range mappingListCrossConnectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCrossConnectsSortOrderEnumStringValues Enumerates the set of values in String for ListCrossConnectsSortOrderEnum
func GetListCrossConnectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCrossConnectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCrossConnectsSortOrderEnum(val string) (ListCrossConnectsSortOrderEnum, bool) {
	enum, ok := mappingListCrossConnectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
