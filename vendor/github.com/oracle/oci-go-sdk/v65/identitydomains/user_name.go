// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserName A complex attribute that contains attributes representing the name
// **SCIM++ Properties:**
//   - idcsCsvAttributeNameMappings: [[columnHeaderName:Formatted Name, mapsTo:name.formatted], [columnHeaderName:Honorific Prefix, mapsTo:name.honorificPrefix], [columnHeaderName:First Name, mapsTo:name.givenName], [columnHeaderName:Middle Name, mapsTo:name.middleName], [columnHeaderName:Last Name, mapsTo:name.familyName], [columnHeaderName:Honorific Suffix, mapsTo:name.honorificSuffix]]
//   - idcsPii: true
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type UserName struct {

	// Full name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Formatted *string `mandatory:"false" json:"formatted"`

	// Last name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Last Name
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FamilyName *string `mandatory:"false" json:"familyName"`

	// First name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: First Name
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GivenName *string `mandatory:"false" json:"givenName"`

	// Middle name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Middle Name
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MiddleName *string `mandatory:"false" json:"middleName"`

	// Prefix
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: Honorific Prefix
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HonorificPrefix *string `mandatory:"false" json:"honorificPrefix"`

	// Suffix
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: Honorific Suffix
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HonorificSuffix *string `mandatory:"false" json:"honorificSuffix"`
}

func (m UserName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
