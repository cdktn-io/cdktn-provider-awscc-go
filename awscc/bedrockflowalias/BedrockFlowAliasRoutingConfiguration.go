// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockflowalias


type BedrockFlowAliasRoutingConfiguration struct {
	// Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_flow_alias#flow_version BedrockFlowAlias#flow_version}
	FlowVersion *string `field:"optional" json:"flowVersion" yaml:"flowVersion"`
}

