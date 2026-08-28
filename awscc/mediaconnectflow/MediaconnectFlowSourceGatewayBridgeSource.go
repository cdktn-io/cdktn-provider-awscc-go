// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceGatewayBridgeSource struct {
	// The ARN of the bridge feeding this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#bridge_arn MediaconnectFlow#bridge_arn}
	BridgeArn *string `field:"optional" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the VPC interface attachment to use for this bridge source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#vpc_interface_attachment MediaconnectFlow#vpc_interface_attachment}
	VpcInterfaceAttachment *MediaconnectFlowSourceGatewayBridgeSourceVpcInterfaceAttachment `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

