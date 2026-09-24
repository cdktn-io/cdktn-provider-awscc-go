// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition


type MgnNetworkMigrationDefinitionSourceConfigurations struct {
	// The source environment type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#source_environment MgnNetworkMigrationDefinition#source_environment}
	SourceEnvironment *string `field:"required" json:"sourceEnvironment" yaml:"sourceEnvironment"`
	// S3 configuration for source network data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#source_s3_configuration MgnNetworkMigrationDefinition#source_s3_configuration}
	SourceS3Configuration *MgnNetworkMigrationDefinitionSourceConfigurationsSourceS3Configuration `field:"required" json:"sourceS3Configuration" yaml:"sourceS3Configuration"`
}

