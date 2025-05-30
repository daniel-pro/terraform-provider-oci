// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OnPremiseVantagePoint The information about an On-premise vantage point.
type OnPremiseVantagePoint struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the On-premise vantage point.
	Id *string `mandatory:"true" json:"id"`

	// Unique permanent name of the On-premise vantage point.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique On-premise vantage point name that cannot be edited. The name should not contain any confidential information.
	Name *string `mandatory:"true" json:"name"`

	// Type of On-premise vantage point.
	Type OnPremiseVantagePointTypeEnum `mandatory:"true" json:"type"`

	WorkersSummary *WorkersSummary `mandatory:"true" json:"workersSummary"`

	// A short description about the On-premise vantage point.
	Description *string `mandatory:"false" json:"description"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OnPremiseVantagePoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OnPremiseVantagePoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOnPremiseVantagePointTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOnPremiseVantagePointTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OnPremiseVantagePointTypeEnum Enum with underlying type: string
type OnPremiseVantagePointTypeEnum string

// Set of constants representing the allowable values for OnPremiseVantagePointTypeEnum
const (
	OnPremiseVantagePointTypeOnPremiseDockerVantagePoint OnPremiseVantagePointTypeEnum = "ON_PREMISE_DOCKER_VANTAGE_POINT"
)

var mappingOnPremiseVantagePointTypeEnum = map[string]OnPremiseVantagePointTypeEnum{
	"ON_PREMISE_DOCKER_VANTAGE_POINT": OnPremiseVantagePointTypeOnPremiseDockerVantagePoint,
}

var mappingOnPremiseVantagePointTypeEnumLowerCase = map[string]OnPremiseVantagePointTypeEnum{
	"on_premise_docker_vantage_point": OnPremiseVantagePointTypeOnPremiseDockerVantagePoint,
}

// GetOnPremiseVantagePointTypeEnumValues Enumerates the set of values for OnPremiseVantagePointTypeEnum
func GetOnPremiseVantagePointTypeEnumValues() []OnPremiseVantagePointTypeEnum {
	values := make([]OnPremiseVantagePointTypeEnum, 0)
	for _, v := range mappingOnPremiseVantagePointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOnPremiseVantagePointTypeEnumStringValues Enumerates the set of values in String for OnPremiseVantagePointTypeEnum
func GetOnPremiseVantagePointTypeEnumStringValues() []string {
	return []string{
		"ON_PREMISE_DOCKER_VANTAGE_POINT",
	}
}

// GetMappingOnPremiseVantagePointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOnPremiseVantagePointTypeEnum(val string) (OnPremiseVantagePointTypeEnum, bool) {
	enum, ok := mappingOnPremiseVantagePointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
