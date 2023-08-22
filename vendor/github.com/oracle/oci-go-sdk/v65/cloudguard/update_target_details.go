// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateTargetDetails The information to be updated.
type UpdateTargetDetails struct {

	// Display name of a target.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the Target.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The details of target detector recipes to be updated.
	TargetDetectorRecipes []UpdateTargetDetectorRecipe `mandatory:"false" json:"targetDetectorRecipes"`

	// The details of target responder recipes to be updated.
	TargetResponderRecipes []UpdateTargetResponderRecipe `mandatory:"false" json:"targetResponderRecipes"`

	// Emit problems to OCI Events
	DoesEmitProblemsToEvents *bool `mandatory:"false" json:"doesEmitProblemsToEvents"`

	TargetDetails UpdateTargetAdditionalDetails `mandatory:"false" json:"targetDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTargetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		LifecycleState           LifecycleStateEnum                `json:"lifecycleState"`
		TargetDetectorRecipes    []UpdateTargetDetectorRecipe      `json:"targetDetectorRecipes"`
		TargetResponderRecipes   []UpdateTargetResponderRecipe     `json:"targetResponderRecipes"`
		DoesEmitProblemsToEvents *bool                             `json:"doesEmitProblemsToEvents"`
		TargetDetails            updatetargetadditionaldetails     `json:"targetDetails"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.TargetDetectorRecipes = make([]UpdateTargetDetectorRecipe, len(model.TargetDetectorRecipes))
	copy(m.TargetDetectorRecipes, model.TargetDetectorRecipes)
	m.TargetResponderRecipes = make([]UpdateTargetResponderRecipe, len(model.TargetResponderRecipes))
	copy(m.TargetResponderRecipes, model.TargetResponderRecipes)
	m.DoesEmitProblemsToEvents = model.DoesEmitProblemsToEvents

	nn, e = model.TargetDetails.UnmarshalPolymorphicJSON(model.TargetDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetDetails = nn.(UpdateTargetAdditionalDetails)
	} else {
		m.TargetDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
