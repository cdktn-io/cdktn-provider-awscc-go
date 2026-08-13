// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockflowalias


type BedrockFlowAliasConcurrencyConfiguration struct {
	// Number of nodes executed concurrently at a time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_flow_alias#max_concurrency BedrockFlowAlias#max_concurrency}
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_flow_alias#type BedrockFlowAlias#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

