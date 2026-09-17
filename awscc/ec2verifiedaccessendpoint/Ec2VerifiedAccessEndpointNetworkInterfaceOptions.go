// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessendpoint


type Ec2VerifiedAccessEndpointNetworkInterfaceOptions struct {
	// The ID of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#network_interface_id Ec2VerifiedAccessEndpoint#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The IP port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#port Ec2VerifiedAccessEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The list of port ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#port_ranges Ec2VerifiedAccessEndpoint#port_ranges}
	PortRanges interface{} `field:"optional" json:"portRanges" yaml:"portRanges"`
	// The IP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#protocol Ec2VerifiedAccessEndpoint#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

