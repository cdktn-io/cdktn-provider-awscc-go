// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicy


type BedrockAutomatedReasoningPolicyPolicyDefinitionRules struct {
	// An alternate expression for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_automated_reasoning_policy#alternate_expression BedrockAutomatedReasoningPolicy#alternate_expression}
	AlternateExpression *string `field:"optional" json:"alternateExpression" yaml:"alternateExpression"`
	// The SMT expression for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_automated_reasoning_policy#expression BedrockAutomatedReasoningPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A unique id within the PolicyDefinition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_automated_reasoning_policy#id BedrockAutomatedReasoningPolicy#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

