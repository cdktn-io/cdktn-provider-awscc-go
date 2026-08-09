// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6Prefixes struct {
	// The IPv6 prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#ipv_6_prefix Ec2LaunchTemplate#ipv_6_prefix}
	Ipv6Prefix *string `field:"optional" json:"ipv6Prefix" yaml:"ipv6Prefix"`
}

