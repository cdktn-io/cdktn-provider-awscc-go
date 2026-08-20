// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceNetworkInterfacesEnaSrdSpecification struct {
	// Specifies whether ENA Express is enabled for the network interface when you launch an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_instance#ena_srd_enabled Ec2Instance#ena_srd_enabled}
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// Contains ENA Express settings for UDP network traffic for the network interface that's attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_instance#ena_srd_udp_specification Ec2Instance#ena_srd_udp_specification}
	EnaSrdUdpSpecification *Ec2InstanceNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

