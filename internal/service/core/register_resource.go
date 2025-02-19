// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_core_app_catalog_listing_resource_version_agreement", AppCatalogListingResourceVersionAgreementResource())
	tfresource.RegisterResource("oci_core_app_catalog_subscription", CoreAppCatalogSubscriptionResource())
	tfresource.RegisterResource("oci_core_boot_volume", CoreBootVolumeResource())
	tfresource.RegisterResource("oci_core_boot_volume_backup", CoreBootVolumeBackupResource())
	tfresource.RegisterResource("oci_core_byoasn", CoreByoasnResource())
	tfresource.RegisterResource("oci_core_capture_filter", CoreCaptureFilterResource())
	tfresource.RegisterResource("oci_core_cluster_network", CoreClusterNetworkResource())
	tfresource.RegisterResource("oci_core_compute_capacity_report", CoreComputeCapacityReportResource())
	tfresource.RegisterResource("oci_core_compute_capacity_reservation", CoreComputeCapacityReservationResource())
	tfresource.RegisterResource("oci_core_compute_capacity_topology", CoreComputeCapacityTopologyResource())
	tfresource.RegisterResource("oci_core_compute_cluster", CoreComputeClusterResource())
	tfresource.RegisterResource("oci_core_compute_gpu_memory_cluster", CoreComputeGpuMemoryClusterResource())
	tfresource.RegisterResource("oci_core_compute_gpu_memory_fabric", CoreComputeGpuMemoryFabricResource())
	tfresource.RegisterResource("oci_core_compute_image_capability_schema", CoreComputeImageCapabilitySchemaResource())
	tfresource.RegisterResource("oci_core_console_history", CoreConsoleHistoryResource())
	tfresource.RegisterResource("oci_core_cpe", CoreCpeResource())
	tfresource.RegisterResource("oci_core_cross_connect", CoreCrossConnectResource())
	tfresource.RegisterResource("oci_core_default_dhcp_options", DefaultCoreDhcpOptionsResource())
	tfresource.RegisterResource("oci_core_default_route_table", DefaultCoreRouteTableResource())
	tfresource.RegisterResource("oci_core_default_security_list", CoreDefaultSecurityListResource())
	tfresource.RegisterResource("oci_core_cross_connect_group", CoreCrossConnectGroupResource())
	tfresource.RegisterResource("oci_core_dedicated_vm_host", CoreDedicatedVmHostResource())
	tfresource.RegisterResource("oci_core_dhcp_options", CoreDhcpOptionsResource())
	tfresource.RegisterResource("oci_core_drg", CoreDrgResource())
	tfresource.RegisterResource("oci_core_drg_attachment", CoreDrgAttachmentResource())
	tfresource.RegisterResource("oci_core_drg_attachments_list", CoreDrgAttachmentsListResource())
	tfresource.RegisterResource("oci_core_drg_attachment_management", CoreDrgAttachmentManagementResource())
	tfresource.RegisterResource("oci_core_drg_route_distribution", CoreDrgRouteDistributionResource())
	tfresource.RegisterResource("oci_core_drg_route_distribution_statement", CoreDrgRouteDistributionStatementResource())
	tfresource.RegisterResource("oci_core_drg_route_table", CoreDrgRouteTableResource())
	tfresource.RegisterResource("oci_core_drg_route_table_route_rule", CoreDrgRouteTableRouteRuleResource())
	tfresource.RegisterResource("oci_core_image", CoreImageResource())
	tfresource.RegisterResource("oci_core_instance", CoreInstanceResource())
	tfresource.RegisterResource("oci_core_instance_configuration", CoreInstanceConfigurationResource())
	tfresource.RegisterResource("oci_core_instance_console_connection", CoreInstanceConsoleConnectionResource())
	tfresource.RegisterResource("oci_core_instance_maintenance_event", CoreInstanceMaintenanceEventResource())
	tfresource.RegisterResource("oci_core_instance_pool", CoreInstancePoolResource())
	tfresource.RegisterResource("oci_core_instance_pool_instance", CoreInstancePoolInstanceResource())
	tfresource.RegisterResource("oci_core_internet_gateway", CoreInternetGatewayResource())
	tfresource.RegisterResource("oci_core_ipsec", CoreIpSecConnectionResource())
	tfresource.RegisterResource("oci_core_ipsec_connection_tunnel_management", CoreIpSecConnectionTunnelManagementResource())
	tfresource.RegisterResource("oci_core_ipv6", CoreIpv6Resource())
	tfresource.RegisterResource("oci_core_listing_resource_version_agreement", AppCatalogListingResourceVersionAgreementResource())
	tfresource.RegisterResource("oci_core_local_peering_gateway", CoreLocalPeeringGatewayResource())
	tfresource.RegisterResource("oci_core_nat_gateway", CoreNatGatewayResource())
	tfresource.RegisterResource("oci_core_network_security_group", CoreNetworkSecurityGroupResource())
	tfresource.RegisterResource("oci_core_network_security_group_security_rule", CoreNetworkSecurityGroupSecurityRuleResource())
	tfresource.RegisterResource("oci_core_private_ip", CorePrivateIpResource())
	tfresource.RegisterResource("oci_core_public_ip", CorePublicIpResource())
	tfresource.RegisterResource("oci_core_public_ip_pool", CorePublicIpPoolResource())
	tfresource.RegisterResource("oci_core_public_ip_pool_capacity", PublicIpPoolCapacityResource())
	tfresource.RegisterResource("oci_core_remote_peering_connection", CoreRemotePeeringConnectionResource())
	tfresource.RegisterResource("oci_core_route_table", CoreRouteTableResource())
	tfresource.RegisterResource("oci_core_route_table_attachment", CoreRouteTableAttachmentResource())
	tfresource.RegisterResource("oci_core_security_list", CoreSecurityListResource())
	tfresource.RegisterResource("oci_core_service_gateway", CoreServiceGatewayResource())
	tfresource.RegisterResource("oci_core_shape_management", CoreShapeResource())
	tfresource.RegisterResource("oci_core_subnet", CoreSubnetResource())
	tfresource.RegisterResource("oci_core_vcn", CoreVcnResource())
	tfresource.RegisterResource("oci_core_virtual_circuit", CoreVirtualCircuitResource())
	tfresource.RegisterResource("oci_core_vlan", CoreVlanResource())
	tfresource.RegisterResource("oci_core_vnic_attachment", CoreVnicAttachmentResource())
	tfresource.RegisterResource("oci_core_volume", CoreVolumeResource())
	tfresource.RegisterResource("oci_core_volume_attachment", CoreVolumeAttachmentResource())
	tfresource.RegisterResource("oci_core_volume_backup", CoreVolumeBackupResource())
	tfresource.RegisterResource("oci_core_volume_backup_policy", CoreVolumeBackupPolicyResource())
	tfresource.RegisterResource("oci_core_volume_backup_policy_assignment", CoreVolumeBackupPolicyAssignmentResource())
	tfresource.RegisterResource("oci_core_volume_group", CoreVolumeGroupResource())
	tfresource.RegisterResource("oci_core_volume_group_backup", CoreVolumeGroupBackupResource())
	tfresource.RegisterResource("oci_core_vtap", CoreVtapResource())
}
