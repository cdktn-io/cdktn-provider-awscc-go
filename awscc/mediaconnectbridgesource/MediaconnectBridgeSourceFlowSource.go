// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridgesource


type MediaconnectBridgeSourceFlowSource struct {
	// The ARN of the cloud flow used as a source of this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#flow_arn MediaconnectBridgeSource#flow_arn}
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The name of the VPC interface attachment to use for this source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#flow_vpc_interface_attachment MediaconnectBridgeSource#flow_vpc_interface_attachment}
	FlowVpcInterfaceAttachment *MediaconnectBridgeSourceFlowSourceFlowVpcInterfaceAttachment `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
}

