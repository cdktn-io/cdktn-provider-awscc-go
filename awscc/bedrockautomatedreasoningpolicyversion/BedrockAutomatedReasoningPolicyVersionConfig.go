// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicyversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockAutomatedReasoningPolicyVersionConfig struct {
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
	// Arn of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_automated_reasoning_policy_version#policy_arn BedrockAutomatedReasoningPolicyVersion#policy_arn}
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// The hash for this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_automated_reasoning_policy_version#last_updated_definition_hash BedrockAutomatedReasoningPolicyVersion#last_updated_definition_hash}
	LastUpdatedDefinitionHash *string `field:"optional" json:"lastUpdatedDefinitionHash" yaml:"lastUpdatedDefinitionHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_automated_reasoning_policy_version#tags BedrockAutomatedReasoningPolicyVersion#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

