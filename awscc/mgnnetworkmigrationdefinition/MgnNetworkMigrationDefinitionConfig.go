// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MgnNetworkMigrationDefinitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the network migration definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#name MgnNetworkMigrationDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of source configurations for the network migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#source_configurations MgnNetworkMigrationDefinition#source_configurations}
	SourceConfigurations interface{} `field:"required" json:"sourceConfigurations" yaml:"sourceConfigurations"`
	// The target network configuration including topology and CIDR ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#target_network MgnNetworkMigrationDefinition#target_network}
	TargetNetwork *MgnNetworkMigrationDefinitionTargetNetwork `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// The S3 configuration for storing the target network artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#target_s3_configuration MgnNetworkMigrationDefinition#target_s3_configuration}
	TargetS3Configuration *MgnNetworkMigrationDefinitionTargetS3Configuration `field:"required" json:"targetS3Configuration" yaml:"targetS3Configuration"`
	// A description of the network migration definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#description MgnNetworkMigrationDefinition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Scope tags for the network migration definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#scope_tags MgnNetworkMigrationDefinition#scope_tags}
	ScopeTags *map[string]*string `field:"optional" json:"scopeTags" yaml:"scopeTags"`
	// Tags to assign to the network migration definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#tags MgnNetworkMigrationDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The target deployment configuration for the migrated network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mgn_network_migration_definition#target_deployment MgnNetworkMigrationDefinition#target_deployment}
	TargetDeployment *string `field:"optional" json:"targetDeployment" yaml:"targetDeployment"`
}

