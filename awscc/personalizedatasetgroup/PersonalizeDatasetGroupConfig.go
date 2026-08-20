// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizedatasetgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeDatasetGroupConfig struct {
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
	// The name for the new dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset_group#name PersonalizeDatasetGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain of a Domain dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset_group#domain PersonalizeDatasetGroup#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset_group#kms_key_arn PersonalizeDatasetGroup#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key.
	//
	// Supplying an IAM role is only valid when also specifying a KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset_group#role_arn PersonalizeDatasetGroup#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset_group#tags PersonalizeDatasetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

