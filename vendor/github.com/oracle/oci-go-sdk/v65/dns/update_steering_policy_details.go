// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSteeringPolicyDetails The body for updating a steering policy. New rules and answers provided in the request will
// replace the existing rules and answers in the policy.
//
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateSteeringPolicyDetails struct {

	// A user-friendly name for the steering policy. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Time To Live (TTL) for responses from the steering policy, in seconds.
	// If not specified during creation, a value of 30 seconds will be used.
	Ttl *int `mandatory:"false" json:"ttl"`

	// The OCID of the health check monitor providing health data about the answers of the
	// steering policy. A steering policy answer with `rdata` matching a monitored endpoint
	// will use the health data of that endpoint. A steering policy answer with `rdata` not
	// matching any monitored endpoint will be assumed healthy.
	//
	// **Note:** To use the Health Check monitoring feature in a steering policy, a monitor
	// must be created using the Health Checks service first. For more information on how to
	// create a monitor, please see Managing Health Checks (https://docs.oracle.com/iaas/Content/HealthChecks/Tasks/managinghealthchecks.htm).
	HealthCheckMonitorId *string `mandatory:"false" json:"healthCheckMonitorId"`

	// A set of predefined rules based on the desired purpose of the steering policy. Each
	// template utilizes Traffic Management's rules in a different order to produce the desired
	// results when answering DNS queries.
	//
	// **Example:** The `FAILOVER` template determines answers by filtering the policy's answers
	// using the `FILTER` rule first, then the following rules in succession: `HEALTH`, `PRIORITY`,
	// and `LIMIT`. This gives the domain dynamic failover capability.
	//
	// It is **strongly recommended** to use a template other than `CUSTOM` when creating
	// a steering policy.
	//
	// All templates require the rule order to begin with an unconditional `FILTER` rule that keeps
	// answers contingent upon `answer.isDisabled != true`, except for `CUSTOM`. A defined
	// `HEALTH` rule must follow the `FILTER` rule if the policy references a `healthCheckMonitorId`.
	// The last rule of a template must must be a `LIMIT` rule. For more information about templates
	// and code examples, see Traffic Management API Guide (https://docs.oracle.com/iaas/Content/TrafficManagement/Concepts/trafficmanagementapi.htm).
	// **Template Types**
	// * `FAILOVER` - Uses health check information on your endpoints to determine which DNS answers
	// to serve. If an endpoint fails a health check, the answer for that endpoint will be removed
	// from the list of available answers until the endpoint is detected as healthy.
	//
	// * `LOAD_BALANCE` - Distributes web traffic to specified endpoints based on defined weights.
	//
	// * `ROUTE_BY_GEO` - Answers DNS queries based on the query's geographic location. For a list of geographic
	// locations to route by, see Traffic Management Geographic Locations (https://docs.oracle.com/iaas/Content/TrafficManagement/Reference/trafficmanagementgeo.htm).
	//
	// * `ROUTE_BY_ASN` - Answers DNS queries based on the query's originating ASN.
	//
	// * `ROUTE_BY_IP` - Answers DNS queries based on the query's IP address.
	//
	// * `CUSTOM` - Allows a customized configuration of rules.
	Template UpdateSteeringPolicyDetailsTemplateEnum `mandatory:"false" json:"template,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The set of all answers that can potentially issue from the steering policy.
	Answers []SteeringPolicyAnswer `mandatory:"false" json:"answers"`

	// The series of rules that will be processed in sequence to reduce the pool of answers
	// to a response for any given request.
	//
	// The first rule receives a shuffled list of all answers, and every other rule receives
	// the list of answers emitted by the one preceding it. The last rule populates the
	// response.
	Rules []SteeringPolicyRule `mandatory:"false" json:"rules"`
}

func (m UpdateSteeringPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSteeringPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateSteeringPolicyDetailsTemplateEnum(string(m.Template)); !ok && m.Template != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Template: %s. Supported values are: %s.", m.Template, strings.Join(GetUpdateSteeringPolicyDetailsTemplateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSteeringPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                                 `json:"displayName"`
		Ttl                  *int                                    `json:"ttl"`
		HealthCheckMonitorId *string                                 `json:"healthCheckMonitorId"`
		Template             UpdateSteeringPolicyDetailsTemplateEnum `json:"template"`
		FreeformTags         map[string]string                       `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{}       `json:"definedTags"`
		Answers              []SteeringPolicyAnswer                  `json:"answers"`
		Rules                []steeringpolicyrule                    `json:"rules"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Ttl = model.Ttl

	m.HealthCheckMonitorId = model.HealthCheckMonitorId

	m.Template = model.Template

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Answers = make([]SteeringPolicyAnswer, len(model.Answers))
	copy(m.Answers, model.Answers)
	m.Rules = make([]SteeringPolicyRule, len(model.Rules))
	for i, n := range model.Rules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Rules[i] = nn.(SteeringPolicyRule)
		} else {
			m.Rules[i] = nil
		}
	}
	return
}

// UpdateSteeringPolicyDetailsTemplateEnum Enum with underlying type: string
type UpdateSteeringPolicyDetailsTemplateEnum string

// Set of constants representing the allowable values for UpdateSteeringPolicyDetailsTemplateEnum
const (
	UpdateSteeringPolicyDetailsTemplateFailover    UpdateSteeringPolicyDetailsTemplateEnum = "FAILOVER"
	UpdateSteeringPolicyDetailsTemplateLoadBalance UpdateSteeringPolicyDetailsTemplateEnum = "LOAD_BALANCE"
	UpdateSteeringPolicyDetailsTemplateRouteByGeo  UpdateSteeringPolicyDetailsTemplateEnum = "ROUTE_BY_GEO"
	UpdateSteeringPolicyDetailsTemplateRouteByAsn  UpdateSteeringPolicyDetailsTemplateEnum = "ROUTE_BY_ASN"
	UpdateSteeringPolicyDetailsTemplateRouteByIp   UpdateSteeringPolicyDetailsTemplateEnum = "ROUTE_BY_IP"
	UpdateSteeringPolicyDetailsTemplateCustom      UpdateSteeringPolicyDetailsTemplateEnum = "CUSTOM"
)

var mappingUpdateSteeringPolicyDetailsTemplateEnum = map[string]UpdateSteeringPolicyDetailsTemplateEnum{
	"FAILOVER":     UpdateSteeringPolicyDetailsTemplateFailover,
	"LOAD_BALANCE": UpdateSteeringPolicyDetailsTemplateLoadBalance,
	"ROUTE_BY_GEO": UpdateSteeringPolicyDetailsTemplateRouteByGeo,
	"ROUTE_BY_ASN": UpdateSteeringPolicyDetailsTemplateRouteByAsn,
	"ROUTE_BY_IP":  UpdateSteeringPolicyDetailsTemplateRouteByIp,
	"CUSTOM":       UpdateSteeringPolicyDetailsTemplateCustom,
}

var mappingUpdateSteeringPolicyDetailsTemplateEnumLowerCase = map[string]UpdateSteeringPolicyDetailsTemplateEnum{
	"failover":     UpdateSteeringPolicyDetailsTemplateFailover,
	"load_balance": UpdateSteeringPolicyDetailsTemplateLoadBalance,
	"route_by_geo": UpdateSteeringPolicyDetailsTemplateRouteByGeo,
	"route_by_asn": UpdateSteeringPolicyDetailsTemplateRouteByAsn,
	"route_by_ip":  UpdateSteeringPolicyDetailsTemplateRouteByIp,
	"custom":       UpdateSteeringPolicyDetailsTemplateCustom,
}

// GetUpdateSteeringPolicyDetailsTemplateEnumValues Enumerates the set of values for UpdateSteeringPolicyDetailsTemplateEnum
func GetUpdateSteeringPolicyDetailsTemplateEnumValues() []UpdateSteeringPolicyDetailsTemplateEnum {
	values := make([]UpdateSteeringPolicyDetailsTemplateEnum, 0)
	for _, v := range mappingUpdateSteeringPolicyDetailsTemplateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSteeringPolicyDetailsTemplateEnumStringValues Enumerates the set of values in String for UpdateSteeringPolicyDetailsTemplateEnum
func GetUpdateSteeringPolicyDetailsTemplateEnumStringValues() []string {
	return []string{
		"FAILOVER",
		"LOAD_BALANCE",
		"ROUTE_BY_GEO",
		"ROUTE_BY_ASN",
		"ROUTE_BY_IP",
		"CUSTOM",
	}
}

// GetMappingUpdateSteeringPolicyDetailsTemplateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSteeringPolicyDetailsTemplateEnum(val string) (UpdateSteeringPolicyDetailsTemplateEnum, bool) {
	enum, ok := mappingUpdateSteeringPolicyDetailsTemplateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
