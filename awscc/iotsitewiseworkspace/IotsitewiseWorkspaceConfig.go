// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseworkspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseWorkspaceConfig struct {
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
	// The encryption configuration for the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_workspace#encryption_configuration IotsitewiseWorkspace#encryption_configuration}
	EncryptionConfiguration *IotsitewiseWorkspaceEncryptionConfiguration `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_workspace#workspace_name IotsitewiseWorkspace#workspace_name}
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// The ARN of the AWS KMS key used for KMS_BASED_ENCRYPTION. Required when EncryptionConfiguration.EncryptionType is KMS_BASED_ENCRYPTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_workspace#kms_key_id IotsitewiseWorkspace#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_workspace#tags IotsitewiseWorkspace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A description of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_workspace#workspace_description IotsitewiseWorkspace#workspace_description}
	WorkspaceDescription *string `field:"optional" json:"workspaceDescription" yaml:"workspaceDescription"`
}

