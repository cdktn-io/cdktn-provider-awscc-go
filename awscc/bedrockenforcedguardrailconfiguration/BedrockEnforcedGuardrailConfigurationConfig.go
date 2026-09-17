// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockenforcedguardrailconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockEnforcedGuardrailConfigurationConfig struct {
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
	// Identifier for the guardrail, could be the ID or the ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_enforced_guardrail_configuration#guardrail_identifier BedrockEnforcedGuardrailConfiguration#guardrail_identifier}
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Numerical guardrail version (not DRAFT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_enforced_guardrail_configuration#guardrail_version BedrockEnforcedGuardrailConfiguration#guardrail_version}
	GuardrailVersion *string `field:"required" json:"guardrailVersion" yaml:"guardrailVersion"`
	// Model-specific information for the enforced guardrail configuration. If not present, the configuration is enforced on all models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_enforced_guardrail_configuration#model_enforcement BedrockEnforcedGuardrailConfiguration#model_enforcement}
	ModelEnforcement *BedrockEnforcedGuardrailConfigurationModelEnforcement `field:"optional" json:"modelEnforcement" yaml:"modelEnforcement"`
	// Selective content guarding controls for enforced guardrails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_enforced_guardrail_configuration#selective_content_guarding BedrockEnforcedGuardrailConfiguration#selective_content_guarding}
	SelectiveContentGuarding *BedrockEnforcedGuardrailConfigurationSelectiveContentGuarding `field:"optional" json:"selectiveContentGuarding" yaml:"selectiveContentGuarding"`
}

