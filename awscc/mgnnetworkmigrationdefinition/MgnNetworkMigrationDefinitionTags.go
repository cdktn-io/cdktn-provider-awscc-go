// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition


type MgnNetworkMigrationDefinitionTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mgn_network_migration_definition#key MgnNetworkMigrationDefinition#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mgn_network_migration_definition#value MgnNetworkMigrationDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

