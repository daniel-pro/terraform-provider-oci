// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"

	oci_core "github.com/oracle/oci-go-sdk/v35/core"
)

func init() {
	RegisterResource("oci_core_default_route_table", DefaultCoreRouteTableResource())
}

func DefaultCoreRouteTableResource() *schema.Resource {
	defaultResourceSchema := ConvertToDefaultVcnResourceSchema(CoreRouteTableResource())

	defaultResourceSchema.Create = createDefaultRouteTable
	defaultResourceSchema.Delete = deleteDefaultRouteTable

	return defaultResourceSchema
}

type DefaultRouteTableResourceCrud struct {
	CoreRouteTableResourceCrud
}

func createDefaultRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	return CreateResource(d, sync)
}

func deleteDefaultRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

func (s *DefaultRouteTableResourceCrud) Create() error {
	// If we are creating a default resource, then don't have to
	// actually create it. Just set the ID and update it.
	if defaultId, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		s.D.SetId(defaultId.(string))
		return s.Update()
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

// This clears out all of the route table rules and sets it to empty
// This is used to clear out default Route Table resources that can't otherwise be deleted
func (s *DefaultRouteTableResourceCrud) reset() error {
	request := oci_core.UpdateRouteTableRequest{}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RouteRules = []oci_core.RouteRule{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *DefaultRouteTableResourceCrud) Delete() error {
	if _, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		// We can't actually delete a default resource.
		// Clear out its settings and mark it as deleted.
		err := s.reset()
		s.D.Set("state", s.DeletedTarget()[0])
		return err
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

// You can't actually delete a default resource, so the creation target
// states are valid states for when waiting for delete to complete
func (s *DefaultRouteTableResourceCrud) DeletedPending() []string {
	return s.CreatedTarget()
}

func (s *DefaultRouteTableResourceCrud) DeletedTarget() []string {
	return s.CreatedTarget()
}
