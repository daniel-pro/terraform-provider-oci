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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Base64SslCertDetails The 64-based encoded SslCert.
type Base64SslCertDetails struct {

	// The base 64 encoded string for the Ssl Certificate.
	CertContent *string `mandatory:"false" json:"certContent"`
}

func (m Base64SslCertDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Base64SslCertDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Base64SslCertDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBase64SslCertDetails Base64SslCertDetails
	s := struct {
		DiscriminatorParam string `json:"contentType"`
		MarshalTypeBase64SslCertDetails
	}{
		"BASE64ENCODED",
		(MarshalTypeBase64SslCertDetails)(m),
	}

	return json.Marshal(&s)
}
