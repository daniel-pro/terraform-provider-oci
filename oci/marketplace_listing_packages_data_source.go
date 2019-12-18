// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_marketplace "github.com/oracle/oci-go-sdk/marketplace"
)

func MarketplaceListingPackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceListingPackages,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_type": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"IMAGE",
					"ORCHESTRATION",
				}, true),
			},
			"package_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"listing_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMarketplaceListingPackages(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient

	return ReadResource(sync)
}

type MarketplaceListingPackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListPackagesResponse
}

func (s *MarketplaceListingPackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingPackagesDataSourceCrud) Get() error {
	request := oci_marketplace.ListPackagesRequest{}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		tmp := packageType.(string)
		request.PackageType = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "marketplace")

	response, err := s.Client.ListPackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplaceListingPackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		listingPackage := map[string]interface{}{
			"listing_id": *r.ListingId,
		}

		if r.PackageVersion != nil {
			listingPackage["package_version"] = *r.PackageVersion
		}

		if r.ResourceId != nil {
			listingPackage["resource_id"] = *r.ResourceId
		}

		if r.TimeCreated != nil {
			listingPackage["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, listingPackage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MarketplaceListingPackagesDataSource().Schema["listing_packages"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("listing_packages", resources); err != nil {
		return err
	}

	return nil
}