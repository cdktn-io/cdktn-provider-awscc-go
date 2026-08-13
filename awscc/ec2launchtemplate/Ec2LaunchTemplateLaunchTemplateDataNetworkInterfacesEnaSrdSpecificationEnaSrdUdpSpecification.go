// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification struct {
	// Indicates whether UDP traffic to and from the instance uses ENA Express.
	//
	// To specify this setting, you must first enable ENA Express.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_launch_template#ena_srd_udp_enabled Ec2LaunchTemplate#ena_srd_udp_enabled}
	EnaSrdUdpEnabled interface{} `field:"optional" json:"enaSrdUdpEnabled" yaml:"enaSrdUdpEnabled"`
}

