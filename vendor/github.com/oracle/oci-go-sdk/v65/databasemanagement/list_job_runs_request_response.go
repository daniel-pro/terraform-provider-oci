// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJobRunsRequest wrapper for the ListJobRuns operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListJobRuns.go.html to see an example of how to use ListJobRunsRequest.
type ListJobRunsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The identifier of the resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The identifier of the job.
	JobId *string `mandatory:"false" contributesTo:"query" name:"jobId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// The status of the job run.
	RunStatus *string `mandatory:"false" contributesTo:"query" name:"runStatus"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListJobRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListJobRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobRunsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobRunsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJobRunsResponse wrapper for the ListJobRuns operation
type ListJobRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobRunCollection instances
	JobRunCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobRunsSortByEnum Enum with underlying type: string
type ListJobRunsSortByEnum string

// Set of constants representing the allowable values for ListJobRunsSortByEnum
const (
	ListJobRunsSortByTimecreated ListJobRunsSortByEnum = "TIMECREATED"
	ListJobRunsSortByName        ListJobRunsSortByEnum = "NAME"
)

var mappingListJobRunsSortByEnum = map[string]ListJobRunsSortByEnum{
	"TIMECREATED": ListJobRunsSortByTimecreated,
	"NAME":        ListJobRunsSortByName,
}

var mappingListJobRunsSortByEnumLowerCase = map[string]ListJobRunsSortByEnum{
	"timecreated": ListJobRunsSortByTimecreated,
	"name":        ListJobRunsSortByName,
}

// GetListJobRunsSortByEnumValues Enumerates the set of values for ListJobRunsSortByEnum
func GetListJobRunsSortByEnumValues() []ListJobRunsSortByEnum {
	values := make([]ListJobRunsSortByEnum, 0)
	for _, v := range mappingListJobRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobRunsSortByEnumStringValues Enumerates the set of values in String for ListJobRunsSortByEnum
func GetListJobRunsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListJobRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobRunsSortByEnum(val string) (ListJobRunsSortByEnum, bool) {
	enum, ok := mappingListJobRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobRunsSortOrderEnum Enum with underlying type: string
type ListJobRunsSortOrderEnum string

// Set of constants representing the allowable values for ListJobRunsSortOrderEnum
const (
	ListJobRunsSortOrderAsc  ListJobRunsSortOrderEnum = "ASC"
	ListJobRunsSortOrderDesc ListJobRunsSortOrderEnum = "DESC"
)

var mappingListJobRunsSortOrderEnum = map[string]ListJobRunsSortOrderEnum{
	"ASC":  ListJobRunsSortOrderAsc,
	"DESC": ListJobRunsSortOrderDesc,
}

var mappingListJobRunsSortOrderEnumLowerCase = map[string]ListJobRunsSortOrderEnum{
	"asc":  ListJobRunsSortOrderAsc,
	"desc": ListJobRunsSortOrderDesc,
}

// GetListJobRunsSortOrderEnumValues Enumerates the set of values for ListJobRunsSortOrderEnum
func GetListJobRunsSortOrderEnumValues() []ListJobRunsSortOrderEnum {
	values := make([]ListJobRunsSortOrderEnum, 0)
	for _, v := range mappingListJobRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobRunsSortOrderEnumStringValues Enumerates the set of values in String for ListJobRunsSortOrderEnum
func GetListJobRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobRunsSortOrderEnum(val string) (ListJobRunsSortOrderEnum, bool) {
	enum, ok := mappingListJobRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
