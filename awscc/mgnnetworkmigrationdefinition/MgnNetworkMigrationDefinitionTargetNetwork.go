// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition


type MgnNetworkMigrationDefinitionTargetNetwork struct {
	// The network topology type for the target environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#topology MgnNetworkMigrationDefinition#topology}
	Topology *string `field:"required" json:"topology" yaml:"topology"`
	// The CIDR block for inbound traffic in the target network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#inbound_cidr MgnNetworkMigrationDefinition#inbound_cidr}
	InboundCidr *string `field:"optional" json:"inboundCidr" yaml:"inboundCidr"`
	// The CIDR block for inspection traffic in the target network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#inspection_cidr MgnNetworkMigrationDefinition#inspection_cidr}
	InspectionCidr *string `field:"optional" json:"inspectionCidr" yaml:"inspectionCidr"`
	// The CIDR block for outbound traffic in the target network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#outbound_cidr MgnNetworkMigrationDefinition#outbound_cidr}
	OutboundCidr *string `field:"optional" json:"outboundCidr" yaml:"outboundCidr"`
}

