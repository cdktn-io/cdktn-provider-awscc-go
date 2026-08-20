// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraIndexConfig struct {
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
	// Edition of index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#edition KendraIndex#edition}
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// Name of index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#name KendraIndex#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#role_arn KendraIndex#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Capacity units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#capacity_units KendraIndex#capacity_units}
	CapacityUnits *KendraIndexCapacityUnits `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description for the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#description KendraIndex#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Document metadata configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#document_metadata_configurations KendraIndex#document_metadata_configurations}
	DocumentMetadataConfigurations interface{} `field:"optional" json:"documentMetadataConfigurations" yaml:"documentMetadataConfigurations"`
	// Server side encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#server_side_encryption_configuration KendraIndex#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration *KendraIndexServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Tags for labeling the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#tags KendraIndex#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#user_context_policy KendraIndex#user_context_policy}.
	UserContextPolicy *string `field:"optional" json:"userContextPolicy" yaml:"userContextPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#user_token_configurations KendraIndex#user_token_configurations}.
	UserTokenConfigurations interface{} `field:"optional" json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}

