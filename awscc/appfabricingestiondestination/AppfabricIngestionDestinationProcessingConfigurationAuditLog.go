// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationProcessingConfigurationAuditLog struct {
	// The format in which the audit logs need to be formatted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#format AppfabricIngestionDestination#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The event schema in which the audit logs need to be formatted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#schema AppfabricIngestionDestination#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
}

