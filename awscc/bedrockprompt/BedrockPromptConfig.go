// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockPromptConfig struct {
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
	// Name for a prompt resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#name BedrockPrompt#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A KMS key ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#customer_encryption_key_arn BedrockPrompt#customer_encryption_key_arn}
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// Name for a variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#default_variant BedrockPrompt#default_variant}
	DefaultVariant *string `field:"optional" json:"defaultVariant" yaml:"defaultVariant"`
	// Name for a prompt resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#description BedrockPrompt#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#tags BedrockPrompt#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// List of prompt variants.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#variants BedrockPrompt#variants}
	Variants interface{} `field:"optional" json:"variants" yaml:"variants"`
}

