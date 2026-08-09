// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationProjectConfig struct {
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
	// Name of the DataAutomationProject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#project_name BedrockDataAutomationProject#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Custom output configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#custom_output_configuration BedrockDataAutomationProject#custom_output_configuration}
	CustomOutputConfiguration *BedrockDataAutomationProjectCustomOutputConfiguration `field:"optional" json:"customOutputConfiguration" yaml:"customOutputConfiguration"`
	// KMS encryption context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#kms_encryption_context BedrockDataAutomationProject#kms_encryption_context}
	KmsEncryptionContext *map[string]*string `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// KMS key identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#kms_key_id BedrockDataAutomationProject#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Override configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#override_configuration BedrockDataAutomationProject#override_configuration}
	OverrideConfiguration *BedrockDataAutomationProjectOverrideConfiguration `field:"optional" json:"overrideConfiguration" yaml:"overrideConfiguration"`
	// Description of the DataAutomationProject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#project_description BedrockDataAutomationProject#project_description}
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// Type of the DataAutomationProject - Sync or Async.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#project_type BedrockDataAutomationProject#project_type}
	ProjectType *string `field:"optional" json:"projectType" yaml:"projectType"`
	// Standard output configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#standard_output_configuration BedrockDataAutomationProject#standard_output_configuration}
	StandardOutputConfiguration *BedrockDataAutomationProjectStandardOutputConfiguration `field:"optional" json:"standardOutputConfiguration" yaml:"standardOutputConfiguration"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#tags BedrockDataAutomationProject#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

