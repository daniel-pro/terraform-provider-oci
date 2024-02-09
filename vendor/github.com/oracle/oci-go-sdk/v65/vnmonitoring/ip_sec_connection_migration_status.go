// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpSecConnectionMigrationStatus The IPSec connection's migration status.
type IpSecConnectionMigrationStatus struct {

	// The IPSec connection's migration status.
	MigrationStatus IpSecConnectionMigrationStatusMigrationStatusEnum `mandatory:"true" json:"migrationStatus"`

	// The start timestamp for Site-to-Site VPN migration work.
	StartTimeStamp *common.SDKTime `mandatory:"true" json:"startTimeStamp"`
}

func (m IpSecConnectionMigrationStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpSecConnectionMigrationStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIpSecConnectionMigrationStatusMigrationStatusEnum(string(m.MigrationStatus)); !ok && m.MigrationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MigrationStatus: %s. Supported values are: %s.", m.MigrationStatus, strings.Join(GetIpSecConnectionMigrationStatusMigrationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IpSecConnectionMigrationStatusMigrationStatusEnum Enum with underlying type: string
type IpSecConnectionMigrationStatusMigrationStatusEnum string

// Set of constants representing the allowable values for IpSecConnectionMigrationStatusMigrationStatusEnum
const (
	IpSecConnectionMigrationStatusMigrationStatusReady           IpSecConnectionMigrationStatusMigrationStatusEnum = "READY"
	IpSecConnectionMigrationStatusMigrationStatusMigrated        IpSecConnectionMigrationStatusMigrationStatusEnum = "MIGRATED"
	IpSecConnectionMigrationStatusMigrationStatusMigrating       IpSecConnectionMigrationStatusMigrationStatusEnum = "MIGRATING"
	IpSecConnectionMigrationStatusMigrationStatusMigrationFailed IpSecConnectionMigrationStatusMigrationStatusEnum = "MIGRATION_FAILED"
	IpSecConnectionMigrationStatusMigrationStatusRolledBack      IpSecConnectionMigrationStatusMigrationStatusEnum = "ROLLED_BACK"
	IpSecConnectionMigrationStatusMigrationStatusRollingBack     IpSecConnectionMigrationStatusMigrationStatusEnum = "ROLLING_BACK"
	IpSecConnectionMigrationStatusMigrationStatusRollbackFailed  IpSecConnectionMigrationStatusMigrationStatusEnum = "ROLLBACK_FAILED"
	IpSecConnectionMigrationStatusMigrationStatusNotApplicable   IpSecConnectionMigrationStatusMigrationStatusEnum = "NOT_APPLICABLE"
	IpSecConnectionMigrationStatusMigrationStatusManual          IpSecConnectionMigrationStatusMigrationStatusEnum = "MANUAL"
	IpSecConnectionMigrationStatusMigrationStatusValidating      IpSecConnectionMigrationStatusMigrationStatusEnum = "VALIDATING"
)

var mappingIpSecConnectionMigrationStatusMigrationStatusEnum = map[string]IpSecConnectionMigrationStatusMigrationStatusEnum{
	"READY":            IpSecConnectionMigrationStatusMigrationStatusReady,
	"MIGRATED":         IpSecConnectionMigrationStatusMigrationStatusMigrated,
	"MIGRATING":        IpSecConnectionMigrationStatusMigrationStatusMigrating,
	"MIGRATION_FAILED": IpSecConnectionMigrationStatusMigrationStatusMigrationFailed,
	"ROLLED_BACK":      IpSecConnectionMigrationStatusMigrationStatusRolledBack,
	"ROLLING_BACK":     IpSecConnectionMigrationStatusMigrationStatusRollingBack,
	"ROLLBACK_FAILED":  IpSecConnectionMigrationStatusMigrationStatusRollbackFailed,
	"NOT_APPLICABLE":   IpSecConnectionMigrationStatusMigrationStatusNotApplicable,
	"MANUAL":           IpSecConnectionMigrationStatusMigrationStatusManual,
	"VALIDATING":       IpSecConnectionMigrationStatusMigrationStatusValidating,
}

var mappingIpSecConnectionMigrationStatusMigrationStatusEnumLowerCase = map[string]IpSecConnectionMigrationStatusMigrationStatusEnum{
	"ready":            IpSecConnectionMigrationStatusMigrationStatusReady,
	"migrated":         IpSecConnectionMigrationStatusMigrationStatusMigrated,
	"migrating":        IpSecConnectionMigrationStatusMigrationStatusMigrating,
	"migration_failed": IpSecConnectionMigrationStatusMigrationStatusMigrationFailed,
	"rolled_back":      IpSecConnectionMigrationStatusMigrationStatusRolledBack,
	"rolling_back":     IpSecConnectionMigrationStatusMigrationStatusRollingBack,
	"rollback_failed":  IpSecConnectionMigrationStatusMigrationStatusRollbackFailed,
	"not_applicable":   IpSecConnectionMigrationStatusMigrationStatusNotApplicable,
	"manual":           IpSecConnectionMigrationStatusMigrationStatusManual,
	"validating":       IpSecConnectionMigrationStatusMigrationStatusValidating,
}

// GetIpSecConnectionMigrationStatusMigrationStatusEnumValues Enumerates the set of values for IpSecConnectionMigrationStatusMigrationStatusEnum
func GetIpSecConnectionMigrationStatusMigrationStatusEnumValues() []IpSecConnectionMigrationStatusMigrationStatusEnum {
	values := make([]IpSecConnectionMigrationStatusMigrationStatusEnum, 0)
	for _, v := range mappingIpSecConnectionMigrationStatusMigrationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetIpSecConnectionMigrationStatusMigrationStatusEnumStringValues Enumerates the set of values in String for IpSecConnectionMigrationStatusMigrationStatusEnum
func GetIpSecConnectionMigrationStatusMigrationStatusEnumStringValues() []string {
	return []string{
		"READY",
		"MIGRATED",
		"MIGRATING",
		"MIGRATION_FAILED",
		"ROLLED_BACK",
		"ROLLING_BACK",
		"ROLLBACK_FAILED",
		"NOT_APPLICABLE",
		"MANUAL",
		"VALIDATING",
	}
}

// GetMappingIpSecConnectionMigrationStatusMigrationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpSecConnectionMigrationStatusMigrationStatusEnum(val string) (IpSecConnectionMigrationStatusMigrationStatusEnum, bool) {
	enum, ok := mappingIpSecConnectionMigrationStatusMigrationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
