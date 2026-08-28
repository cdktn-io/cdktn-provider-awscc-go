// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationlibrary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationLibraryConfig struct {
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
	// Name of the DataAutomationLibrary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_library#library_name BedrockDataAutomationLibrary#library_name}
	LibraryName *string `field:"required" json:"libraryName" yaml:"libraryName"`
	// KMS Encryption Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_library#encryption_configuration BedrockDataAutomationLibrary#encryption_configuration}
	EncryptionConfiguration *BedrockDataAutomationLibraryEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Description of the DataAutomationLibrary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_library#library_description BedrockDataAutomationLibrary#library_description}
	LibraryDescription *string `field:"optional" json:"libraryDescription" yaml:"libraryDescription"`
	// List of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_automation_library#tags BedrockDataAutomationLibrary#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

