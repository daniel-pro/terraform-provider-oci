// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LicenseEnterpriseExtensibilityConfigDetails A configuration of the LICENSE_ENTERPRISE_EXTENSIBILITY type, consists of a boolean which determines
// whether enterprise extensibility is enabled.
type LicenseEnterpriseExtensibilityConfigDetails struct {

	// The Unique Oracle ID (OCID) that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// True if enterprise extensibility is enabled, false if it is not enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the configuration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Config was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the configuration.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m LicenseEnterpriseExtensibilityConfigDetails) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m LicenseEnterpriseExtensibilityConfigDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m LicenseEnterpriseExtensibilityConfigDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m LicenseEnterpriseExtensibilityConfigDetails) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m LicenseEnterpriseExtensibilityConfigDetails) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m LicenseEnterpriseExtensibilityConfigDetails) GetLifecycleState() ConfigLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m LicenseEnterpriseExtensibilityConfigDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m LicenseEnterpriseExtensibilityConfigDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m LicenseEnterpriseExtensibilityConfigDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m LicenseEnterpriseExtensibilityConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LicenseEnterpriseExtensibilityConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LicenseEnterpriseExtensibilityConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLicenseEnterpriseExtensibilityConfigDetails LicenseEnterpriseExtensibilityConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeLicenseEnterpriseExtensibilityConfigDetails
	}{
		"LICENSE_ENTERPRISE_EXTENSIBILITY",
		(MarshalTypeLicenseEnterpriseExtensibilityConfigDetails)(m),
	}

	return json.Marshal(&s)
}
