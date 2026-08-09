// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeSourcesFlowSource struct {
	// The ARN of the cloud flow used as a source of this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_bridge#flow_arn MediaconnectBridge#flow_arn}
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The name of the VPC interface attachment to use for this source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_bridge#flow_vpc_interface_attachment MediaconnectBridge#flow_vpc_interface_attachment}
	FlowVpcInterfaceAttachment *MediaconnectBridgeSourcesFlowSourceFlowVpcInterfaceAttachment `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
	// The name of the flow source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

