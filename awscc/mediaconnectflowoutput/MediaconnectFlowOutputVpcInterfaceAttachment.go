// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput


type MediaconnectFlowOutputVpcInterfaceAttachment struct {
	// The name of the VPC interface to use for this output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow_output#vpc_interface_name MediaconnectFlowOutput#vpc_interface_name}
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

