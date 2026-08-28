// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockintelligentpromptrouter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockIntelligentPromptRouterConfig struct {
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
	// Model configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_intelligent_prompt_router#fallback_model BedrockIntelligentPromptRouter#fallback_model}
	FallbackModel *BedrockIntelligentPromptRouterFallbackModel `field:"required" json:"fallbackModel" yaml:"fallbackModel"`
	// List of model configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_intelligent_prompt_router#models BedrockIntelligentPromptRouter#models}
	Models interface{} `field:"required" json:"models" yaml:"models"`
	// Name of the Prompt Router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_intelligent_prompt_router#prompt_router_name BedrockIntelligentPromptRouter#prompt_router_name}
	PromptRouterName *string `field:"required" json:"promptRouterName" yaml:"promptRouterName"`
	// Represents the criteria used for routing requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_intelligent_prompt_router#routing_criteria BedrockIntelligentPromptRouter#routing_criteria}
	RoutingCriteria *BedrockIntelligentPromptRouterRoutingCriteria `field:"required" json:"routingCriteria" yaml:"routingCriteria"`
	// Description of the Prompt Router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_intelligent_prompt_router#description BedrockIntelligentPromptRouter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_intelligent_prompt_router#tags BedrockIntelligentPromptRouter#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

