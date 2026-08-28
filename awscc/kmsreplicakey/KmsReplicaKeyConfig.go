// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmsreplicakey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KmsReplicaKeyConfig struct {
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
	// The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kms_replica_key#key_policy KmsReplicaKey#key_policy}
	KeyPolicy *string `field:"required" json:"keyPolicy" yaml:"keyPolicy"`
	// Identifies the primary AWS KMS key to create a replica of.
	//
	// Specify the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kms_replica_key#primary_key_arn KmsReplicaKey#primary_key_arn}
	PrimaryKeyArn *string `field:"required" json:"primaryKeyArn" yaml:"primaryKeyArn"`
	// A description of the AWS KMS key.
	//
	// Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kms_replica_key#description KmsReplicaKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kms_replica_key#enabled KmsReplicaKey#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack.
	//
	// Enter a value between 7 and 30 days. The default value is 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kms_replica_key#pending_window_in_days KmsReplicaKey#pending_window_in_days}
	PendingWindowInDays *float64 `field:"optional" json:"pendingWindowInDays" yaml:"pendingWindowInDays"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kms_replica_key#tags KmsReplicaKey#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

