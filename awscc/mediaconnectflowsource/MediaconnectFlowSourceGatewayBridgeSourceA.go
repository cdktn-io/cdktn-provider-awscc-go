// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowsource


type MediaconnectFlowSourceGatewayBridgeSourceA struct {
	// The ARN of the bridge feeding this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow_source#bridge_arn MediaconnectFlowSourceA#bridge_arn}
	BridgeArn *string `field:"optional" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the VPC interface attachment to use for this bridge source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow_source#vpc_interface_attachment MediaconnectFlowSourceA#vpc_interface_attachment}
	VpcInterfaceAttachment *MediaconnectFlowSourceGatewayBridgeSourceVpcInterfaceAttachmentA `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

