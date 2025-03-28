// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemediationRunStageSummary The summary of a remediation run stages.
type RemediationRunStageSummary struct {

	// The current status of remediation run stage.
	Status RemediationRunStageStatusEnum `mandatory:"true" json:"status"`

	// The creation date and time of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of the remediation run stage.
	Type RemediationRunStageTypeEnum `mandatory:"true" json:"type"`

	// The Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
	RemediationRunId *string `mandatory:"true" json:"remediationRunId"`

	// The date and time of the start of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time of the finish of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Information about the current step within the stage.
	Summary *string `mandatory:"false" json:"summary"`
}

func (m RemediationRunStageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemediationRunStageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemediationRunStageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRemediationRunStageStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRemediationRunStageTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemediationRunStageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
