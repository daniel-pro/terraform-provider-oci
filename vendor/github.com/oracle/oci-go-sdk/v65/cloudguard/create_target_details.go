// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateTargetDetails Parameters used to create a new target.
type CreateTargetDetails struct {

	// Display name for the target.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of resource that target support (COMPARTMENT/FACLOUD)
	TargetResourceType TargetResourceTypeEnum `mandatory:"true" json:"targetResourceType"`

	// Resource ID which the target uses to monitor
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// The target description.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// List of detector recipes to attach to target
	TargetDetectorRecipes []CreateTargetDetectorRecipeDetails `mandatory:"false" json:"targetDetectorRecipes"`

	// List of responder recipes to attach to target
	TargetResponderRecipes []CreateTargetResponderRecipeDetails `mandatory:"false" json:"targetResponderRecipes"`

	// The enablement state of the detector rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Should problems be emitted to OCI Events service?
	DoesEmitProblemsToEvents *bool `mandatory:"false" json:"doesEmitProblemsToEvents"`

	TargetDetails CreateTargetAdditionalDetails `mandatory:"false" json:"targetDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetResourceTypeEnum(string(m.TargetResourceType)); !ok && m.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", m.TargetResourceType, strings.Join(GetTargetResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateTargetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description              *string                              `json:"description"`
		TargetDetectorRecipes    []CreateTargetDetectorRecipeDetails  `json:"targetDetectorRecipes"`
		TargetResponderRecipes   []CreateTargetResponderRecipeDetails `json:"targetResponderRecipes"`
		LifecycleState           LifecycleStateEnum                   `json:"lifecycleState"`
		DoesEmitProblemsToEvents *bool                                `json:"doesEmitProblemsToEvents"`
		TargetDetails            createtargetadditionaldetails        `json:"targetDetails"`
		FreeformTags             map[string]string                    `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}    `json:"definedTags"`
		DisplayName              *string                              `json:"displayName"`
		CompartmentId            *string                              `json:"compartmentId"`
		TargetResourceType       TargetResourceTypeEnum               `json:"targetResourceType"`
		TargetResourceId         *string                              `json:"targetResourceId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TargetDetectorRecipes = make([]CreateTargetDetectorRecipeDetails, len(model.TargetDetectorRecipes))
	copy(m.TargetDetectorRecipes, model.TargetDetectorRecipes)
	m.TargetResponderRecipes = make([]CreateTargetResponderRecipeDetails, len(model.TargetResponderRecipes))
	copy(m.TargetResponderRecipes, model.TargetResponderRecipes)
	m.LifecycleState = model.LifecycleState

	m.DoesEmitProblemsToEvents = model.DoesEmitProblemsToEvents

	nn, e = model.TargetDetails.UnmarshalPolymorphicJSON(model.TargetDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetDetails = nn.(CreateTargetAdditionalDetails)
	} else {
		m.TargetDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TargetResourceType = model.TargetResourceType

	m.TargetResourceId = model.TargetResourceId

	return
}
