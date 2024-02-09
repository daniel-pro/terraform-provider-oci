// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAssetSourceDetails Asset source creation request.
type CreateAssetSourceDetails interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for the resource.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the environment.
	GetEnvironmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the inventory that will contain created assets.
	GetInventoryId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	GetAssetsCompartmentId() *string

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information. The name is generated by the service if it is not
	// explicitly provided.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the discovery schedule that is going to be attached to the created asset.
	GetDiscoveryScheduleId() *string

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	GetSystemTags() map[string]map[string]interface{}
}

type createassetsourcedetails struct {
	JsonData            []byte
	DisplayName         *string                           `mandatory:"false" json:"displayName"`
	DiscoveryScheduleId *string                           `mandatory:"false" json:"discoveryScheduleId"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	EnvironmentId       *string                           `mandatory:"true" json:"environmentId"`
	InventoryId         *string                           `mandatory:"true" json:"inventoryId"`
	AssetsCompartmentId *string                           `mandatory:"true" json:"assetsCompartmentId"`
	Type                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createassetsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateassetsourcedetails createassetsourcedetails
	s := struct {
		Model Unmarshalercreateassetsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.EnvironmentId = s.Model.EnvironmentId
	m.InventoryId = s.Model.InventoryId
	m.AssetsCompartmentId = s.Model.AssetsCompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DiscoveryScheduleId = s.Model.DiscoveryScheduleId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createassetsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DB":
		mm := CreateOracleDbAssetSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VMWARE":
		mm := CreateVmWareAssetSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWS":
		mm := CreateAwsAssetSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateAssetSourceDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createassetsourcedetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDiscoveryScheduleId returns DiscoveryScheduleId
func (m createassetsourcedetails) GetDiscoveryScheduleId() *string {
	return m.DiscoveryScheduleId
}

// GetFreeformTags returns FreeformTags
func (m createassetsourcedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createassetsourcedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m createassetsourcedetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetCompartmentId returns CompartmentId
func (m createassetsourcedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetEnvironmentId returns EnvironmentId
func (m createassetsourcedetails) GetEnvironmentId() *string {
	return m.EnvironmentId
}

// GetInventoryId returns InventoryId
func (m createassetsourcedetails) GetInventoryId() *string {
	return m.InventoryId
}

// GetAssetsCompartmentId returns AssetsCompartmentId
func (m createassetsourcedetails) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

func (m createassetsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createassetsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
