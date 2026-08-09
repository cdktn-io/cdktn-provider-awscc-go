// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification struct {
	// Indicates whether UDP traffic uses ENA Express for your instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_instance#ena_srd_udp_enabled Ec2Instance#ena_srd_udp_enabled}
	EnaSrdUdpEnabled interface{} `field:"optional" json:"enaSrdUdpEnabled" yaml:"enaSrdUdpEnabled"`
}

