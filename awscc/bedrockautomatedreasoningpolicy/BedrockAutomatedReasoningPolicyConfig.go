// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockAutomatedReasoningPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#name BedrockAutomatedReasoningPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#description BedrockAutomatedReasoningPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether to force delete the automated reasoning policy even if it has active resources.
	//
	// When false , Amazon Bedrock validates if all artifacts have been deleted (e.g. policy version, test case, test result) for a policy before deletion. When true , Amazon Bedrock will delete the policy and all its artifacts without validation. Default is false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#force_delete BedrockAutomatedReasoningPolicy#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// The KMS key with which the Policy's assets will be encrypted at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#kms_key_id BedrockAutomatedReasoningPolicy#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#policy_definition BedrockAutomatedReasoningPolicy#policy_definition}.
	PolicyDefinition *BedrockAutomatedReasoningPolicyPolicyDefinition `field:"optional" json:"policyDefinition" yaml:"policyDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#tags BedrockAutomatedReasoningPolicy#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

