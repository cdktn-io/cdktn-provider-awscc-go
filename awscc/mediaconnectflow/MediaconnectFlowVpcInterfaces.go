// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowVpcInterfaces struct {
	// Immutable and has to be a unique against other VpcInterfaces in this Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#name MediaconnectFlow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// IDs of the network interfaces created in customer's account by MediaConnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#network_interface_ids MediaconnectFlow#network_interface_ids}
	NetworkInterfaceIds *[]*string `field:"optional" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// The type of network adapter that you want MediaConnect to use on this interface.
	//
	// If you don't set this value, it defaults to ENA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#network_interface_type MediaconnectFlow#network_interface_type}
	NetworkInterfaceType *string `field:"optional" json:"networkInterfaceType" yaml:"networkInterfaceType"`
	// Role Arn MediaConnect can assume to create ENIs in customer's account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#role_arn MediaconnectFlow#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Security Group IDs to be used on ENI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#security_group_ids MediaconnectFlow#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Subnet must be in the AZ of the Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#subnet_id MediaconnectFlow#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

