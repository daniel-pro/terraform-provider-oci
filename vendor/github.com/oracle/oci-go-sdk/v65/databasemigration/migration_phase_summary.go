// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MigrationPhaseSummary Migration Phase Summary of details.
type MigrationPhaseSummary struct {

	// ODMS Job phase name
	Name OdmsJobPhasesEnum `mandatory:"true" json:"name"`

	// Array of actions for the corresponding phase. Empty array would indicate there is no supported action for the phase.
	SupportedActions []OdmsPhaseActionsEnum `mandatory:"true" json:"supportedActions"`

	// Action recommended for this phase. If not included in the response, there is no recommended action for the phase.
	RecommendedAction OdmsPhaseActionsEnum `mandatory:"false" json:"recommendedAction,omitempty"`
}

func (m MigrationPhaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationPhaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdmsJobPhasesEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetOdmsJobPhasesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOdmsPhaseActionsEnum(string(m.RecommendedAction)); !ok && m.RecommendedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecommendedAction: %s. Supported values are: %s.", m.RecommendedAction, strings.Join(GetOdmsPhaseActionsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
