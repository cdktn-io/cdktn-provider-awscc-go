// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockpromptversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockPromptVersionConfig struct {
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
	// ARN of a prompt resource possibly with a version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt_version#prompt_arn BedrockPromptVersion#prompt_arn}
	PromptArn *string `field:"required" json:"promptArn" yaml:"promptArn"`
	// Description for a prompt version resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt_version#description BedrockPromptVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt_version#tags BedrockPromptVersion#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

