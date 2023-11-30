// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOracleMigrationDetails Create Migration resource parameters.
type UpdateOracleMigrationDetails struct {

	// An object's Description.
	Description *string `mandatory:"false" json:"description"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// OCI resource ID.
	SourceDatabaseConnectionId *string `mandatory:"false" json:"sourceDatabaseConnectionId"`

	// OCI resource ID.
	TargetDatabaseConnectionId *string `mandatory:"false" json:"targetDatabaseConnectionId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataTransferMediumDetails UpdateOracleDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *UpdateOracleInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *UpdateOracleAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	HubDetails *UpdateGoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *UpdateOracleGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// OCI resource ID.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// Migration type (ONLINE/OFFLINE).
	Type MigrationTypesEnum `mandatory:"false" json:"type,omitempty"`
}

// GetDescription returns Description
func (m UpdateOracleMigrationDetails) GetDescription() *string {
	return m.Description
}

// GetType returns Type
func (m UpdateOracleMigrationDetails) GetType() MigrationTypesEnum {
	return m.Type
}

// GetDisplayName returns DisplayName
func (m UpdateOracleMigrationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m UpdateOracleMigrationDetails) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m UpdateOracleMigrationDetails) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetFreeformTags returns FreeformTags
func (m UpdateOracleMigrationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOracleMigrationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateOracleMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleMigrationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleMigrationDetails UpdateOracleMigrationDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeUpdateOracleMigrationDetails
	}{
		"ORACLE",
		(MarshalTypeUpdateOracleMigrationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateOracleMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                         *string                               `json:"description"`
		Type                                MigrationTypesEnum                    `json:"type"`
		DisplayName                         *string                               `json:"displayName"`
		SourceDatabaseConnectionId          *string                               `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId          *string                               `json:"targetDatabaseConnectionId"`
		FreeformTags                        map[string]string                     `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{}     `json:"definedTags"`
		DataTransferMediumDetails           updateoracledatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings                 *UpdateOracleInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings                     *UpdateOracleAdvisorSettings          `json:"advisorSettings"`
		HubDetails                          *UpdateGoldenGateHubDetails           `json:"hubDetails"`
		GgsDetails                          *UpdateOracleGgsDeploymentDetails     `json:"ggsDetails"`
		SourceContainerDatabaseConnectionId *string                               `json:"sourceContainerDatabaseConnectionId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Type = model.Type

	m.DisplayName = model.DisplayName

	m.SourceDatabaseConnectionId = model.SourceDatabaseConnectionId

	m.TargetDatabaseConnectionId = model.TargetDatabaseConnectionId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.DataTransferMediumDetails.UnmarshalPolymorphicJSON(model.DataTransferMediumDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataTransferMediumDetails = nn.(UpdateOracleDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

	m.SourceContainerDatabaseConnectionId = model.SourceContainerDatabaseConnectionId

	return
}
