// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v48/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SecurityAssessmentRequiredOnlyResource = SecurityAssessmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Required, Create, securityAssessmentRepresentation)

	SecurityAssessmentResourceConfig = SecurityAssessmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Optional, Update, securityAssessmentRepresentation)

	securityAssessmentSingularDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id": Representation{repType: Required, create: `${oci_data_safe_security_assessment.test_security_assessment.id}`},
	}

	securityAssessmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"target_id":      Representation{repType: Required, create: `${oci_data_safe_target_database.test_target_database.id}`},
		"type":           Representation{repType: Optional, create: `SAVED`},
		"filter":         RepresentationGroup{Required, securityAssessmentDataSourceFilterRepresentation}}
	securityAssessmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_data_safe_security_assessment.test_security_assessment.id}`}},
	}

	securityAssessmentRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"target_id":      Representation{repType: Required, create: `${oci_data_safe_target_database.test_target_database.id}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `EBS assessment`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      RepresentationGroup{Required, ignoreChangesSecurityAssessmentRepresentation},
	}

	securityAssessmentChangeCompartmentRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"target_id":      Representation{repType: Required, create: `${oci_data_safe_target_database.test_target_database.id}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `EBS assessment`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"schedule":       Representation{repType: Optional, create: `schedule`, update: `schedule2`},
		"lifecycle":      RepresentationGroup{Required, ignoreChangesSecurityAssessmentRepresentation},
	}

	ignoreChangesSecurityAssessmentRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	SecurityAssessmentResourceDependencies = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", Required, Create, targetDatabaseRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_security_assessment.test_security_assessment"
	datasourceName := "data.oci_data_safe_security_assessments.test_security_assessments"
	singularDatasourceName := "data.oci_data_safe_security_assessment.test_security_assessment"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+SecurityAssessmentResourceDependencies+
		generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Optional, Create, securityAssessmentRepresentation), "datasafe", "securityAssessment", t)

	ResourceTest(t, testAccCheckDataSafeSecurityAssessmentDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + SecurityAssessmentResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Required, Create, securityAssessmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + SecurityAssessmentResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + SecurityAssessmentResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Optional, Create, securityAssessmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EBS assessment"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SecurityAssessmentResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Optional, Create,
					representationCopyWithNewProperties(securityAssessmentRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EBS assessment"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SecurityAssessmentResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Required, Update, securityAssessmentChangeCompartmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_data_safe_security_assessments", "test_security_assessments", Optional, Update, securityAssessmentDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAssessmentResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Optional, Update, securityAssessmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttr(datasourceName, "type", "SAVED"),

				resource.TestCheckResourceAttr(datasourceName, "security_assessments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_assessments.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "security_assessments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.is_baseline"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "security_assessments.0.statistics.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.triggered_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", Required, Create, securityAssessmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAssessmentRequiredOnlyResource,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_assessment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_baseline"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statistics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_ids.#", "1"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "target_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "triggered_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SecurityAssessmentResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"target_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataSafeSecurityAssessmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_security_assessment" {
			noResourceFound = false
			request := oci_data_safe.GetSecurityAssessmentRequest{}

			tmp := rs.Primary.ID
			request.SecurityAssessmentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "data_safe")

			_, err := client.GetSecurityAssessment(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DataSafeSecurityAssessment") {
		resource.AddTestSweepers("DataSafeSecurityAssessment", &resource.Sweeper{
			Name:         "DataSafeSecurityAssessment",
			Dependencies: DependencyGraph["securityassessment"],
			F:            sweepDataSafeSecurityAssessmentResource,
		})
	}
}

func sweepDataSafeSecurityAssessmentResource(compartment string) error {
	dataSafeClient := GetTestClients(&schema.ResourceData{}).dataSafeClient()
	securityAssessmentIds, err := getSecurityAssessmentIds(compartment)
	if err != nil {
		return err
	}
	for _, securityAssessmentId := range securityAssessmentIds {
		if ok := SweeperDefaultResourceId[securityAssessmentId]; !ok {
			deleteSecurityAssessmentRequest := oci_data_safe.DeleteSecurityAssessmentRequest{}

			deleteSecurityAssessmentRequest.SecurityAssessmentId = &securityAssessmentId

			deleteSecurityAssessmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSecurityAssessment(context.Background(), deleteSecurityAssessmentRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityAssessment %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityAssessmentId, error)
				continue
			}
		}
	}
	return nil
}

func getSecurityAssessmentIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "SecurityAssessmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := GetTestClients(&schema.ResourceData{}).dataSafeClient()

	listSecurityAssessmentsRequest := oci_data_safe.ListSecurityAssessmentsRequest{}
	listSecurityAssessmentsRequest.CompartmentId = &compartmentId
	listSecurityAssessmentsResponse, err := dataSafeClient.ListSecurityAssessments(context.Background(), listSecurityAssessmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityAssessment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityAssessment := range listSecurityAssessmentsResponse.Items {
		id := *securityAssessment.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "SecurityAssessmentId", id)
	}
	return resourceIds, nil
}
