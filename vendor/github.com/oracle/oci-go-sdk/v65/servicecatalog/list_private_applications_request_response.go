// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPrivateApplicationsRequest wrapper for the ListPrivateApplications operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListPrivateApplications.go.html to see an example of how to use ListPrivateApplicationsRequest.
type ListPrivateApplicationsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique identifier for the private application.
	PrivateApplicationId *string `mandatory:"false" contributesTo:"query" name:"privateApplicationId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// Default is `TIMECREATED`.
	SortBy ListPrivateApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to apply, either `ASC` or `DESC`. Default is `ASC`.
	SortOrder ListPrivateApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Exact match name filter.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivateApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivateApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivateApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivateApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivateApplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPrivateApplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivateApplicationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateApplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivateApplicationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivateApplicationsResponse wrapper for the ListPrivateApplications operation
type ListPrivateApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivateApplicationCollection instances
	PrivateApplicationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPrivateApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivateApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivateApplicationsSortByEnum Enum with underlying type: string
type ListPrivateApplicationsSortByEnum string

// Set of constants representing the allowable values for ListPrivateApplicationsSortByEnum
const (
	ListPrivateApplicationsSortByTimecreated    ListPrivateApplicationsSortByEnum = "TIMECREATED"
	ListPrivateApplicationsSortByLifecyclestate ListPrivateApplicationsSortByEnum = "LIFECYCLESTATE"
)

var mappingListPrivateApplicationsSortByEnum = map[string]ListPrivateApplicationsSortByEnum{
	"TIMECREATED":    ListPrivateApplicationsSortByTimecreated,
	"LIFECYCLESTATE": ListPrivateApplicationsSortByLifecyclestate,
}

var mappingListPrivateApplicationsSortByEnumLowerCase = map[string]ListPrivateApplicationsSortByEnum{
	"timecreated":    ListPrivateApplicationsSortByTimecreated,
	"lifecyclestate": ListPrivateApplicationsSortByLifecyclestate,
}

// GetListPrivateApplicationsSortByEnumValues Enumerates the set of values for ListPrivateApplicationsSortByEnum
func GetListPrivateApplicationsSortByEnumValues() []ListPrivateApplicationsSortByEnum {
	values := make([]ListPrivateApplicationsSortByEnum, 0)
	for _, v := range mappingListPrivateApplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateApplicationsSortByEnumStringValues Enumerates the set of values in String for ListPrivateApplicationsSortByEnum
func GetListPrivateApplicationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"LIFECYCLESTATE",
	}
}

// GetMappingListPrivateApplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateApplicationsSortByEnum(val string) (ListPrivateApplicationsSortByEnum, bool) {
	enum, ok := mappingListPrivateApplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivateApplicationsSortOrderEnum Enum with underlying type: string
type ListPrivateApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListPrivateApplicationsSortOrderEnum
const (
	ListPrivateApplicationsSortOrderAsc  ListPrivateApplicationsSortOrderEnum = "ASC"
	ListPrivateApplicationsSortOrderDesc ListPrivateApplicationsSortOrderEnum = "DESC"
)

var mappingListPrivateApplicationsSortOrderEnum = map[string]ListPrivateApplicationsSortOrderEnum{
	"ASC":  ListPrivateApplicationsSortOrderAsc,
	"DESC": ListPrivateApplicationsSortOrderDesc,
}

var mappingListPrivateApplicationsSortOrderEnumLowerCase = map[string]ListPrivateApplicationsSortOrderEnum{
	"asc":  ListPrivateApplicationsSortOrderAsc,
	"desc": ListPrivateApplicationsSortOrderDesc,
}

// GetListPrivateApplicationsSortOrderEnumValues Enumerates the set of values for ListPrivateApplicationsSortOrderEnum
func GetListPrivateApplicationsSortOrderEnumValues() []ListPrivateApplicationsSortOrderEnum {
	values := make([]ListPrivateApplicationsSortOrderEnum, 0)
	for _, v := range mappingListPrivateApplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateApplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListPrivateApplicationsSortOrderEnum
func GetListPrivateApplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivateApplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateApplicationsSortOrderEnum(val string) (ListPrivateApplicationsSortOrderEnum, bool) {
	enum, ok := mappingListPrivateApplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
