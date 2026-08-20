// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMediaConnectFlow struct {
	// The ARN of the flow to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#flow_arn MediaconnectRouterInput#flow_arn}
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The ARN of the flow output to connect to this router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#flow_output_arn MediaconnectRouterInput#flow_output_arn}
	FlowOutputArn *string `field:"optional" json:"flowOutputArn" yaml:"flowOutputArn"`
	// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#source_transit_decryption MediaconnectRouterInput#source_transit_decryption}
	SourceTransitDecryption *MediaconnectRouterInputConfigurationMediaConnectFlowSourceTransitDecryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

