// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationDestinationConfiguration struct {
	// Contains information about an audit log destination configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#audit_log AppfabricIngestionDestination#audit_log}
	AuditLog *AppfabricIngestionDestinationDestinationConfigurationAuditLog `field:"required" json:"auditLog" yaml:"auditLog"`
}

