// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package finspaceenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FinspaceEnvironmentConfig struct {
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
	// Name of the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#name FinspaceEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ARNs of FinSpace Data Bundles to install.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#data_bundles FinspaceEnvironment#data_bundles}
	DataBundles *[]*string `field:"optional" json:"dataBundles" yaml:"dataBundles"`
	// Description of the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#description FinspaceEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Federation mode used with the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#federation_mode FinspaceEnvironment#federation_mode}
	FederationMode *string `field:"optional" json:"federationMode" yaml:"federationMode"`
	// Additional parameters to identify Federation mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#federation_parameters FinspaceEnvironment#federation_parameters}
	FederationParameters *FinspaceEnvironmentFederationParameters `field:"optional" json:"federationParameters" yaml:"federationParameters"`
	// KMS key used to encrypt customer data within FinSpace Environment infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#kms_key_id FinspaceEnvironment#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Parameters of the first Superuser for the FinSpace Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#superuser_parameters FinspaceEnvironment#superuser_parameters}
	SuperuserParameters *FinspaceEnvironmentSuperuserParameters `field:"optional" json:"superuserParameters" yaml:"superuserParameters"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#tags FinspaceEnvironment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

