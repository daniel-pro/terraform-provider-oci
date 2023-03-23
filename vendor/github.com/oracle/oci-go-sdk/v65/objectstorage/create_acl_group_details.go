// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAclGroupDetails The details to create an ACL Group.
type CreateAclGroupDetails struct {

	// User-friendly name to describe the ACL Group. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the ACL Group. This can be useful for identifying the purpose of the ACL Group.
	Description *string `mandatory:"false" json:"description"`

	// The compartment containing this ACL Group.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Determines whether this ACL Group can be associated with a bucket, compartment, or tenancy.
	AclType CreateAclGroupDetailsAclTypeEnum `mandatory:"false" json:"aclType,omitempty"`

	// Specifies requests that are not subject to any ACLs in the ACL group.
	AclGroupExceptions []CreateAclGroupDetailsAclGroupExceptionsEnum `mandatory:"false" json:"aclGroupExceptions,omitempty"`

	// An ordered list of ACL IDs. The ACLs are evaluated first to last.
	AclIds []string `mandatory:"false" json:"aclIds"`
}

func (m CreateAclGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAclGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateAclGroupDetailsAclTypeEnum(string(m.AclType)); !ok && m.AclType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclType: %s. Supported values are: %s.", m.AclType, strings.Join(GetCreateAclGroupDetailsAclTypeEnumStringValues(), ",")))
	}
	for _, val := range m.AclGroupExceptions {
		if _, ok := GetMappingCreateAclGroupDetailsAclGroupExceptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupExceptions: %s. Supported values are: %s.", val, strings.Join(GetCreateAclGroupDetailsAclGroupExceptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAclGroupDetailsAclTypeEnum Enum with underlying type: string
type CreateAclGroupDetailsAclTypeEnum string

// Set of constants representing the allowable values for CreateAclGroupDetailsAclTypeEnum
const (
	CreateAclGroupDetailsAclTypeTenancy     CreateAclGroupDetailsAclTypeEnum = "TENANCY"
	CreateAclGroupDetailsAclTypeCompartment CreateAclGroupDetailsAclTypeEnum = "COMPARTMENT"
	CreateAclGroupDetailsAclTypeBucket      CreateAclGroupDetailsAclTypeEnum = "BUCKET"
)

var mappingCreateAclGroupDetailsAclTypeEnum = map[string]CreateAclGroupDetailsAclTypeEnum{
	"TENANCY":     CreateAclGroupDetailsAclTypeTenancy,
	"COMPARTMENT": CreateAclGroupDetailsAclTypeCompartment,
	"BUCKET":      CreateAclGroupDetailsAclTypeBucket,
}

var mappingCreateAclGroupDetailsAclTypeEnumLowerCase = map[string]CreateAclGroupDetailsAclTypeEnum{
	"tenancy":     CreateAclGroupDetailsAclTypeTenancy,
	"compartment": CreateAclGroupDetailsAclTypeCompartment,
	"bucket":      CreateAclGroupDetailsAclTypeBucket,
}

// GetCreateAclGroupDetailsAclTypeEnumValues Enumerates the set of values for CreateAclGroupDetailsAclTypeEnum
func GetCreateAclGroupDetailsAclTypeEnumValues() []CreateAclGroupDetailsAclTypeEnum {
	values := make([]CreateAclGroupDetailsAclTypeEnum, 0)
	for _, v := range mappingCreateAclGroupDetailsAclTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAclGroupDetailsAclTypeEnumStringValues Enumerates the set of values in String for CreateAclGroupDetailsAclTypeEnum
func GetCreateAclGroupDetailsAclTypeEnumStringValues() []string {
	return []string{
		"TENANCY",
		"COMPARTMENT",
		"BUCKET",
	}
}

// GetMappingCreateAclGroupDetailsAclTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAclGroupDetailsAclTypeEnum(val string) (CreateAclGroupDetailsAclTypeEnum, bool) {
	enum, ok := mappingCreateAclGroupDetailsAclTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAclGroupDetailsAclGroupExceptionsEnum Enum with underlying type: string
type CreateAclGroupDetailsAclGroupExceptionsEnum string

// Set of constants representing the allowable values for CreateAclGroupDetailsAclGroupExceptionsEnum
const (
	CreateAclGroupDetailsAclGroupExceptionsCopy        CreateAclGroupDetailsAclGroupExceptionsEnum = "COPY"
	CreateAclGroupDetailsAclGroupExceptionsReplication CreateAclGroupDetailsAclGroupExceptionsEnum = "REPLICATION"
)

var mappingCreateAclGroupDetailsAclGroupExceptionsEnum = map[string]CreateAclGroupDetailsAclGroupExceptionsEnum{
	"COPY":        CreateAclGroupDetailsAclGroupExceptionsCopy,
	"REPLICATION": CreateAclGroupDetailsAclGroupExceptionsReplication,
}

var mappingCreateAclGroupDetailsAclGroupExceptionsEnumLowerCase = map[string]CreateAclGroupDetailsAclGroupExceptionsEnum{
	"copy":        CreateAclGroupDetailsAclGroupExceptionsCopy,
	"replication": CreateAclGroupDetailsAclGroupExceptionsReplication,
}

// GetCreateAclGroupDetailsAclGroupExceptionsEnumValues Enumerates the set of values for CreateAclGroupDetailsAclGroupExceptionsEnum
func GetCreateAclGroupDetailsAclGroupExceptionsEnumValues() []CreateAclGroupDetailsAclGroupExceptionsEnum {
	values := make([]CreateAclGroupDetailsAclGroupExceptionsEnum, 0)
	for _, v := range mappingCreateAclGroupDetailsAclGroupExceptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAclGroupDetailsAclGroupExceptionsEnumStringValues Enumerates the set of values in String for CreateAclGroupDetailsAclGroupExceptionsEnum
func GetCreateAclGroupDetailsAclGroupExceptionsEnumStringValues() []string {
	return []string{
		"COPY",
		"REPLICATION",
	}
}

// GetMappingCreateAclGroupDetailsAclGroupExceptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAclGroupDetailsAclGroupExceptionsEnum(val string) (CreateAclGroupDetailsAclGroupExceptionsEnum, bool) {
	enum, ok := mappingCreateAclGroupDetailsAclGroupExceptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
