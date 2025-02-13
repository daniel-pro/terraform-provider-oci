// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeUserAssessmentPasswordExpiryDateAnalyticDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.user_assessment_id}`},
		"access_level":                   acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"time_password_expiry_less_than": acctest.Representation{RepType: acctest.Optional, Create: `timePasswordExpiryLessThan`},
		"user_category":                  acctest.Representation{RepType: acctest.Optional, Create: `userCategory`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentPasswordExpiryDateAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentPasswordExpiryDateAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	assessmentId := utils.GetEnvSettingWithBlankDefault("data_safe_user_assessment_id")
	userAssessmentIdVariableStr := fmt.Sprintf("variable \"user_assessment_id\" { default = \"%s\" }\n", assessmentId)

	datasourceName := "data.oci_data_safe_user_assessment_password_expiry_date_analytics.test_user_assessment_password_expiry_date_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_password_expiry_date_analytics", "test_user_assessment_password_expiry_date_analytics", acctest.Required, acctest.Create, DataSafeUserAssessmentPasswordExpiryDateAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + userAssessmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_aggregations.#"),
				resource.TestCheckResourceAttr(datasourceName, "user_aggregations.0.items.#", "3"),
			),
		},
	})
}
