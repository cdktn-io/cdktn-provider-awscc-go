// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationDestinationConfigurationAuditLog struct {
	// Contains information about an audit log destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#destination AppfabricIngestionDestination#destination}
	Destination *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestination `field:"required" json:"destination" yaml:"destination"`
}

