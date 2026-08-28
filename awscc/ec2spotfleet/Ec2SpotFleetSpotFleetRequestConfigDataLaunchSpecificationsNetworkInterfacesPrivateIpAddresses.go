// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesPrivateIpAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_spot_fleet#primary Ec2SpotFleet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_spot_fleet#private_ip_address Ec2SpotFleet#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

