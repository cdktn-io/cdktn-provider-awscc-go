// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowsource


type MediaconnectFlowSourceGatewayBridgeSourceVpcInterfaceAttachmentA struct {
	// The name of the VPC interface to use for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow_source#vpc_interface_name MediaconnectFlowSourceA#vpc_interface_name}
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

